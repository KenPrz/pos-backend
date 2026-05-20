package database

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Migrator struct {
	DB            *sql.DB
	MigrationsDir string
}

func NewMigrator(db *sql.DB, dir string) *Migrator {
	return &Migrator{
		DB:            db,
		MigrationsDir: dir,
	}
}

func (m *Migrator) Run() error {
	files, err := m.loadFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := m.runFile(file); err != nil {
			return fmt.Errorf("migration failed on %s: %w", file, err)
		}
	}

	return nil
}

func (m *Migrator) loadFiles() ([]string, error) {
	var files []string

	err := filepath.WalkDir(m.MigrationsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".sql") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(files)
	return files, nil
}

func (m *Migrator) runFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	queries := splitSQL(string(content))

	for _, q := range queries {
		q = strings.TrimSpace(q)
		if q == "" {
			continue
		}

		if _, err := tx.Exec(q); err != nil {
			return fmt.Errorf("query failed: %w\nSQL: %s", err, q)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	fmt.Println("Applied:", filepath.Base(path))
	return nil
}

func splitSQL(script string) []string {
	return strings.Split(script, ";")
}
