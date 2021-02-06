BEGIN TRANSACTION;

CREATE TABLE trace (
    trace_id SERIAL PRIMARY KEY,
    case_id INTEGER,
    user_id INTEGER,
    CONSTRAINT user_id_constraint FOREIGN KEY (user_id) REFERENCES users(user_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
    CONSTRAINT case_id_constraint FOREIGN KEY (user_id) REFERENCES cases(case_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

COMMIT;