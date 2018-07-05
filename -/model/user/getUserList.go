package user

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetUserListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT t_user.id, t_user.name, t_user.email, t_user.company, t_user.phone, t_user.address, t_role.name AS role FROM t_user LEFT JOIN t_role ON t_user.role=t_role.id ORDER BY t_user.role,t_user.name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebList - ", err)
		}
		return nil
	}
}

func GetUserList() string {
	rows := GetUserListMap()
	if rows == nil {
		return ""
	} else {
		js_str := "["
		for i := 0; i < len(rows); i++ {
			js_str += "["
			js_str += rows[i]["id"] + ",'" + rows[i]["name"] + "','" + rows[i]["email"] + "','" + rows[i]["company"] + "','" + rows[i]["phone"] + "','" + rows[i]["address"] + "','" + rows[i]["role"]
			js_str += "']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}

func GetUserDetail(id string) map[string]interface{} {

	rows, err := database.ExecuteQuery("SELECT t_user.id, t_user.name, t_user.email, t_user.company, t_user.phone, t_user.address, t_role.name AS role FROM t_user LEFT JOIN t_role ON t_user.role=t_role.id WHERE t_user.id = $1", id)
	if len(rows) > 0 && err == nil {
		result := make(map[string]interface{})
		result["id"] = rows[0]["id"]
		result["name"] = rows[0]["name"]
		result["email"] = rows[0]["email"]
		result["company"] = rows[0]["company"]
		result["address"] = rows[0]["address"]
		result["role"] = rows[0]["role"]
		result["phone"] = rows[0]["phone"]
		return result
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserDetail - ", err)
		}
		return nil
	}
}
