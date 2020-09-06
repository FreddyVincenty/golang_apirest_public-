package database

import (
	"database/sql"

	//comentario
	_ "github.com/go-sql-driver/mysql"
)

//ResultRows para enviar resultado de consultas select
type ResultRows struct {
	Result  string
	Message string
	Rows    *sql.Rows
}

var userdb string = "root"
var passuserdb string = "root"
var dbname string = "mydb"
var ipserver string = "127.0.0.1:3306"

//Exito indica todo ok
var Exito string = "ok"

//Fallo indica que la operacion consulta sql fallo
var Fallo string = "fail"
var err error

var db *sql.DB
var rows *sql.Rows
var command *sql.Rows

func opendb() {

	db, err = sql.Open("mysql", userdb+":"+passuserdb+"@tcp("+ipserver+")/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()
}

func closedb() {
	db.Close()
}

//Insertar insertar en db
func Insertar(sql string) string {
	opendb()
	var respuesta string = Exito

	command, err = db.Query(sql)
	if err != nil {
		respuesta = "Error:" + err.Error()
		//panic(err.Error())
	}

	if respuesta == Exito {
		defer command.Close()
	}

	closedb()
	return respuesta
}

//QuerySelect consulta select
func QuerySelect(sql string) ResultRows {
	opendb()
	var resultado ResultRows

	rows, err = db.Query(sql)
	if err != nil {
		//panic(err.Error())
		resultado = ResultRows{Fallo, err.Error(), rows}
	} else {
		resultado = ResultRows{Exito, Exito, rows}
	}
	closedb()
	return resultado
}
