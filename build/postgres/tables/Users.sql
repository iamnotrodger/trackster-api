BEGIN TRANSACTION;
CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(30) UNIQUE,
    provider_id TEXT,
    email VARCHAR(254),
    picture TEXT,
    given_name TEXT,
    family_name TEXT,
    joined TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
COMMIT;