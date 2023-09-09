SELECT *
FROM users
WHERE gender = 'M';

SELECT *
FROM products
WHERE id = 3;

SELECT *
FROM users
WHERE created_at >= DATE(NOW() - INTERVAL 7 DAY) 
  AND name LIKE '%a%';

SELECT COUNT(*)
FROM users
WHERE gender = 'F';

SELECT *
FROM users
ORDER BY name ASC;

SELECT *
FROM products
LIMIT 5;
