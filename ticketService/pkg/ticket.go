package pkg

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/types/known/emptypb"
	dp "intelXlabs/dbService/proto"
	"intelXlabs/ticketService/proto"
	"log"
	"strconv"
	"time"
)

// kafka topic to which the msg will be sent to
var topic = "tickets"

type TicketService struct {
	proto.TicketServiceServer
}

type kafkaMsg struct {
	TicketID    int
	Description string
	Message     string
}

func (s *TicketService) CreateTicket(ctx context.Context, in *proto.TicketRequest) (*proto.TicketResponse, error) {
	log.Println("CreateTicket method called")

	dbticket := &dp.TicketRequest{
		Name:        in.Name,
		Description: in.Description,
		AssignedTo:  in.AssignedTo,
		Priority:    in.Priority,
		Status:      in.Status,
		Timestamp:   time.Now().String(),
		Labels:      in.Labels,
		CreatedBy:   in.CreatedBy,
	}

	log.Println("calling CreateTicket method in dbservice")

	_, err := client.DbServiceClient.CreateTicket(ctx, dbticket)

	if err != nil {
		log.Println("calling CreateTicket method failed:", err)
		return nil, err
	}

	log.Println("got response from CreateTicket method in dbservice")

	kafkaMsg := kafkaMsg{
		TicketID:    int(in.Id),
		Message:     "Ticket created successfully",
		Description: "msg to notifiy the creation of ticket with ticked id" + strconv.Itoa(int(in.Id)),
	}

	msg, err := json.Marshal(kafkaMsg)

	if err != nil {
		log.Println("unable to marshall json object for kafka", err.Error())
		return nil, err
	}

	log.Println("Pushing message to kafka")

	err = PushMsgToQueue(topic, msg)

	if err != nil {
		log.Println("calling PushMessageToQueue method failed:", err)
		return nil, err
	}

	return &proto.TicketResponse{}, nil
}

func (s *TicketService) GetTicket(ctx context.Context, id *proto.TicketID) (*proto.TicketResponse, error) {

	log.Println("GetTicket method called form ticket service")

	dbticket := &dp.TicketID{
		ID: id.ID,
	}

	log.Println("calling GetTicket method in dbservice")

	dbres, err := client.DbServiceClient.GetTicket(ctx, dbticket)

	if err != nil {
		log.Println("calling GetTicket method failed:", err)

		return nil, err
	}

	response := &proto.TicketResponse{
		ID:          dbres.ID,
		Name:        dbres.Name,
		AssignedTo:  dbres.AssignedTo,
		Priority:    dbres.Priority,
		Description: dbres.Description,
		Status:      dbres.Status,
		Timestamp:   dbres.Timestamp,
		CreatedBy:   dbres.CreatedBy,
	}

	log.Println("got response from GetTicket method in dbservice")

	return response, nil
}
func (s *TicketService) GetAllTicketsForUser(ctx context.Context, in *proto.TicketID) (*proto.ListTicketResponse, error) {

	log.Println("GetTicket method called form ticket service")

	response := &proto.ListTicketResponse{}
	dbticket := &dp.TicketID{
		ID: in.ID,
	}

	log.Println("calling GetAllTicketsForUser method in dbservice")

	tickets, err := client.DbServiceClient.GetAllTicketsForUser(ctx, dbticket)

	if err != nil {
		log.Println("calling GetAllTicketsForUser method failed:", err)
		return nil, err
	}

	for _, userTicket := range tickets.Tickets {
		ticket := &proto.TicketResponse{ID: userTicket.ID, Name: userTicket.Name, Description: userTicket.Description, AssignedTo: userTicket.AssignedTo, Priority: userTicket.Priority, Timestamp: userTicket.Timestamp, Status: userTicket.Status, CreatedBy: userTicket.CreatedBy}
		response.Tickets = append(response.Tickets, ticket)
	}

	log.Println("got response from GetAllTicketsForUser method in dbservice")

	return response, nil
}
func (s *TicketService) UpdateTicket(ctx context.Context, ticket *proto.TicketRequest) (*emptypb.Empty, error) {

	log.Println("UpdateTicket method called form ticket service")

	dbticket := &dp.TicketRequest{
		ID:          ticket.Id,
		Name:        ticket.Name,
		Description: ticket.Description,
		AssignedTo:  ticket.AssignedTo,
		Priority:    ticket.Priority,
		Status:      ticket.Status,
		Timestamp:   ticket.Timestamp,
		CreatedBy:   ticket.CreatedBy,
	}

	log.Println("calling UpdateTicket method in dbservice")

	_, err := client.DbServiceClient.UpdateTicket(ctx, dbticket)

	if err != nil {
		log.Println("calling UpdateTicket method failed:", err)
		return nil, err
	}

	log.Println("got response from UpdateTicket method in dbservice")

	kafkaMsg := kafkaMsg{
		TicketID:    int(ticket.Id),
		Message:     "Ticket updated successfully",
		Description: "msg to notifiy the updation of ticket with ticked id" + strconv.Itoa(int(ticket.Id)),
	}

	msg, err := json.Marshal(kafkaMsg)

	if err != nil {
		log.Println("unable to marshall json object for kafka", err.Error())
		return nil, err
	}

	log.Println("Pushing message to kafka")

	err = PushMsgToQueue(topic, msg)

	if err != nil {
		log.Println("calling PushMessageToQueue method failed:", err)
		return nil, err
	}

	return nil, nil
}
func (s *TicketService) DeleteTicket(ctx context.Context, ticket *proto.TicketID) (*emptypb.Empty, error) {

	log.Println("delete ticket method called form ticket service")

	dbticket := &dp.TicketID{
		ID: ticket.ID}

	log.Println("calling DeleteTicket method in dbservice")

	_, err := client.DbServiceClient.DeleteTicket(ctx, dbticket)

	if err != nil {
		log.Println("calling DeleteTicket method failed:", err)
		return nil, err
	}

	log.Println("got response from DeleteTicket method in dbservice")

	kafkaMsg := kafkaMsg{
		TicketID:    int(ticket.ID),
		Message:     "Ticket UpdateTicket successfully",
		Description: "msg to notifiy the Deletion of ticket with ticked id" + strconv.Itoa(int(ticket.ID)),
	}

	msg, err := json.Marshal(kafkaMsg)

	if err != nil {
		log.Println("unable to marshall json object for kafka", err.Error())
		return nil, err
	}

	log.Println("Pushing message to kafka")

	err = PushMsgToQueue(topic, msg)

	if err != nil {
		log.Println("calling PushMessageToQueue method failed:", err)
		return nil, err
	}

	return nil, nil
}
func (s *TicketService) GetTicketsByLabel(ctx context.Context, in *proto.TicketLabel) (*proto.ListTicketResponse, error) {

	log.Println("GetTicketsByLabel method called form ticket service")

	dbticket := &dp.TicketLabel{
		UserID: in.UserID,
		Label:  in.Label,
	}

	log.Println("calling GetTicketsByLabel method in dbservice")

	tickets, err := client.DbServiceClient.GetTicketsByLabel(ctx, dbticket)

	if err != nil {
		log.Println("calling GetTicketsByLabel method failed:", err)
		return nil, err
	}

	response := &proto.ListTicketResponse{}

	for _, userTicket := range tickets.Tickets {
		ticket := &proto.TicketResponse{CreatedBy: userTicket.CreatedBy, ID: userTicket.ID, Name: userTicket.Name, Description: userTicket.Description, AssignedTo: userTicket.AssignedTo, Priority: userTicket.Priority, Timestamp: userTicket.Timestamp, Status: userTicket.Status}
		response.Tickets = append(response.Tickets, ticket)
	}

	log.Println("got response from GetTicketsByLabel method in dbservice")

	return response, nil
}
func (s *TicketService) GetTicketsByPriority(ctx context.Context, label *proto.TicketLabel) (*proto.ListTicketResponse, error) {
	log.Println("GetTicketsByPriority method called form ticket service")

	dbticket := &dp.TicketLabel{
		UserID: label.UserID,
		Label:  label.Label,
	}

	log.Println("calling GetTicketsByPriority method in dbservice")

	tickets, err := client.DbServiceClient.GetTicketsByPriority(ctx, dbticket)

	if err != nil {
		log.Println("calling GetTicketsByPriority method failed:", err)
		return nil, err
	}

	response := &proto.ListTicketResponse{}

	for _, userTicket := range tickets.Tickets {
		ticket := &proto.TicketResponse{CreatedBy: userTicket.CreatedBy,
			ID: userTicket.ID, Name: userTicket.Name, Description: userTicket.Description, AssignedTo: userTicket.AssignedTo, Priority: userTicket.Priority, Timestamp: userTicket.Timestamp, Status: userTicket.Status}
		response.Tickets = append(response.Tickets, ticket)
	}

	log.Println("got response from GetTicketsByPriority method in dbservice")

	return response, nil
}

func (s *TicketService) AssignTicketToUser(ctx context.Context, in *proto.AssignTicket) (*emptypb.Empty, error) {
	log.Println("AssignTicketToUser method called form ticket service")

	dbticket := &dp.AssignTicket{
		UserID: in.UserID,
		ID:     in.ID,
	}

	log.Println("calling AssignTicketToUser method in dbservice")

	_, err := client.DbServiceClient.AssignTicketToUser(ctx, dbticket)

	if err != nil {
		log.Println("calling AssignTicketToUser method failed:", err)
		return nil, err
	}

	log.Println("got response from AssignTicketToUser method in dbservice")

	return nil, nil
}
