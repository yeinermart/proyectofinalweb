package main

import (
	"log" //referencia al paquete de registro estándar
	"net/http"

	GorillaHandlers "github.com/gorilla/handlers"
	"github.com/yeinermart/proyecto/controllers"
	"github.com/yeinermart/proyecto/handlers"
	"github.com/yeinermart/proyecto/models"
	repositorio "github.com/yeinermart/proyecto/repository" /* importando el paquete de repositorio */

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

/*
función para conectarse a la instancia de PostgreSQL, en general sirve para cualquier base de datos SQL.
Necesita la URL del host donde está instalada la base de datos y el tipo de base datos (driver)
*/
func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	/* creando un objeto de conexión a PostgreSQL - le paso la URL de la instancia a elephant*/
	/*db almacena la conexion, err almacena si hubo error - la url y postgress es un controlador de postgresSQL para realizar la conexion */
	db, err := ConectarDB("postgres://xbrjoevu:e3kIS8s-HnaFjyt3dNsy4rOKt5-Lt2Hm@berry.db.elephantsql.com/xbrjoevu", "postgres")
	if err != nil { //si hubo un error entra en el condicional - db interatua con la base de datos
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/* creando una instancia del tipo Repository del paquete repository
	se debe especificar el tipo de struct que va a manejar la base de datos
	para este ejemplo es Amigo y se le pasa como parámetro el objeto de
	conexión a PostgreSQL - ejemplo de una instancia{ se crea una clase persona {nombre, edad}} persona1 y persona 2 son instancias*/
	repo, err := repositorio.NewRepository[models.Estudiante](db) //carpeta models y la struct estudiante - crea una instancia
	if err != nil {                                               //crea instancia del repositorio para el modelo Estudiante usando la funcion constructora NewRepo
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return //Es útil para manejar errores críticos y detener la ejecución cuando algo va mal.
	}

	controller, err := controllers.NewController(repo) //crear una nueva instancia de controlador utilizando el repositorio previamente creado (repo).
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := handlers.NewHandler(controller) //se utiliza para instanciar un manejador (handler) que probablemente se encargará de gestionar las solicitudes relacionadas con el controlador de Estudiante
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux)
	está llamando a la función NewRouter() del paquete mux, que devuelve un nuevo enrutador.
	 Este enrutador se utiliza para manejar las rutas y las solicitudes HTTP */
	router := mux.NewRouter() //implementa funciones que permite recortar codigo y hacer lo de abajo

	/* rutas a los endpoints de la API - donde se usan los metodos get=leer, post=crear, patch=actualizardeauno, delete=borrar*/
	router.Handle("/estudiantes", http.HandlerFunc(handler.LeerAmigos)).Methods(http.MethodGet)
	router.Handle("/estudiantes", http.HandlerFunc(handler.CrearAmigo)).Methods(http.MethodPost)
	router.Handle("/estudiantes/{id}", http.HandlerFunc(handler.LeerUnAmigo)).Methods(http.MethodGet)
	router.Handle("/estudiantes/{id}", http.HandlerFunc(handler.ActualizarUnAmigo)).Methods(http.MethodPatch)
	router.Handle("/estudiantes/{id}", http.HandlerFunc(handler.EliminarUnAmigo)).Methods(http.MethodDelete)

	headers := GorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := GorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"})
	origins := GorillaHandlers.AllowedOrigins([]string{"*"})

	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router */
	http.ListenAndServe(":8080", GorillaHandlers.CORS(headers, methods, origins)(router))
}
