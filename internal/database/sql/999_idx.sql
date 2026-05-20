-- ALWAYS LAST SINCE THESE ARE INDEXES
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(user_name);
CREATE INDEX idx_user_passwords_user_id ON user_passwords(user_id);
