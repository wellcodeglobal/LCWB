package db

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func UpdateHTMLCode(web_id, index, css, js string) {
	dat, err := ioutil.ReadFile("-/view/partial/base/html/index_base.html")
	if err == nil {
		temp_index := string(dat)
		temp_index = strings.Replace(temp_index, "{{.index}}", index, -1)
		_, err := ExecuteQuery("UPDATE t_html SET html_code=$1, css_code=$2, js_code=$3 WHERE id=$4;", temp_index, css, js, web_id)
		if err != nil {
			fmt.Println("Error UpdateHTMLCode1 : ", err)
		}
	} else {
		fmt.Println("Error - UpdateHTMLCode2 : ", err)
	}
}
