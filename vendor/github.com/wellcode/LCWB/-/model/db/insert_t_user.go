package db

import (
	"fmt"
)

func InsertUser(name, email, phone, address, company, role, password string) (error, string) {
	msg := ""
	rows, err := ExecuteQuery("SELECT * from t_user WHERE email=$1;", email)
	if err == nil {
		if len(rows) > 0 {
			msg = "Error! Email used by another person. Please Change another email!"
			return err, msg
		} else {
			_, err = ExecuteQuery("INSERT INTO t_user(name, email, password, phone, address, company, role) VALUES ($1,$2,$3,$4,$5,$6,$7);", name, email, password, phone, address, company, role)
			if err != nil {
				fmt.Println("Err : InsertUser - ", err)
			}
			msg = "Create user \"" + email + "\"" + " done! refreshing this page."
			return err, msg
		}
	} else {
		if err != nil {
			fmt.Println("Err : InsertUser - ", err)
		}
		return err, "Error! Please try again."
	}
}
