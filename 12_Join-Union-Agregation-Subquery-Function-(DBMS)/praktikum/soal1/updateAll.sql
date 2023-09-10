-- Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE product 
SET name = 'product dummy' 
WHERE id = 1;

-- Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transaction_details
SET qty = 3
WHERE product_id = 3;