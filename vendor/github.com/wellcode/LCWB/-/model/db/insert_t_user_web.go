package db

import ()

func InsertUserWeb(user, web int) bool {
	_, err := ExecuteQuery("INSERT INTO t_user_web (user_id, web) VALUES ($1,$2);", user, web)
	if err == nil {
		return true
	} else {
		return false
	}
}
