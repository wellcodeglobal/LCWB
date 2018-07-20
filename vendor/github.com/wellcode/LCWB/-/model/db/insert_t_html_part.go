package db

import (
	"fmt"
)

func InsertHTMLPart(name, price, types, html, css, js string) (error, string) {
	msg := ""
	rows, err := ExecuteQuery("SELECT * from t_html_part WHERE html_code=$1;", html)
	if err == nil {
		if len(rows) > 0 {
			msg = "Error! Duplicate HTML code with\"" + rows[0]["name"] + "\"!"
			return err, msg
		} else {
			_, err = ExecuteQuery("INSERT INTO t_html_part(name, price, type, html_code, css_code, js_code) VALUES ($1,$2,$3,$4,$5,$6);", name, price, types, html, css, js)
			if err != nil {
				fmt.Println("Err : InsertHTMLPart - ", err)
			}
			msg = "Create HTML part \"" + name + "\"" + " done! refreshing this page."
			return err, msg
		}
	} else {
		if err != nil {
			fmt.Println("Err : InsertHTMLPart - ", err)
		}
		return err, "Error! Please try again."
	}
}
