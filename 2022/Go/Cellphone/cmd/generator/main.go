package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const URL_ROOT = "http://localhost:8080/"
const FROM = 1_000_000_000
const TO = 1_001_000_000
const STEP = 100_000

const PROVIDER_QUERY = "INSERT INTO PROVIDER (NAME, TOTAL) VALUES (?, ?);"
const CELLPHONE_QUERY = "INSERT INTO CELLPHONE (PROVIDER_ID, NUMBER) VALUES "

func newProvider(tx *sql.Tx, name string, total int) error {
	_, err := tx.Exec(PROVIDER_QUERY, name, total)

	if err != nil {
		return err
	}

	return nil
}

func newCellphone(tx *sql.Tx, queryData string) error {
	_, err := tx.Exec(queryData)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := sql.Open("mysql", "celluser:cellpass123@tcp(127.0.0.1:3306)/CELLDB")

	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()

	defer func() {
		tx.Commit()
	}()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prefix := "Provider"

	sb := strings.Builder{}
	sb.Grow(1_000_000_000)
	sb.WriteString(CELLPHONE_QUERY)

	j, k, p := 0, 1, 1

out:
	for i := FROM; i < TO; {
		// fib
		r := k + j
		j = k
		k = r
		name := fmt.Sprintf("%s%d", prefix, p)
		err := newProvider(tx, name, j)
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			return
		}
		fmt.Println("Generating", j, "entries")
		for m := 0; m < j; m++ {
			if (i+1)%STEP == 0 {
				sb.WriteString(fmt.Sprintf("(%d, '%d')", p, i))
			} else {
				sb.WriteString(fmt.Sprintf("(%d, '%d'),", p, i))
			}
			if err != nil {
				fmt.Println(err.Error())
				tx.Rollback()
				return
			}
			i++
			if i%STEP == 0 {
				fmt.Println("Inserting...", sb.Len())
				sb.WriteString(";")
				err = newCellphone(tx, sb.String())

				sb.Reset()
				sb.WriteString(CELLPHONE_QUERY)

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				err := tx.Commit()

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				tx, err = db.Begin()

				if err != nil {
					fmt.Println(err.Error())
					return
				}
			}
			if i >= TO {
				break out
			}
		}
		p++
	}
}
