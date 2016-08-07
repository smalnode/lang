package main

import (
    _ "github.com/Go-SQL-Driver/MySQL"
    "database/sql";
    "fmt";
    //"time"
)

func main() {
    db, err := sql.Open("mysql", "wdeqin:7972@/test?charset=utf8")
    checkErr(err)

    /*
    //Insert data
    stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
    for i := 'A'; i <= 'Z'; i++ {
        res, err := stmt.Exec(string(i), "Research Dept.", "2013-10-24")
        checkErr(err)
        affect, err := res.RowsAffected()
        checkErr(err)
        fmt.Println(affect)
    }
    */

    /*
    fmt.Println(id)

    //Update data
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("wdeqinupdate", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)
    */

    //retrieve data
    rows, err := db.Query("SELECT userinfo.username, userdetail.intro, userdetail.profile FROM userinfo, userdetail where userinfo.uid = userdetail.uid")
    checkErr(err)

    for rows.Next() {
        // var uid int
        var username string
        // var department string
        // var created string
        var text string
        var profile string
        err = rows.Scan(&username, &text, &profile)
        checkErr(err)
        //fmt.Println(uid)
        fmt.Println(username)
        //fmt.Println(department)
        //fmt.Println(created)
        fmt.Println(text)
        fmt.Println(profile)
        fmt.Println()
    }

    /*
    // delete data
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect,err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)
    */

    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
