CREATE TABLE recipes (
    id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    keywords text,
    description text,
    url text,
    yield number,
    ingredients text NOT NULL,
    steps JSONB NOT NULL
);