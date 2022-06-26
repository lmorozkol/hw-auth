CREATE TABLE IF NOT EXISTS credentials (
                                           id SERIAL PRIMARY KEY,
                                           timestamp TIMESTAMP NOT NULL,
                                           login varchar(250) NOT NULL,
    password varchar(250) NOT NULL
    );
CREATE INDEX login_index ON credentials (login);
CREATE INDEX password_index ON credentials (password);