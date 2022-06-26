CREATE TABLE IF NOT EXISTS credentials (
    id SERIAL PRIMARY KEY,
    login varchar(250) NOT NULL,
    password varchar(250) NOT NULL,
    timestamp TIMESTAMP NOT NULL
    );
CREATE INDEX login_index ON credentials (login);
CREATE INDEX password_index ON credentials (password);

INSERT INTO credentials (login, password, "timestamp") VALUES ('admin', 'admin', current_timestamp);