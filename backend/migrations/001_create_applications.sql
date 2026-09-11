CREATE TABLE IF NOT EXISTS applications (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    contact TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS applications_created_at_idx
    ON applications (created_at DESC);
