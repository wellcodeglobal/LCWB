package html

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetHTMLPartListType(session, html_id, types string) (string, [][]string) {
	selected_id := GetUserWebPartSection(session, html_id, types)
	rows, err := database.ExecuteQuery("SELECT t_html_part.id, t_html_part.html_code, t_html_part.css_code, t_html_part.js_code FROM t_html_part LEFT JOIN t_html_part_type ON t_html_part.type = t_html_part_type.id WHERE t_html_part_type.id=$1;", types)
	if len(rows) > 0 && err == nil {
		var result [][]string
		for i := 0; i < len(rows); i++ {
			selected := ""
			if selected_id == rows[i]["id"] {
				selected = "selected"
			}
			result = append(result, []string{rows[i]["id"], rows[i]["html_code"], rows[i]["css_code"], rows[i]["js_code"], selected})
		}
		return selected_id, result
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebList - ", err)
		}
		return selected_id, nil
	}
}

func GetUserWebPartSection(session, id, types string) string {
	rows, err := database.ExecuteQuery("SELECT t_html_part_of.html_part AS html_part_id, t_web.*, t_html_part_of.id FROM t_web LEFT JOIN t_user_web ON t_web.id=t_user_web.web LEFT JOIN t_user ON t_user.id=t_user_web.user_id LEFT JOIN t_web_html ON t_web_html.web = t_web.id LEFT JOIN t_html_part_of on t_web_html.html = t_html_part_of.html LEFT JOIN t_html_part ON t_html_part_of.html_part = t_html_part.id WHERE t_user.session_token=$1 AND t_web.id=$2 AND t_html_part.type = $3 ;", session, id, types)
	if len(rows) > 0 && err == nil {
		return fmt.Sprint(rows[0]["html_part_id"])
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebPartSection - ", err)
		}
		return ""
	}
}
