package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", "eurus:eurus@tcp(localhost:3306)/eurus")
	checkErr(err)
	defer db.Close()

	var stmt *sql.Stmt

	//retrieve data
	stmt, err = db.Prepare(`SELECT seqseq, seqnam, curval, incval FROM seq_dta_t
	WHERE seqseq >= ?`)
	checkErr(err)
	defer stmt.Close()

	var (
		rows   *sql.Rows
		seqseq int
		seqnam int
		curval int
		incval int
	)
	rows, err = stmt.Query(1)
	checkErr(err)

	columns, err := rows.Columns()
	checkErr(err)
	for _, col := range columns {
		fmt.Printf("%s\t", col)
	}
	println()

	colTypes, err := rows.ColumnTypes()
	checkErr(err)
	for _, ct := range colTypes {
		fmt.Printf("%s\t", ct.ScanType().Name())
	}
	println()

	for rows.Next() {
		rows.Scan(&seqseq, &seqnam, &curval, incval)
		fmt.Printf("%d\t%d\t%d\t%d\n", seqseq, seqnam, curval, incval)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
