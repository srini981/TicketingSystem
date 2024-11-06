package main

import (
	"fmt"
	"intelXlabs/chatService/internals"
	handlers "intelXlabs/chatService/pkg"
	"intelXlabs/chatService/proto"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	godotenv.Load()

	fmt.Println(
		fmt.Sprintf("%s%s%s%s", "Server will start at http://localhost:8007"),
	)
	go startGrpcServer()
	internals.ConnectDatabase()

	route := mux.NewRouter()

	AddApproutes(route)

	serverPath := ":8007"

	cors := internals.GetCorsConfig()

	err := http.ListenAndServe(serverPath, cors.Handler(route))
	if err != nil {
		log.Fatalf("failed to start chatService", err.Error())
		return
	}

}
func startGrpcServer() {
	port := ":8008"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterChatServiceServer(s, &handlers.ChatService{})
	fmt.Println("started grpc server")
	if err = s.Serve(lis); err != nil {
		fmt.Println("failed to start grpc server :", err)
		return
	}
}

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {

	log.Println("Loading Routes...")

	hub := handlers.NewHub()
	go hub.Run()
	route.HandleFunc("/", handlers.RenderHome)
	route.HandleFunc("/isUsernameAvailable/{username}", handlers.IsUsernameAvailable)
	route.HandleFunc("/login", handlers.Login).Methods("POST")
	route.HandleFunc("/userSessionCheck/{userID}", handlers.UserSessionCheck)
	route.HandleFunc("/getConversation/{toUserID}/{fromUserID}", handlers.GetMessagesHandler)
	route.HandleFunc("/ws/{userID}", func(responseWriter http.ResponseWriter, request *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		// Reading username from request parameter
		userID := mux.Vars(request)["userID"]

		// Upgrading the HTTP connection socket connection
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		connection, err := upgrader.Upgrade(responseWriter, request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		handlers.CreateNewSocketUser(hub, connection, userID)

	})
	log.Println("Routes are Loaded.")
}
