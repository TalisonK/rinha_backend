
CREATE TABLE pessoas(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    nome varchar(255) NOT NULL,
    apelido varchar(255) NOT NULL,
    nascimento date NOT NULL,
    stack varchar(255)[] NOT NULL
);

INSERT INTO pessoas (nome, apelido, nascimento, stack) VALUES ('João', 'Joãozinho', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Maria', 'Mariazinha', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('José', 'Zézinho', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Joana', 'Joaninha', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Pedro', 'Pedrinho', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Paula', 'Paulinha', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Carlos', 'Carlinhos', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Carla', 'Carlinda', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Fernando', 'Fernandinho', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Fernanda', 'Fernandinha', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Ricardo', 'Ricardinho', '1990-01-01', '{Java, Python}');
Insert INTO pessoas (nome, apelido, nascimento, stack) VALUES ('Ricarda', 'Ricardinha', '1990-01-01', '{Java, Python}');