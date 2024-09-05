package repositories

import (
	"database/sql"
	"search-engine/internal/models"
)

func SearchUsers(db *sql.DB, searchTerm string) ([]models.User, error) {

	query := `
	SELECT id, name, username, icon, bio WHERE username LIKE ? OR name LIKE ?
	`

	rows, err := db.Query(query, "%"+searchTerm+"%", "%"+searchTerm+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err != rows.Scan(&user.ID, &user.Name, &user.Username, &user.Icon, user.Bio) {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil

}
