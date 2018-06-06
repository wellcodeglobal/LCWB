package user

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetUserData(session string) map[string]interface{} {
	rows, err := database.ExecuteQuery("SELECT t_user.id, t_user.name, t_user.email, t_user.company, t_user.phone, t_user.address, t_role.name AS role FROM t_user LEFT JOIN t_role ON t_user.role=t_role.id WHERE t_user.session_token = $1;", session)
	if len(rows) > 0 && err == nil {
		data := make(map[string]interface{})
		data["id"] = rows[0]["id"]
		data["name"] = rows[0]["name"]
		data["email"] = rows[0]["email"]
		data["company"] = rows[0]["company"]
		data["phone"] = rows[0]["phone"]
		data["address"] = rows[0]["address"]
		data["role"] = rows[0]["role"]
		return data
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserData - ", err)
		}
		return nil
	}
}
