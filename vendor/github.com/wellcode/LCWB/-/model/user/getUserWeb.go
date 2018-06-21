package user

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetUserWebListMap(session string) map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT t_web.* FROM t_web LEFT JOIN t_user_web ON t_web.id=t_user_web.web LEFT JOIN t_user ON t_user.id=t_user_web.user_id WHERE t_user.session_token=$1", session)
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebList - ", err)
		}
		return nil
	}
}

func GetUserWebList(session string) string {
	rows := GetUserWebListMap(session)
	if rows == nil {
		return ""
	} else {
		js_str := "["
		for i := 0; i < len(rows); i++ {
			js_str += "["
			js_str += rows[i]["id"] + ",'" + rows[i]["domain"] + "'," + rows[i]["base_cost"] + "," + rows[i]["maintenance"] + "," + rows[i]["service"]
			js_str += "]"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}

func GetUserWebDetail(session, id string) map[string]interface{} {
	rows, err := database.ExecuteQuery("SELECT t_web.id AS web_id, t_web.* FROM t_web LEFT JOIN t_user_web ON t_web.id=t_user_web.web LEFT JOIN t_user ON t_user.id=t_user_web.user_id WHERE t_user.session_token=$1 AND t_web.id=$2", session, id)
	if len(rows) > 0 && err == nil {
		result := make(map[string]interface{})
		result["web_id"] = rows[0]["web_id"]
		result["domain"] = rows[0]["domain"]
		result["base_cost"] = rows[0]["base_cost"]
		result["service"] = rows[0]["service"]
		result["maintenance"] = rows[0]["maintenance"]
		return result
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebList - ", err)
		}
		return nil
	}
}
