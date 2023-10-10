package groups

import (
	"database/sql"
	"easyvpn/src/database"
	"easyvpn/src/groups/groups_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
)

func GetGroups() (*[]groups_dtos.Group, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM Groups")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var groups []groups_dtos.Group

	for rows.Next() {
		var group groups_dtos.Group

		err := rows.Scan(&group.ID, &group.Name, &group.MemberCount, &group.Enabled, &group.Enabled)

		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err // Return the error
	}

	return &groups, nil

}

func GetMemberships(groupId string) (*[]user_dtos.User, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT U.id, U.username, U.name, U.is_admin, U.enabled
		FROM GroupMembership AS GM
		JOIN Users AS U ON GM.userId = U.id
		WHERE GM.groupId = ?;
	`

	rows, err := db.Query(query, groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user_dtos.User

	for rows.Next() {
		var user user_dtos.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.IsAdmin, &user.Enabled); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil

}
