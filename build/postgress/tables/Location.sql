BEGIN TRANSACTION;

CREATE TABLE locations (
    location_id SERIAL PRIMARY KEY,
    addr TEXT
)

COMMIT;