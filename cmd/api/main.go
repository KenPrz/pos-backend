package main

import (
	nethttp "net/http"

	transporthttp "github.com/KenPrz/pos-backend/internal/transport/http"
)

func main() {
	r := transporthttp.NewRouter()

	nethttp.ListenAndServe(":8000", r)
}
