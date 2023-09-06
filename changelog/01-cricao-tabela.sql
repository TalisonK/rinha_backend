
CREATE TABLE pessoas(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    nome varchar(255) NOT NULL,
    apelido varchar(255) NOT NULL,
    nascimento date NOT NULL,
    stack varchar[] NOT NULL
);

INSERT INTO pessoas (nome, apelido, nascimento, stack) VALUES ('João', 'Joãozinho', '1990-01-01', '{Java, Python}');