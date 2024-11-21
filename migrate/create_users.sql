CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    name TEXT,
    surname TEXT,
    patronymic TEXT,
    email TEXT UNIQUE
);