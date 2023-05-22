package queries

const (
	GetListProducts = `
	SELECT
		p.product_id,
		p.code,
		p.title,
		p.price,
		p.weight,
		p.stock,
		p.address,
		p.category,
		p.created_at,
		p.updated_at
	FROM products as p
	WHERE p.category IS NULL 
	OR p.category = $1
	ORDER BY p.created_at
	LIMIT $2
	OFFSET $3
	`
)
