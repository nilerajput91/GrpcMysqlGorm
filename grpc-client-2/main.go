package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nilerajput91/RPc/grpc-client-2/model"
	userHandler "github.com/nilerajput91/RPc/grpc-client-2/user/handler"
	"google.golang.org/grpc"
)

func main() {
	port := "8083"
	targetPort := "8888"
	conn, err := grpc.Dial(":"+targetPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %v %v", targetPort, err)
	}
	user := model.NewUsersClient(conn)
	router := gin.Default()
	userHandler.CreateUserHandler(router, user)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
