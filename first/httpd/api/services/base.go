package services

import (
	"first/httpd/api/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB     *gorm.DB
	ROUTER *mux.Router
}

func (server *Server) Init() {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "paceecommercedb")
	server.DB, err = gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database")
	}

	server.ROUTER = mux.NewRouter()

	server.setup_router()
}

func (server *Server) setup_router() {
	server.ROUTER.HandleFunc("/user/register", middleware.SetMiddlewareJSON(server.Register)).Methods("POST")
	//router.POST("/user/login", user.Login)
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.ROUTER))
}
