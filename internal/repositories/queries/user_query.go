package queries

const (
	CreateUsers = `
	INSERT INTO users (
		username,
		full_name, 
		address, 
		role, 
		password)
	VALUES($1, $2, $3, $4, $5)
	RETURNING username
	`

	GetOneUsersByUsername = `
	SELECT
		u.user_id,
		u.username,
		u.full_name,
		u.address,
		u.role,
		u.balance,
		u.created_at,
		u.updated_at
	FROM users as u
	WHERE u.username = $1
	LIMIT 1
	`

	GetPasswordByUsername = `
	SELECT
		password
	FROM users
	WHERE username = $1
	LIMIT 1
	`
)
