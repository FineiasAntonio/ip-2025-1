CREATE TABLE IF NOT EXISTS Professores (
    idProfessor SERIAL PRIMARY KEY,
    nome VARCHAR(10) NOT NULL,
    formacao VARCHAR(50) NOT NULL,
    ativo BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE Materias (
    idMateria SERIAL PRIMARY KEY,
    nome VARCHAR(10) NOT NULL,
    cargaHoraria INT NOT NULL,
    idProfessor INT NOT NULL,
    
    FOREIGN KEY (idProfessor) REFERENCES Professores(idProfessor) ON UPDATE CASCADE
);