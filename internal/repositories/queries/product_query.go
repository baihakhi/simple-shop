package queries

const (
	GetListProducts = `
	SELECT
		product_id,
		code,
		title,
		price,
		weight,
		stock,
		address,
		category,
		created_at,
		updated_at
	FROM products as p
	WHERE p.category IS NULL 
	OR p.category = $1
	ORDER BY p.created_at
	LIMIT $2
	OFFSET $3
	`
	GetOneProductByID = `
	SELECT
		product_id,
		code,
		title,
		price,
		weight,
		stock,
		address,
		category,
		created_at,
		updated_at
	FROM products
	WHERE product_id = $1
	LIMIT 1
	`
)
