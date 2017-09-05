package main

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbmap := initDB()
	defer dbmap.Db.Close()

	var seqs []SeqDta

	_, err := dbmap.Select(&seqs, "SELECT seqseq, seqnam, curval, incval FROM seq_dta_t")
	checkErr(err)

	for _, seq := range seqs {
		fmt.Println(seq)
	}

}

type SeqDta struct {
	SeqSeq int    `db:"seqseq"`
	SeqNam string `db:"seqnam,size:10"`
	CurVal int    `db:"curval"`
	IncVal int    `db:"incval"`
}

func initDB() *gorp.DbMap {
	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", "eurus:eurus@tcp(localhost:3306)/eurus")
	checkErr(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	dbmap.AddTableWithName(SeqDta{}, "seq_dta_t").SetKeys(true, "seqseq")

	return dbmap
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
