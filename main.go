package main

import (
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/nilerajput91/RPc/grpc-server-1/model"
	userHandler "github.com/nilerajput91/RPc/user/handler"
	userRepo "github.com/nilerajput91/RPc/user/repo"
	userUsecase "github.com/nilerajput91/RPc/user/usecase"
	"google.golang.org/grpc"
)

func main() {
	var port = "8888"
	db, err := gorm.Open("mysql", "nilesh:nilesh@123@tcp(127.0.0.1:3306)/grpcclientserver?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.Debug().AutoMigrate(model.UserDB{})
	server := grpc.NewServer()
	userRepo := userRepo.CreateUserRepoImpl(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)
	userHandler.CreateUserHandler(server, userUsecase)

	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server starting at Port :", port)
	log.Fatal(server.Serve(conn))

}
