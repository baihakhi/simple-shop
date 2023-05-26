package queries

const (
	CreateCart = `
	INSERT INTO carts (user_id, product_id, amount, amount_price)
	VALUES($1, $2, $3, $4)
	RETURNING cart_id
	`

	DeleteCart = `
	DELETE from carts
	WHERE
		cart_id = $1
	`

	GetListCart = `
	SELECT
		cart_id,
		user_id,
		product_id,
		amount,
		amount_price,
		created_at,
		updated_at
	FROM carts as c
	WHERE c.user_id=$1
	ORDER BY c.created_at
	LIMIT $2
	OFFSET $3
	`

	GetOneCartsById = `
	SELECT
		cart_id,
		user_id,
		product_id,
		amount,
		amount_price,
		created_at,
		updated_at
	FROM carts 
	WHERE cart_id = $1
	LIMIT 1
	`
)
