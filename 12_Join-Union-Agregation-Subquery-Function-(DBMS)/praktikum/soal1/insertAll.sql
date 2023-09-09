INSERT INTO operators (name) VALUES
  ("Nibras"),
  ("Sarbin"),
  ("Udin"),
  ("Rena"),
  ("Didin");
  
INSERT INTO product_types (name) VALUES
  ("kitchen"),
  ("electronic"),
  ("furniture");
  
INSERT INTO products (operator_id, product_type_id, name, code, status) VALUES
  (3, 1, "Small Knife", "COOKING_KNIFE", 10),
  (3, 1, "Medium Spatulla", "COOKING_SPATULLA", 3);
  
INSERT INTO products (operator_id, product_type_id, name, code, status) VALUES
  (1, 2, "Sumsang Phone Charger", "PHONE_CHARGER", 50),
  (1, 2, "APPA Tablet", "TABLET", 20),
  (1, 2, "Qeykron Keyboard Ligth", "KEYBOARD", 34);
  
INSERT INTO products (operator_id, product_type_id, name, code, status) VALUES
  (4, 3, "AX5 Office Chair", "OFFICE_CHAIR", 15),
  (4, 3, "X7D Dinner Table", "TABLE", 20),
  (4, 3, "IKAE Office Chair", "OFFICE_CHAIR", 32);
  
INSERT INTO product_descriptions (product_id, description) VALUES
  SELECT products.id, "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur,"
  FROM products;
  
INSERT INTO payment_methods (name, status) VALUES
  ("CBA VA", 50),
  ("DEPENDEN MOBILE", 100),
  ("ONGKOS", 45);
  
INSERT INTO users (name, dob, status, gender) VALUES
  ("Nibras", "2003-03-03", 98, "M"),
  ("Fikri", "2003-07-07", 67, "M"),
  ("Richard", "2003-08-08", 79, "M"),
  ("Doni", "2003-04-04", 99, "M"),
  ("Dini", "2003-05-05", 100, "F");
  
DELIMITER //
CREATE OR REPLACE PROCEDURE InsertTransactionsThreeTimes()
BEGIN
DECLARE i INT DEFAULT 1;
WHILE(i <= 3) DO
  INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
    SELECT users.id, 2, "PENDING", 0, 0
    FROM users;
  SET i = i + 1;
END WHILE;
END;
//
DELIMITER ;

CALL InsertTransactionsThreeTimes();

DELIMITER //
CREATE OR REPLACE PROCEDURE InsertProductToDTSThreeTimes()
BEGIN
DECLARE i INT DEFAULT 1;
WHILE(i <= 3) DO
  INSERT INTO transactions (transaction_id, product_id, status, qty, price)
    SELECT transactions.id, (SELECT FLOOR(RAND()*(3 - 1 + 1)) + 1), "CHECKED", 5, 15000
    FROM transactions;
  SET i = i + 1;
END WHILE;
END;
//
DELIMITER ;

CALL InsertProductToDTSThreeTimes();