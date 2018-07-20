package utility

import (
	database "github.com/wellcode/LCWB/-/model/db"
)

func CheckVerified(usertoken string) bool {
	rows, err := database.ExecuteQuery("SELECT email FROM users WHERE registration_token=$1 AND verified=$2 AND approval_token is null", usertoken, "false")
	if len(rows) > 0 && err == nil {
		newToken := GenerateToken(usertoken)
		_, err := database.ExecuteQuery("update users set approval_token=$1 ,status=$2 where registration_token=$3", newToken, "Need Approve", usertoken)
		if err == nil {
			return true
		}
	}
	return false
}
