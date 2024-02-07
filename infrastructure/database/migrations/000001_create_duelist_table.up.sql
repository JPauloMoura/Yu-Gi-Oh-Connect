CREATE TABLE  IF NOT EXISTS duelists (
    ID VARCHAR(50) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    BirthDate DATE NOT NULL,
    Street VARCHAR(255) NOT NULL,
    City VARCHAR(255) NOT NULL,
    State CHAR(2) NOT NULL,
    PostalCode VARCHAR(10) NOT NULL,
    Complement VARCHAR(255),
    Email VARCHAR(255) NOT NULL,
    Phone VARCHAR(20) NOT NULL,
    Presentation TEXT
);