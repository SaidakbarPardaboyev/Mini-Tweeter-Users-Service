package static

var (
	InsertUserQuery = `
		INSERT INTO 
			users (
				id,
				name,
				surname,
				bio,
				profile_image,
				username,
				password_hash
			) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	GetUserByIDQuery = `
		SELECT 
			id, 
			name, 
			surname,
			bio, 
			profile_image, 
			username, 
			password_hash, 
			created_at, 
			updated_at
		FROM 
			users
		WHERE 
			id = $1
		`

	GetUserWithLoginQuery = `
		SELECT 
			id, 
			name, 
			surname,
			bio, 
			profile_image, 
			username, 
			password_hash, 
			created_at, 
			updated_at
		FROM 
			users
		WHERE 
			username = $1
		`

	GetListUserQuery = `
		SELECT 
			id, 
			name, 
			surname,
			bio, 
			profile_image, 
			username, 
			password_hash, 
			created_at, 
			updated_at
		FROM 
			users
		`

	GetListUserQueryCount = `
		SELECT 	
			COUNT(*)
		FROM 
			users
		`

	UpdateUserQuery = `
		UPDATE 
			users
		SET 
			name = $1, 
			surname = $2,
			bio = $3, 
			profile_image = $4, 
			username = $5, 
			password_hash = $6, 
			updated_at = NOW()
		WHERE
			id = $7
		`

	DeleteUserQuery = `
		DELETE 
			FROM 
				users
		WHERE 
			id = $1
		`
)
