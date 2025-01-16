-- name: CreateProduct :exec
INSERT INTO product (
    product_url,
    product_category,
    description,
    price,
    quantity,
    product_name,
    created_at,
    updated_at
) VALUES (
             ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
         );
-- name: UpdateProduct :exec
UPDATE product
SET
    product_url = ?,
    product_category = ?,
    description = ?,
    price = ?,
    quantity = ?,
    product_name = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: deleteProduct :exec
update product set status = ?;

-- name: GetProductByNameSQLC :one
SELECT p.* FROM product AS p WHERE p.product_name = ? and p.status = 1 LIMIT 1;

-- name: GetProductAllSQLC :many
SELECT p.*
FROM product AS p
WHERE p.status = 1
    LIMIT ? OFFSET ?;
-- name: GetProduct :one
SELECT p.* FROM product AS p WHERE p.id = ? LIMIT 1;

