package db

import (
	"fmt"
)

func DeletePart(id string) error {
	_, err := ExecuteQuery("DELETE FROM t_html_part WHERE id=$1;", id)
	if err != nil {
		fmt.Println("Err : DeleteHTMLPart - ", err)
		return err
	} else {
		return err
	}
}
