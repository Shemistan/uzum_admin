-- name: GetProduct :one
SELECT pr.* FROM products pr
WHERE pr.id = @id;

-- name: GetProducts :many
SELECT pr.* FROM products pr
ORDER BY pr.id;

-- name: CreateProduct :exec
INSERT INTO products (name, description, price, quantity)
VALUES (@name, @description, @price, @quantity);

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = @id;

-- name: UpdateProduct :exec
UPDATE products
SET name = @name, description = @description, price = @price, quantity = @quantity
WHERE id = @id;
