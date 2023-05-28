CREATE TABLE metrics (
                         id SERIAL PRIMARY KEY,
                         timestamp TIMESTAMP DEFAULT NOW(),
                         name TEXT NOT NULL,
                         value DOUBLE PRECISION NOT NULL
);

CREATE INDEX idx_metrics_timestamp ON metrics (timestamp);
CREATE INDEX idx_metrics_name ON metrics (name);
