BEGIN TRANSACTION;

CREATE TABLE contact (
    contact_id SERIAL PRIMARY KEY,
    user_id INTEGER,
    location_id INTEGER, 
    contact_time TIMESTAMP,
    CONSTRAINT user_id_constraint FOREIGN KEY (user_id) REFERENCES users(user_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
    CONSTRAINT location_id_constraint FOREIGN KEY (location_id) REFERENCES locations(location_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

COMMIT;
