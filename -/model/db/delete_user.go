package db

import (
	"fmt"
)

func DeleteUser(id string) error {
	_, err := ExecuteQuery("DELETE FROM t_user WHERE id=$1;", id)
	if err != nil {
		fmt.Println("Err : UpdateUserHTMLPart - ", err)
		return err
	} else {
		return err
	}
}
