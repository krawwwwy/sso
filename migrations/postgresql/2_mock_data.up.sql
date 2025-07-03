INSERT INTO apps (id, name, secret)
VALUES (1, 'test', 'test-secret')
ON CONFLICT DO NOTHING;

INSERT INTO users (id, email, pass_hash)
VALUES (1, 'admin@mail.ru', 'admin123')
ON CONFLICT DO NOTHING;

INSERT INTO admins (user_id)
VALUES (1)
ON CONFLICT DO NOTHING;