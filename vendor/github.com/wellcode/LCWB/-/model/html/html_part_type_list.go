package html

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetHTMLPartTypeMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT * FROM t_html_part_type ORDER BY name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserRoleList - ", err)
		}
		return nil
	}
}

func GetHTMLPartType() [][]string {
	rows := GetHTMLPartTypeMap()
	var result [][]string
	if rows == nil {
		return result
	} else {
		for i := 0; i < len(rows); i++ {
			id := fmt.Sprint(rows[i]["id"])
			name := fmt.Sprint(rows[i]["name"])
			temp := []string{id, name}
			result = append(result, temp)
		}
		return result
	}
}
