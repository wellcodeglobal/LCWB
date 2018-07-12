package html

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
)

func GetAllHTMLPartListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select t_html_part.*,coalesce(t_html_part.html_code,'') as html, coalesce(t_html_part.css_code,'') as css, coalesce(t_html_part.js_code,'') as js, t_html_part_type.name as part_type from t_html_part left join t_html_part_type on t_html_part_type.id=t_html_part.type;")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("Error - GetAllHTMLPartList : ", err)
		}
		return nil
	}
}

func GetAllHTMLPartList() string {
	rows := GetAllHTMLPartListMap()
	if rows == nil {
		return ""
	} else {
		js_str := "["
		for i := 0; i < len(rows); i++ {
			js_str += "["
			js_str += rows[i]["id"] + ",'" + rows[i]["name"] + "','" + rows[i]["part_type"] + "','" + rows[i]["price"] + "','<textarea rows=\"4\" cols=\"30\" style=\"resize: none;font-size: 10px;\" readonly oncontextmenu=\"this.focus();this.select()\">" + rows[i]["html"] + "</textarea>','<textarea rows=\"4\" cols=\"30\" style=\"resize: none;font-size: 10px;\" readonly oncontextmenu=\"this.focus();this.select()\">" + rows[i]["css"] + "</textarea>','<textarea rows=\"4\" cols=\"30\" style=\"resize: none;font-size: 10px;\" readonly oncontextmenu=\"this.focus();this.select()\">" + rows[i]["js"] + "</textarea>"
			js_str += "','<button onclick=\"detailPart(this)\" class=\"btn btn-info\">Detail</button>&nbsp&nbsp&nbsp<button onclick=\"editPart(this)\" class=\"btn btn-warning\">Edit</button>&nbsp&nbsp&nbsp<button onclick=\"deletePart(this)\" class=\"btn btn-danger\">Delete</button>']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}
