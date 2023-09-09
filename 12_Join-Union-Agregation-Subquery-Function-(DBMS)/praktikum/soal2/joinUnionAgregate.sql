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
