package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	// _ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	// _ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	// if Dbdriver == "postgres" {
	// 	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	// 	server.DB, err = gorm.Open(Dbdriver, DBURL)
	// 	if err != nil {
	// 		fmt.Printf("Cannot connect to %s database", Dbdriver)
	// 		log.Fatal("This is the error:", err)
	// 	} else {
	// 		fmt.Printf("We are connected to the %s database", Dbdriver)
	// 	}
	// }
	// if Dbdriver == "sqlite3" {
	// 	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	// 	server.DB, err = gorm.Open(Dbdriver, DbName)
	// 	if err != nil {
	// 		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
	// 		log.Fatal("This is the error:", err)
	// 	} else {
	// 		fmt.Printf("We are connected to the %s database\n", Dbdriver)
	// 	}
	// 	server.DB.Exec("PRAGMA foreign_keys = ON")
	// }

	// server.DB.Debug().AutoMigrate(&models.User{}, &models.Software{}, &models.VideoTutorial{}, &models.DokumenPendukung{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port", addr)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	// server.Router.Use(mux.CORSMethodMiddleware(server.Router))
	// server.Router.Use(testMiddleware(server.Router))

	// // log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(server.Router)))

	// log.Fatal(http.ListenAndServe(addr, server.Router))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "X-CSRF-Token", "X-Custom-Header", "Access-Control-Allow-Origin", "Authorization", "Accept"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3030"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(server.Router)))

	// CORSHandler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
	// 	AllowedHeaders:   []string{"X-Requested-With", "Content-Type"},
	// 	AllowCredentials: false,
	// })
	// server.Router.Use(CORSHandler)
}
