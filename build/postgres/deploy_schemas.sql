\i '/docker-entrypoint-initdb.d/tables/Users.sql' 
\i '/docker-entrypoint-initdb.d/tables/Locations.sql' 
\i '/docker-entrypoint-initdb.d/tables/Contact.sql' 
\i '/docker-entrypoint-initdb.d/tables/Cases.sql' 
\i '/docker-entrypoint-initdb.d/tables/Trace.sql' 

-- Seeds
\i '/docker-entrypoint-initdb.d/seeds/seeds.sql' 