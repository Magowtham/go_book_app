package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Magowtham/go_book_app/book_app"
	"github.com/Magowtham/go_book_app/db"
	"github.com/Magowtham/go_book_app/repository"
	"github.com/Magowtham/go_book_app/types"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type bookService struct {
	db types.DataBase
	book_app.UnimplementedBookAppBackendServer
}

func createBookService(db types.DataBase) *bookService {
	return &bookService{
		db: db,
	}
}

func (b *bookService) SignUpService(context context.Context, signUpRequest *book_app.SignUp) (*book_app.SignUpStatus, error) {

	fmt.Println(signUpRequest)
	user := types.CreateUser(signUpRequest.UserName, signUpRequest.Email, signUpRequest.Password)
	if err := b.db.CreateUser(user); err != nil {
		return nil, err
	}

	response := &book_app.SignUpStatus{
		Status: true,
	}
	return response, nil

}

func init() {
	log.Println("loading env from file...")
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listenAddress, ok := os.LookupEnv("SERVER_ADDRESS")

	if !ok {
		log.Fatalln("LISTEN_ADDRESS not found!!")
	}
	listener, err := net.Listen("tcp", listenAddress)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("gRPC server is listening at: %v\n", listenAddress)

	monogodbConn := db.Connect()

	mongodbRepo := repository.NewMongoDB(monogodbConn)

	gRpcServer := grpc.NewServer()

	bookService := createBookService(mongodbRepo)

	book_app.RegisterBookAppBackendServer(gRpcServer, bookService)

	if err := gRpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
