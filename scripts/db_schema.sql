/* Creating required Tables */

CREATE TABLE patient
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100),
    illness     VARCHAR(200),
    birth_date  DATE,
    location_id INT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES location (id)
);

CREATE TABLE location
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100),
    hospital_id INT,
    FOREIGN KEY (hospital_id) REFERENCES hospital (id)
);

CREATE TABLE hospital
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(100),
    max_patient_amount INT
);

/* Inserting data into Hospital Table */

INSERT INTO hospital (name, max_patient_amount)
VALUES ('National Government Hospital', 2500),
       ('State Government Hospital', 1500);

/* Inserting data into Location Table */

INSERT INTO location (name, hospital_id)
VALUES ('Registration Desk', 1),
       ('Laboratory', 1),
       ('Pharmacy', 2),
       ('Isolation Ward', 2);

/* Inserting data into Patient Table */

INSERT INTO patient (name, illness, birth_date, location_id)
VALUES ('Ray Ma', 'Fever', '2000-05-15', 1),
       ('Tejas Mate', 'Cold', '2004-10-12', 2),
       ('Ribhav Sharma', 'Diabetes', '1990-08-01', 3),
       ('Hrishikesh Patil', 'Asthama', '1981-12-23', 4),
       ('Nayan Jindal', 'Flu', '2006-01-30', 1),
       ('Ishaan Das', 'Dengue', '2003-11-19', 2),
       ('Mihir Aggarwal', 'Cough', '1995-03-07', 3),
       ('Khushi Walia', 'Cold', '1985-06-01', 4),
       ('Inesh Tickoo', 'Malaria', '2009-10-31', 1);
