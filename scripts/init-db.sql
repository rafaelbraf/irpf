SELECT 'CREATE DATABASE irpf' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'irpf');

CREATE TABLE IF NOT EXISTS contribuintes (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(11) NOT NULL,
    nome VARCHAR(255),
    celular VARCHAR(20),
    endereco VARCHAR(255),
    data_nascimento DATE,
    email VARCHAR(255),
    natureza_ocupacao VARCHAR(255),
    ocupacao_principal VARCHAR(255)
);