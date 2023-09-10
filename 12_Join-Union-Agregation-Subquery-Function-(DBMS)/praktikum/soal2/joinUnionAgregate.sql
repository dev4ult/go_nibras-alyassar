-- Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT *
FROM transactions
JOIN users
ON users.id = transactions.user_id
WHERE users.id <= 2;

-- Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(transaction_details.price) as Total_Price
FROM transaction_details
JOIN transactions
ON transaction_details.transaction_id = transactions.id
WHERE transactions.user_id = 1;

-- Tampilkan total transaksi dengan product type 2.
SELECT SUM(transaction_details.price) as Total_Price
FROM transaction_details
JOIN transactions
ON transaction_details.transaction_id = transactions.id
JOIN products
ON transaction_details.product_id = products.id
WHERE products.product_type_id = 2;

-- Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT 
  products.*,  product_types.name as 'product type'
FROM products
JOIN product_types
ON products.product_type_id = product_types.id;

-- Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT
  users.name as 'user name', products.name as 'product name', transaction_details.*
FROM transaction_details
JOIN products
ON products.id = transaction_details.product_id
JOIN transactions
ON transactions.id = transaction_details.transaction_id
JOIN users
ON users.id = transactions.user_id;

-- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER //
CREATE OR REPLACE TRIGGER delete_all_details_transaction
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
DECLARE v_transaction_id INT;
SET v_transaction_id = OLD.id;
DELETE FROM transaction_details
  WHERE transaction_id = v_transaction_id;
END //

-- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER //
CREATE OR REPLACE TRIGGER update_total_qty
BEFORE DELETE ON transaction_details FOR EACH ROW
BEGIN
DECLARE 
  v_transaction_id INT, v_qty INT,
SET v_transaction_id = OLD.transaction_id, v_qty = OLD.qty;
UPDATE transactions 
  SET total_qty = ((SELECT SUM(qty) 
    FROM transaction_details 
    WHERE transaction_id = v_transaction_id) - v_qty)
  WHERE id = v_transaction_id;
END //

-- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. 
SELECT products.* 
FROM products
WHERE 
NOT EXISTS (SELECT NULL 
            FROM transaction_details as td 
            WHERE products.id = td.product_id);