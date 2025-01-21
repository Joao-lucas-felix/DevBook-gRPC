package users_service

import (
	"log"

	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/gRPC/server"
)

func BootstrapService() {
	if err := server.BootstrapServer(); err != nil{
		log.Fatalln("error while trying to bootstrap the users service: ", err)
	}
}
