package html

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetUserWebPart(session, id string) [][]string {
	rows, err := database.ExecuteQuery("SELECT t_html.id AS html_id,t_html_part.id AS html_part_id,t_html_part.name AS html_part_name,t_html_part.html_code AS part_html_html,t_html_part.css_code AS part_html_css,t_html_part.js_code AS part_html_js, t_html_part_type.name AS type, t_html_part_type.id FROM t_html LEFT JOIN t_html_part_of ON t_html.id = t_html_part_of.html LEFT JOIN t_html_part ON t_html_part.id = t_html_part_of.html_part LEFT JOIN t_html_part_type ON t_html_part_type.id = t_html_part.type LEFT JOIN t_user ON t_user.id = t_html.user_id LEFT JOIN t_web_html ON t_web_html.html = t_html.id WHERE t_user.session_token = $1 AND t_web_html.web = $2 ORDER BY t_html_part_type.id ASC", session, id)

	if len(rows) > 0 && err == nil {
		var result [][]string
		for i := 0; i < len(rows); i++ {
			result = append(result, []string{rows[i]["html_id"], rows[i]["html_part_id"], rows[i]["html_part_name"], rows[i]["part_html_html"], rows[i]["part_html_css"], rows[i]["part_html_js"], rows[i]["type"]})
		}
		return result
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebPart - ", err)
		}
		return nil
	}
}

func GetUserWebPartAdmin(id string) [][]string {
	rows, err := database.ExecuteQuery("SELECT t_html.id AS html_id,t_html_part.id AS html_part_id,t_html_part.name AS html_part_name,t_html_part.html_code AS part_html_html,t_html_part.css_code AS part_html_css,t_html_part.js_code AS part_html_js, t_html_part_type.name AS type, t_html_part_type.id FROM t_html LEFT JOIN t_html_part_of ON t_html.id = t_html_part_of.html LEFT JOIN t_html_part ON t_html_part.id = t_html_part_of.html_part LEFT JOIN t_html_part_type ON t_html_part_type.id = t_html_part.type LEFT JOIN t_user ON t_user.id = t_html.user_id LEFT JOIN t_web_html ON t_web_html.html = t_html.id WHERE t_web_html.web = $1 ORDER BY t_html_part_type.id ASC", id)

	if len(rows) > 0 && err == nil {
		var result [][]string
		for i := 0; i < len(rows); i++ {
			result = append(result, []string{rows[i]["html_id"], rows[i]["html_part_id"], rows[i]["html_part_name"], rows[i]["part_html_html"], rows[i]["part_html_css"], rows[i]["part_html_js"], rows[i]["type"]})
		}
		return result
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebPart - ", err)
		}
		return nil
	}
}

func GetUserWebPartCode(id string) (string, string, string) {
	rows, err := database.ExecuteQuery("SELECT COALESCE(t_html_part.html_code,'') AS html_code,COALESCE(t_html_part.css_code,'') AS css_code,COALESCE(t_html_part.js_code,'') AS js_code, t_html_part_type.id FROM t_html LEFT JOIN t_html_part_of ON t_html.id = t_html_part_of.html LEFT JOIN t_html_part ON t_html_part.id = t_html_part_of.html_part LEFT JOIN t_html_part_type ON t_html_part_type.id = t_html_part.type LEFT JOIN t_user ON t_user.id = t_html.user_id LEFT JOIN t_web_html ON t_web_html.html = t_html.id WHERE t_web_html.web = $1 ORDER BY t_html_part_type.id ASC", id)

	if len(rows) > 0 && err == nil {
		var js string
		var index string
		var css string
		for i := 0; i < len(rows); i++ {
			if rows[i]["html_code"] != "" {
				index += fmt.Sprintln(rows[i]["html_code"])
			}
			if rows[i]["css_code"] != "" {
				css += fmt.Sprintln(rows[i]["css_code"])
			}
			if rows[i]["js_code"] != "" {
				js += fmt.Sprintln(rows[i]["js_code"])
			}
		}
		return index, css, js
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebPart - ", err)
		}
		return "", "", ""
	}
}
