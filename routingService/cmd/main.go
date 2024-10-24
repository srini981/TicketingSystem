package main

import (
	"github.com/gin-gonic/gin"
	"intelXlabs/routingService/pkg"
	"log"
	"net/http"
)

func main() {
	log.Println("starting routing service at port 8004")

	router := gin.New()

	log.Println("adding routes to routing service")

	addRoutes(router)

	log.Println("added routes to routing service")

	router.Use(gin.Logger())
	router.Run(":8004")

	log.Println("routing service started")
}

// all the handlers to handle all the incoming http requests for all micrservices
func addRoutes(route *gin.Engine) {
	route.Handle(http.MethodGet, "health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
		return
	})
	route.Handle(http.MethodPost, "Login", pkg.Login)
	route.Use(pkg.JwtAuthentication)
	route.Use(pkg.Authorize)
	route.Handle(http.MethodPost, "user", pkg.CreateUser)
	route.Handle(http.MethodGet, "user/:ID", pkg.GetUser)
	route.Handle(http.MethodGet, "user", pkg.GetAllusers)
	route.Handle(http.MethodDelete, "user/:ID", pkg.Deleteuser)
	route.Handle(http.MethodPut, "user", pkg.Updateuser)
	route.Handle(http.MethodPost, "ticket", pkg.CreateTicket)
	route.Handle(http.MethodGet, "ticket/:ID", pkg.GetTicket)
	route.Handle(http.MethodGet, "tickets/user/:ID", pkg.GetAllTicketsForUser)
	route.Handle(http.MethodPut, "ticket", pkg.UpdateTicket)
	route.Handle(http.MethodDelete, "ticket/:ID", pkg.DeleteTicket)
	route.Handle(http.MethodGet, "ticket/:ID/label/:label", pkg.GetTicketsByLabel)
	route.Handle(http.MethodGet, "ticket/:ID/priority/:priority", pkg.GetTicketsByPriority)
	route.Handle(http.MethodPost, "ticket/:ID/user/:userID", pkg.AssignTicketsForUser)
	route.Handle(http.MethodPost, "role", pkg.CreateRole)
	route.Handle(http.MethodGet, "role/:ID", pkg.GetRole)
	route.Handle(http.MethodPut, "role", pkg.UpdateRole)
	route.Handle(http.MethodDelete, "role/:ID", pkg.DeleteRole)
	route.Handle(http.MethodGet, "roles", pkg.GetAllRoles)
	route.Handle(http.MethodGet, "roles/User/:ID", pkg.GetAllUserRoles)
}
