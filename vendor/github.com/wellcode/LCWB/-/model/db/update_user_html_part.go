package db

import (
	"fmt"
)

func UpdateUserHTMLPart(html_id, html_part_id, previous string) {
	// fmt.Println("-----------------")
	// fmt.Println("HTML      : " + html_id)
	// fmt.Println("HTML part : " + html_part_id)
	// fmt.Println("Previous  : " + previous)
	rows, err := ExecuteQuery("SELECT id FROM t_html_part_of WHERE html=$1 AND html_part=$2", html_id, previous)
	if len(rows) > 0 && err == nil {
		id := rows[0]["id"]
		ExecuteQuery("UPDATE t_html_part_of SET html_part=$1 WHERE id=$2;", html_part_id, id)
	} else {
		if err != nil {
			fmt.Println("Err : UpdateUserHTMLPart - ", err)
		}
	}
}
