CREATE TABLE departments (
   id VARCHAR (36) UNIQUE PRIMARY KEY,
   name VARCHAR (255) NOT NULL,
   description VARCHAR (255)
);

CREATE TABLE employees (
   id VARCHAR (36) UNIQUE PRIMARY KEY,
   department_id VARCHAR (36) NOT NULL,
   name VARCHAR (255) NOT NULL,
   email VARCHAR (255) NOT NULL,
   phone VARCHAR (255),
   FOREIGN KEY (department_id) REFERENCES departments (id)
);
