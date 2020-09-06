package serv

//GITHUB PUBLICO
import (
	"apirest/database"
	"apirest/objects"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//StartServer iniciar el servidor // normalmente se usa el puerto 8000
func StartServer(port string) {

	//TOMAR ENCUENTA QUE LOS DATOS ?id=2 se conocen como consultas o querys para filtrar
	//EN CAMBIO USUARIOS, VENTAS, COMPRAS, PRODUCTOS, ETC SE CONOCEN COMO RECURSOS Y PARA ELLO SE USA VARIABLES DE RUTA EJEMPLO: /{id} QUE DEVUELVEN DIRECTAMENTE UN RECURSO EN ESPECIFICO

	router := mux.NewRouter()
	router.HandleFunc("/api", home)                              //uso ==> http://127.0.0.1:8000/api  //OJO AQUI SI MANDAS COMO GET,POST,PUT,DELETE HACE DIFERENTE COSA
	router.HandleFunc("/api/welcome", welcome)                   // uso ==> http://127.0.0.1:8000/api/welcome //solo muestra un simple mensaje de bienvenida
	router.HandleFunc("/api/users", getUsers).Methods("GET")     //uso ==> http://127.0.0.1:8000/api/users
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET") //uso ==> http://127.0.0.1:8000/api/users/1
	router.HandleFunc("/api/querys", withQuery).Methods("GET")   //uso ==> http://127.0.0.1:8000//api/querys?id=2&name=Fredd

	//log.Fatal muestra error en consola si hay algun problema
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bienvenido a mi APIREST! v 1.1")
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var users []objects.User

	key := r.FormValue("key")
	fmt.Println(key)

	resultado := database.QuerySelect("SELECT id, nombre FROM usuarios")
	if resultado.Result == database.Exito {
		for resultado.Rows.Next() {
			var user objects.User
			err := resultado.Rows.Scan(&user.ID, &user.Name)
			if err != nil {
				panic(err.Error())
			}
			users = append(users, user)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	} else {
		fmt.Println(resultado.Message)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resultado.Message)
	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var users []objects.User
	resultado := database.QuerySelect("SELECT id, nombre FROM usuarios WHERE id = " + params["id"])
	if resultado.Result == database.Exito {
		for resultado.Rows.Next() {
			var user objects.User
			err := resultado.Rows.Scan(&user.ID, &user.Name)
			if err != nil {
				panic(err.Error())
			}
			users = append(users, user)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	} else {
		fmt.Println(resultado.Message)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resultado.Message)
	}
}

func withQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.FormValue("id")
	name := r.FormValue("name")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("id:" + id + " name:" + name)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
