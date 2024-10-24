package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"intelXlabs/routingService/internals"
	tp "intelXlabs/ticketService/proto"
	"log"
	"strconv"

	"io/ioutil"
	"net/http"
)

// handler for creating tickets which sends request to ticket service
func CreateTicket(c *gin.Context) {
	ticket := tp.TicketRequest{}
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println("unable to read json body")
		Response := internals.Response{Err: err, Message: "unable to read json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = json.Unmarshal(reqBody, &ticket)
	userid, check := c.Get("id")

	if !check {
		log.Println("unable to find  user id")
		Response := internals.Response{Err: err, Message: "unable to find  user id"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling create ticket using  ticket client")

	ticket.CreatedBy = int32(userid.(uint))

	_, err = client.TicketClient.CreateTicket(ctx, &ticket)

	if err != nil {
		log.Println("unable to create ticket using ticket client")
		Response := internals.Response{Err: err, Message: "unable to create ticket"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("created ticket successfully using ticket client ")

	Response := internals.Response{Err: nil, Message: "ticket created successfully"}
	c.JSON(http.StatusOK, Response)
}

// handler for fetching ticket based in ticket id which sends request to ticket service
func GetTicket(c *gin.Context) {
	ticketID, exists := c.Params.Get("ID")

	if !exists {
		log.Println("ticket ID not found in path")
		Response := internals.Response{Err: errors.New("ticket ID required"), Message: "ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(ticketID)

	if err != nil {
		log.Println("invalid ticket ID  found in path")
		Response := internals.Response{Err: err, Message: "invalid ticket ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &tp.TicketID{ID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get ticket using  ticket client")

	ticket, err := client.TicketClient.GetTicket(ctx, getrequest)

	if err != nil {
		log.Println("failed to get ticket using ticket client ", err.Error())
		Response := internals.Response{Err: err, Message: "unable to get ticket"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("fetched ticket successfully using ticket client ")

	Response := internals.Response{Err: nil, Message: " successfully fetched ticket", Response: ticket}
	c.JSON(http.StatusOK, Response)
}

// handler for fetching all  ticket based on user id which sends request to ticket service
func GetAllTicketsForUser(c *gin.Context) {
	ticketID, exists := c.Params.Get("ID")

	if !exists {
		log.Println("user ID not found in path")

		Response := internals.Response{Err: errors.New("user ID required"), Message: "user ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(ticketID)

	if err != nil {
		log.Println("invalid user ID", err.Error())
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &tp.TicketID{ID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get all tickets for user using  ticket client")

	tickets, err := client.TicketClient.GetAllTicketsForUser(ctx, getrequest)

	if err != nil {
		log.Println("failed to  get all tickets for user using  ticket client ", err.Error())

		Response := internals.Response{Err: nil, Message: "failed to get tickets for user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("successfully fetched all tickets for user using  ticket client ")

	Response := internals.Response{Err: nil, Message: " successfully fetched tickets for user", Response: tickets}
	c.JSON(http.StatusOK, Response)
}

// handler for updating ticket which sends request to ticket service
func UpdateTicket(c *gin.Context) {
	ticket := tp.TicketRequest{}
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println("unable to read json body")
		Response := internals.Response{Err: err, Message: "unable to read json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = json.Unmarshal(reqBody, &ticket)
	if err != nil {
		log.Println("unable to parse json body")
		Response := internals.Response{Err: err, Message: "unable to parse json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling update ticket using  ticket client")

	_, err = client.TicketClient.UpdateTicket(ctx, &ticket)

	if err != nil {
		Response := internals.Response{Err: err, Message: "unable to update ticket"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	Response := internals.Response{Err: nil, Message: "ticket updated successfully"}
	c.JSON(http.StatusOK, Response)
}

// handler for deleting tickets which sends request to ticket service
func DeleteTicket(c *gin.Context) {
	ticketID, exists := c.Params.Get("ID")

	if !exists {
		log.Println("ticket ID not found in path")

		Response := internals.Response{Err: errors.New("ticket ID required"), Message: "ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(ticketID)

	if err != nil {
		log.Println("invalid ticket ID  found in path")
		Response := internals.Response{Err: err, Message: "invalid ticket ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ticket := &tp.TicketID{ID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling delete ticket using  ticket client")

	_, err = client.TicketClient.DeleteTicket(ctx, ticket)

	if err != nil {
		Response := internals.Response{Err: err, Message: "unable to delete ticket"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	Response := internals.Response{Err: nil, Message: "ticket deleted successfully"}
	c.JSON(http.StatusOK, Response)
}

// handler for fetching all the tickets based on ticket label which sends request to ticket service
func GetTicketsByLabel(c *gin.Context) {
	ticketID, exists := c.Params.Get("ID")

	if !exists {
		log.Println("user id not found in path")

		Response := internals.Response{Err: errors.New("user ID required"), Message: "user ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(ticketID)

	if err != nil {
		log.Println("invalid user ID", err.Error())
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ticketLabel, exists := c.Params.Get("label")

	if !exists {
		log.Println("label not found in path")
		Response := internals.Response{Err: errors.New("user ID required"), Message: "user ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &tp.TicketLabel{UserID: int32(id), Label: ticketLabel}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get all tickets by label using  ticket client")

	tickets, err := client.TicketClient.GetTicketsByLabel(ctx, getrequest)

	if err != nil {
		log.Println("failed to  get  tickets for user using  ticket client ", err.Error())

		Response := internals.Response{Err: nil, Message: "failed to get tickets for user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("successfully fetched all tickets by label using  ticket client ")

	Response := internals.Response{Err: nil, Message: " successfully fetched tickets for label", Response: tickets}
	c.JSON(http.StatusOK, Response)
}

// handler for fetching all the tickets based on ticket priority which sends request to ticket service
func GetTicketsByPriority(c *gin.Context) {
	ticketID, exists := c.Params.Get("ID")

	if !exists {
		log.Println("user id not found in path")

		Response := internals.Response{Err: errors.New("user ID required"), Message: "user ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	priority, exists := c.Params.Get("priority")

	if !exists {
		log.Println("ticket priority id not found in path")

		Response := internals.Response{Err: errors.New("priority ID required"), Message: "ticket priority id not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(ticketID)

	if err != nil {
		log.Println("invalid user ID", err.Error())
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &tp.TicketLabel{UserID: int32(id), Label: priority}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get all tickets by priority using  ticket client")

	tickets, err := client.TicketClient.GetTicketsByPriority(ctx, getrequest)

	if err != nil {
		log.Println("failed to  get  tickets by user priority using  ticket client ", err.Error())

		Response := internals.Response{Err: nil, Message: "failed to get tickets for user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("successfully fetched all tickets by priority using  ticket client ")

	Response := internals.Response{Err: nil, Message: " successfully fetched tickets for label", Response: tickets}
	c.JSON(http.StatusOK, Response)
}

// handler for assigning tickets based to user which sends request to ticket service
func AssignTicketsForUser(c *gin.Context) {
	userid, exists := c.Params.Get("userID")

	if !exists {
		log.Println("user id not found in path")

		Response := internals.Response{Err: errors.New("user ID required"), Message: "user ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ticketid, exists := c.Params.Get("ID")

	if !exists {
		log.Println("ticket  id not found in path")

		Response := internals.Response{Err: errors.New("priority ID required"), Message: "ticket priority id not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	user, err := strconv.Atoi(userid)

	if err != nil {
		log.Println("invalid user ID", err.Error())
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ticket, err := strconv.Atoi(ticketid)

	if err != nil {
		log.Println("invalid ticket ID", err.Error())
		Response := internals.Response{Err: err, Message: "invalid ticket ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &tp.AssignTicket{UserID: int32(user), ID: int32(ticket)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling AssignTicketToUser using  ticket client")

	_, err = client.TicketClient.AssignTicketToUser(ctx, getrequest)

	if err != nil {
		log.Println("failed to assign tickets to user using  ticket client ", err.Error())

		Response := internals.Response{Err: nil, Message: "failed to assign tickets to user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("successfully assigned tickets to user using  ticket client ")

	Response := internals.Response{Err: nil, Message: " successfully assigned tickets to user using  ticket client"}
	c.JSON(http.StatusOK, Response)
}
