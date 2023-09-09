SELECT *
FROM transactions
JOIN users
ON users.id = transactions.user_id
WHERE users.id <= 2;

SELECT SUM(transaction_details.price) as Total_Price
FROM transaction_details
JOIN transactions
ON transaction_details.transaction_id = transactions.id
WHERE transactions.user_id = 1;

SELECT SUM(transaction_details.price) as Total_Price
FROM transaction_details
JOIN transactions
ON transaction_details.transaction_id = transactions.id
JOIN products
ON transaction_details.product_id = products.id
WHERE products.product_type_id = 2;

SELECT 
  products.*,  product_types.name as 'product type'
FROM products
JOIN product_types
ON products.product_type_id = product_types.id;

SELECT
  users.name as 'user name', products.name as 'product name', transaction_details.*
FROM transaction_details
JOIN products
ON products.id = transaction_details.product_id
JOIN transactions
ON transactions.id = transaction_details.transaction_id
JOIN users
ON users.id = transactions.user_id;

DELIMITER //
CREATE FUNCTION DeleteDetailsAfterTransaction(id INT)
RETURNS INT




