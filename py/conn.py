import mysql.connector
from mysql.connector import errorcode

config = {
    'user': 'smalu',
    'password': 'smalu',
    'host': '127.0.0.1',
    'database': 'cookbook',
    'raise_on_warnings': True,
}

try:
    cnx = mysql.connector.connect(**config)
    cur = cnx.cursor()

    queryStr = "SELECT cookbook.last_foot_location()" 
    cur.execute(queryStr)

    for location in cur:
    	print ("last_foot_location = %s\n") % (location)
except mysql.connector.Error as err:
    if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
    	print("Something is wrong with your user name or password")
    elif err.errno == errorcode.ER_BAD_DB_ERROR:
    	print("Database does not exists")
    else:
    	print(err)
else:
    cnx.close()
