CREATE TABLE IF NOT EXISTS entries (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    foods TEXT[] NOT NULL,
    food_desc TEXT,
    rating SMALLINT CHECK (rating >= 0 AND rating <= 10),
    rating_desc TEXT
);
