package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type UserRepository struct {
	DB *sql.DB
}

/* --------------- functions that interact with with database --------------- */

// Insert new user in db
func (repo *UserRepository) Add(user models.User) error {
	// example code
	stmt, err := repo.DB.Prepare("INSERT INTO users(user_id, email,first_name, last_name, nickname, about, password, birthday, image) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(user.ID, user.Email, user.FirstName, user.LastName, user.Nickname, user.About, user.Password, user.DateOfBirth, user.ImagePath); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) EmailNotTaken(email string) (bool, error) {
	row := repo.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? LIMIT 1", email)
	var result int
	if err := row.Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
	}
	if result != 0 {
		return false, nil
	}
	return true, nil
}

// find user by email and return user_id and password / mainly for login funcionality
func (repo *UserRepository) FindUserByEmail(email string) (models.User, error) {
	row := repo.DB.QueryRow("SELECT user_id,password FROM users WHERE email = ? LIMIT 1", email)
	var user models.User
	if err := row.Scan(&user.ID, &user.Password); err != nil {
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

// Return list of all users except current and following/follower info
func (repo *UserRepository) GetAllAndFollowing(userID string) ([]models.User, error) {
	var users []models.User

	rows, err := repo.DB.Query("SELECT user_id, first_name || ' ' || last_name as full_name, (SELECT COUNT(*) FROM followers WHERE followers.user_id = "+userID+" AND follower_id = users.user_id) as follower,(SELECT COUNT(*) FROM followers WHERE followers.user_id = users.user_id AND follower_id = "+userID+") as following, nickname FROM users WHERE user_id != ? GROUP BY user_id;", userID)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user models.User
		var fullName string
		var follower int
		var following int
		rows.Scan(&user.ID, &fullName, &follower, &following, &user.Nickname)
		// Replace nickname with fumm name if nickname not provided
		if len(user.Nickname) == 0 {
			user.Nickname = fullName
		}
		// configure followers
		if follower == 1 {
			user.Follower = true
		}
		if following == 1 {
			user.Following = true
		}
		users = append(users, user)
	}
	return users, nil
}
