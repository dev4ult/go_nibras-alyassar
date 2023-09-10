-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT *
FROM users
WHERE gender = 'M';

-- Tampilkan product dengan id = 3.
SELECT *
FROM products
WHERE id = 3;

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT *
FROM users
WHERE created_at >= DATE(NOW() - INTERVAL 7 DAY) 
  AND name LIKE '%a%';

-- Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*)
FROM users
WHERE gender = 'F';

-- Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT *
FROM users
ORDER BY name ASC;

-- Tampilkan 5 data pada data product
SELECT *
FROM products
LIMIT 5;
