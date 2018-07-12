package db

import (
	"fmt"
)

func UpdateUser(id, name, phone, address, company, email, password, role string) error {
	_, err := ExecuteQuery("UPDATE t_user SET name=$1, phone=$2, address=$3, company=$4 WHERE id=$5;", name, phone, address, company, id)
	if err != nil {
		fmt.Println("Err : UpdateUser - ", err)
	}
	return err
}
