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

func GetUserWebListMapId(id string) map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT t_web.* FROM t_web LEFT JOIN t_user_web ON t_web.id=t_user_web.web LEFT JOIN t_user ON t_user.id=t_user_web.user_id WHERE t_user.id=$1", id)
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebListID - ", err)
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

func GetUserWebListId(id string) string {
	rows := GetUserWebListMapId(id)
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
	var rows map[int](map[string]string)
	var err error
	if session == "Admin" {
		rows, err = database.ExecuteQuery("SELECT t_web.id AS web_id, t_web.* FROM t_web LEFT JOIN t_user_web ON t_web.id=t_user_web.web LEFT JOIN t_user ON t_user.id=t_user_web.user_id WHERE t_web.id=$1", id)
	} else {
		rows, err = database.ExecuteQuery("SELECT t_web.id AS web_id, t_web.* FROM t_web LEFT JOIN t_user_web ON t_web.id=t_user_web.web LEFT JOIN t_user ON t_user.id=t_user_web.user_id WHERE t_user.session_token=$1 AND t_web.id=$2", session, id)
	}
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
			fmt.Println("ERR : GetUserWebDetail - ", err)
		}
		return nil
	}
}

func GetUserHTMLPart(session, id, types string) string {
	rows, err := database.ExecuteQuery("SELECT t_html_part_of.html_part AS html_part_id, t_web.*, t_html_part_of.id FROM t_web LEFT JOIN t_user_web ON t_web.id=t_user_web.web LEFT JOIN t_user ON t_user.id=t_user_web.user_id LEFT JOIN t_web_html ON t_web_html.web = t_web.id LEFT JOIN t_html_part_of on t_web_html.html = t_html_part_of.html LEFT JOIN t_html_part ON t_html_part_of.html_part = t_html_part.id WHERE t_user.session_token=$1 AND t_web.id=$2 AND t_html_part.type = $3 ;", session, id, types)
	if len(rows) > 0 && err == nil {
		return rows[0]["html_part_id"]
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserHTMLPart - ", err)
		}
		return ""
	}
}

func CheckUserWeb(session, web_id string) bool {
	rows, err := database.ExecuteQuery("select * from t_user_web left join t_user on t_user.id=t_user_web.user_id WHERE t_user.session_token=$1 and t_user_web.web=$2", session, web_id)
	fmt.Println(len(rows))
	if len(rows) > 0 && err == nil {
		return true
	} else {
		if err != nil {
			fmt.Println("ERR : CheckUserWeb - ", err)
		}
		return false
	}
}
