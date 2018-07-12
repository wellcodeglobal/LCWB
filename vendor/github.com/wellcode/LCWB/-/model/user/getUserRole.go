package user

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetUserRoleMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT * FROM t_role ORDER BY name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserRoleList - ", err)
		}
		return nil
	}
}

func GetUserRole() [][]string {
	rows := GetUserRoleMap()
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
