package repository

func (ur *UserRepository) Register(id, name, email, hashedPassword string, createdAt, updatedAt time.Time) error {
	createSQL := `
	    INSERT INTO users (id, name, email, password, created_at, updated_at)
	    VALUES (?, ?, ?, ?, CONVERT_TZ(?, '+00:00', '+07:00'), CONVERT_TZ(?, '+00:00', '+07:00'))
	`

	_, err := ur.db.Exec(createSQL, id, name, email, hashedPassword, createdAt.UTC(), updatedAt.UTC())
	return err
}