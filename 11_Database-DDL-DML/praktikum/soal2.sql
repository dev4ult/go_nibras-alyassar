CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  birthdate DATE,
  user_status ENUM('active', 'de-active'),
  gender ENUM('male', 'female'),
  created_at TIMESTAMP DEFAULT current_timestamp(),
  updated_at TIMESTAMP DEFAULT current_timestamp()
);

CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  operator_id INT,
  type_id INT,
  product VARCHAR(150) NOT NULL
);

CREATE TABLE product_types (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_type VARCHAR(20) NOT NULL
);

CREATE TABLE operators (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name varchar(100) NOT NULL,
  address varchar(150),
  phone_number varchar(15),
  birthdate DATE,
  gender ENUM('male','female'),
  created_at TIMESTAMP DEFAULT current_timestamp(),
  updated_at TIMESTAMP DEFAULT current_timestamp()
);

CREATE TABLE product_descriptions (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_id INT UNIQE NOT NULL,
  product_description TEXT
);

CREATE TABLE payment_methods (
  id INT AUTO_INCREMENT PRIMARY KEY,
  payment_method VARCHAR(50) NOT NULL
);

CREATE TABLE transactions (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL
);

CREATE TABLE transaction_details (
  id INT AUTO_INCREMENT PRIMARY KEY,
  transaction_id INT NOT NULL,
  product_id INT NOT NULL
);

CREATE TABLE kurir (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp(),
  updated_at TIMESTAMP DEFAULT current_timestamp()
);

ALTER TABLE kurir
  ADD ongkos_dasar INT;
  
ALTER TABLE kurir
  RENAME TO shipping;
  
DROP TABLE shipping;

CREATE TABLE payment_method_descriptions (
  id INT AUTO_INCREMENT PRIMARY KEY,
  payment_method_id INT UNIQUE NOT NULL,
  payment_method_description VARCHAR(255)
);

CREATE TABLE user_payment_method_details (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  payment_method_id NOT NULL
);