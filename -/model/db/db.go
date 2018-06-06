package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	database "github.com/wellcode/LCWB/-/config"
)

var Execute *sql.DB

type info struct {
	adapter  string
	user     string
	password string
	sslmode  string
	dbname   string
}

var (
	infoDb = info{database.DB_ADAPTER, database.DB_USER, database.DB_PASSWORD, database.DB_SSLMODE, database.DB_NAME}
	dbinfo = "user=" + infoDb.user + " password=" + infoDb.password + " dbname=" + infoDb.dbname + " sslmode=" + infoDb.sslmode
)

func Conn() *sql.DB {
	conn, err := sql.Open(infoDb.adapter, dbinfo)
	if err != nil {
		panic(err)
	}
	Execute = conn
	return conn
}

func Connect() {
	conn, err := sql.Open(infoDb.adapter, dbinfo)
	if err != nil {
		panic(err)
	}
	Execute = conn
	fmt.Println("Open Database Success ...")
}

func ExecuteQuery(state string, params ...interface{}) (map[int](map[string]string), error) {
	result := make(map[int](map[string]string))
	rows, err := Execute.Query(state, params...)
	if err == nil {
		cols, err := rows.Columns()
		if err == nil {
			rowsID := 0

			for rows.Next() {
				columns := make([]interface{}, len(cols))
				columnPointers := make([]interface{}, len(cols))
				for i, _ := range columns {
					columnPointers[i] = &columns[i]
				}

				if err := rows.Scan(columnPointers...); err != nil {
					return result, err
				}
				rowscan := make(map[string]string)
				for i, colName := range cols {
					val := columnPointers[i].(*interface{})
					rowscan[colName] = fmt.Sprint(*val)
				}
				result[rowsID] = rowscan
				rowsID++
			}
			return result, nil
		} else {
			fmt.Println(err)
			return result, err
		}
	} else {
		fmt.Println(err)
		return result, err
	}

}
