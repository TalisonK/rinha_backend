CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

\c api_pessoas

CREATE USER user_pessoa;
ALTER USER user_pessoa WITH PASSWORD '1122';
GRANT ALL PRIVILEGES ON DATABASE api_pessoas TO user_pessoa;
Grant ALL PRIVILEGES ON TABLE pessoas TO user_pessoa;
Grant ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO user_pessoa;
