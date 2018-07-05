package db

import (
	"strconv"
)

func InsertWeb(base_cost, maintenance, service int, domain string) (int, bool) {
	rows, err := ExecuteQuery("INSERT INTO t_web (base_cost, maintenance, service, domain) VALUES ($1,$2,$3,$4) returning id;", base_cost, maintenance, service, domain)
	if err == nil {
		id, _ := strconv.Atoi(rows[0]["id"])
		return id, true
	} else {
		return 0, false
	}
}
