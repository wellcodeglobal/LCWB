package sign

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
	utility "github.com/wellcode/LCWB/-/model/utility"
	//"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (bool, string) {
	rows, err := database.ExecuteQuery("SELECT id from t_user where email=$1 and password=$2", email, password)
	if len(rows) > 0 && err == nil {
		sessionToken := utility.GenerateToken(email)
		_, err := database.ExecuteQuery("update t_user set session_token=$1 where email=$2", sessionToken, email)
		if err == nil {
			return true, sessionToken
		} else {
			fmt.Println("Login - 1", err)
			return false, ""
		}
	}
	return false, ""
}
