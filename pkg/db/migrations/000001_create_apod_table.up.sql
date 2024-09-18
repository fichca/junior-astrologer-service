CREATE TABLE apod
(
    id          UUID PRIMARY KEY,
    title       TEXT NOT NULL,
    explanation TEXT NOT NULL,
    copyright   TEXT NOT NULL,
    date        DATE
);