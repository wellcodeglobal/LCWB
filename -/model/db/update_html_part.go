package db

import (
	"fmt"
)

func UpdateHTMLPart(id, name, price, types, html, css, js string) error {
	_, err := ExecuteQuery("UPDATE t_html_part SET name=$1, price=$2, type=$3, css_code=$4, js_code=$5, html_code=$6 WHERE id=$7;", name, price, types, css, js, html, id)
	if err != nil {
		fmt.Println("Err : UpdateHTMLPart - ", err)
	}
	return err
}
