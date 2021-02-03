package main

import (
	"github.com/jinzhu/gorm"
	"github.com/wwwillian/codepix-go/application/grpc"
	"github.com/wwwillian/codepix-go/infrastructure/db"
	"os"
)

var database *gorm.DB

func main()  {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
