package main

import (
	"apirest/serv"
)

func main() {

	//serv.OpenConnectionBDGorm()
	serv.StartServer("8000")

}
