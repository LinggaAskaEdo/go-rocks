package user

const (
	CreateUser = `
		INSERT INTO user 
		(
			username, 
			email,
			phone,
			division_id,
			password,
			created_at
		) 
		VALUES 
			(?,?,?,?,?,?);
	`

	GetUserByID = `
		SELECT
			user.id,
			user.username,
			user.email,
			user.phone,
			user.password,
			user.is_deleted,
			user.created_at,
			user.updated_at,
			user.deleted_at
		FROM 
			user
		WHERE
			user.id = ?
	`

	GetUserByUsername = `
		SELECT
			user.id,
			user.username,
			user.email,
			user.phone,
			user.password,
			user.is_deleted,
			user.created_at,
			user.updated_at,
			user.deleted_at
		FROM 
			user
		WHERE
			user.username = ?
	`
)
