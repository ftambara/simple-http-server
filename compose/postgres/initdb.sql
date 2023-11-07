CREATE TABLE notes (
    id  SERIAL PRIMARY KEY,
    title VARCHAR(100),
    body TEXT,
    created TIMESTAMP WITH TIME ZONE,
    modified TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_notes_created ON notes(created);

