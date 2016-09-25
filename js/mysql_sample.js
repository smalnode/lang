var mysql = require('mysql');
var squel = require('squel');
var http = require('http');

var conn = mysql.createConnection({
    host: 'localhost',
    port: 3306,
    user: 'wdeqin',
    password: 'wdeqin'
});

conn.connect();

var s = squel.select()
        .field('db_nam, infnam, fedseq, fednam, fedtyp, feddcm')
        .field("date_format(upddtm, '%Y%m%d%H%i%s') as upddtm")
        .from('dev_pub.dev_inf_vew_t')
        .where("infnam = 'FDBRWPRFX1'")
        .where("db_nam = 'dev_d00'");

console.log(s.toString());
conn.query(s.toString(), function(err, rows, fields) {
    if (err) throw err;

    rows.forEach(function(elem, index, arr) {
        console.log(index, ":", elem);
    });
    conn.end();
});
