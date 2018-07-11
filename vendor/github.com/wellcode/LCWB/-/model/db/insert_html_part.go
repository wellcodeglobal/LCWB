package db

import ()

func InsertDivPart(types int, html_code, css_code, js_code string) bool {
	_, err := ExecuteQuery("INSERT INTO t_html (type,html_code,css_code,js_code) VALUES ($1,$2,$3,$4);", types, html_code, css_code, js_code)
	if err == nil {
		return true
	} else {
		return false
	}
}
