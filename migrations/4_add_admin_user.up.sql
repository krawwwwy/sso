INSERT INTO users (id, email, pass_hash, is_admin)
VALUES (2, 'admin@mail.ru', 'admin123', 1)
ON CONFLICT DO NOTHING;