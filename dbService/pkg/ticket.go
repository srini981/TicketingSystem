package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"intelXlabs/dbService/internals"
	"intelXlabs/dbService/proto"
	"log"
	"time"
)

type DbService struct {
	proto.DbServiceServer
}

func (d *DbService) CreateTicket(ctx context.Context, in *proto.TicketRequest) (*proto.TicketResponse, error) {
	log.Println("CreateTicket method called in dbservice")

	ticket := internals.Ticket{
		TicketName:  in.Name,
		Description: in.Description,
		AssignedTo:  int(in.AssignedTo),
		Priority:    int(in.Priority),
		Status:      in.Status,
		Timestamp:   time.Now().String(),
		Labels:      in.Labels,
		CreatedBy:   in.CreatedBy,
	}

	tx := dbClient.PostgresClient.Create(&ticket)

	if tx.Error != nil {
		log.Println("failed to create ticket in db ", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil
}
func (d *DbService) GetTicket(ctx context.Context, in *proto.TicketID) (*proto.TicketResponse, error) {
	log.Println("GetTicket method called in dbservice")

	ticket := internals.Ticket{}
	query := fmt.Sprintf("select * from tickets where id = %d;", in.ID)
	result := dbClient.RedisClient.Get(query)

	if result.Err() != nil {
		fmt.Println("key not in redis")
		tx := dbClient.PostgresClient.Raw(query).Scan(&ticket)

		if tx.Error != nil {
			log.Println("failed to get ticket in db ", tx.Error.Error())
			return nil, tx.Error
		}

		temp, err := json.Marshal(ticket)

		if err != nil {
			fmt.Println("failed to marshal json object", err)
			return nil, err
		}

		res := dbClient.RedisClient.Set(query, temp, 10*time.Minute)

		if res.Err() != nil {
			fmt.Println("failed to update in redis", res.Err().Error())
			return nil, res.Err()
		}

	} else {
		err := json.Unmarshal([]byte(result.Val()), &ticket)
		if err != nil {
			log.Println("failed to unmarshal response from redis")
			return nil, err
		}
	}

	ticketres := proto.TicketResponse{
		Name:        ticket.TicketName,
		Description: ticket.Description,
		AssignedTo:  int32(ticket.AssignedTo),
		Priority:    int32(ticket.Priority),
		Status:      ticket.Status,
		Timestamp:   ticket.Timestamp,
		CreatedBy:   ticket.CreatedBy,
	}

	return &ticketres, nil
}

func (d *DbService) GetAllTicketsForUser(ctx context.Context, in *proto.TicketID) (*proto.ListTicketResponse, error) {
	log.Println("GetAllTicketsForUser method called in dbservice")
	tickets := []internals.Ticket{}
	resTickets := proto.ListTicketResponse{}
	query := fmt.Sprintf("select * from tickets where assigned_to = %d;", in.ID)
	result := dbClient.RedisClient.Get(query)

	if result.Err() != nil {
		fmt.Println("key not in redis")
		tx := dbClient.PostgresClient.Raw(query).Scan(&tickets)

		if tx.Error != nil {
			log.Println("failed to fetch tickets for user", tx.Error.Error())
			return nil, tx.Error
		}

		temp, err := json.Marshal(tickets)

		if err != nil {
			fmt.Println("failed to marshal json object", err)
			return nil, err
		}

		res := dbClient.RedisClient.Set(query, temp, 10*time.Minute)

		if res.Err() != nil {
			fmt.Println("failed to update in redis", res.Err().Error())
			return nil, res.Err()
		}
	} else {
		err := json.Unmarshal([]byte(result.Val()), &tickets)
		if err != nil {
			log.Println("failed to unmarshal response from redis", err)
			return nil, err
		}
	}

	for _, ticket := range tickets {
		ticketres := proto.TicketResponse{
			ID:          int32(ticket.ID),
			Name:        ticket.TicketName,
			Description: ticket.Description,
			AssignedTo:  int32(ticket.AssignedTo),
			Priority:    int32(ticket.Priority),
			Status:      ticket.Status,
			Timestamp:   ticket.Timestamp,
			Labels:      ticket.Labels,
			CreatedBy:   ticket.CreatedBy,
		}
		resTickets.Tickets = append(resTickets.Tickets, &ticketres)
	}

	return &resTickets, nil
}
func (d *DbService) UpdateTicket(ctx context.Context, in *proto.TicketRequest) (*emptypb.Empty, error) {
	log.Println("UpdateTicket method called in dbservice")

	ticket := internals.Ticket{
		ID:          int(in.ID),
		TicketName:  in.Name,
		Description: in.Description,
		AssignedTo:  int(in.AssignedTo),
		Priority:    int(in.Priority),
		Status:      in.Status,
		Timestamp:   in.Timestamp,
		CreatedBy:   in.CreatedBy,
	}

	tx := dbClient.PostgresClient.Model(internals.Ticket{}).Where("id = ?", in.ID).Updates(&ticket)

	if tx.Error != nil {
		log.Println("failed to update ticket in db", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil

}
func (d *DbService) DeleteTicket(ctx context.Context, in *proto.TicketID) (*emptypb.Empty, error) {
	log.Println("DeleteTicket method called in dbservice")

	tx := dbClient.PostgresClient.Model(internals.Ticket{}).Where("id = ?", in.ID).Delete(&internals.Ticket{})

	if tx.Error != nil {
		log.Println("failed to delete ticket in db", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil

}
func (d *DbService) GetTicketsByLabel(ctx context.Context, in *proto.TicketLabel) (*proto.ListTicketResponse, error) {
	log.Println("GetTicketsByLabel method called in dbservice")

	tickets := []internals.Ticket{}
	query := fmt.Sprintf("select * from tickets,unnest(labels) as ulabel where ulabel = '%s';", in.Label)
	result := dbClient.RedisClient.Get(query)

	if result.Err() != nil {
		fmt.Println("key not in redis")
		tx := dbClient.PostgresClient.Raw(query).Scan(&tickets)
		if tx.Error != nil {
			log.Println("failed to get tickets by label from db", tx.Error.Error())
			return nil, tx.Error
		}

		temp, err := json.Marshal(tickets)

		if err != nil {
			fmt.Println("failed to marshal json object", err)
			return nil, err
		}

		res := dbClient.RedisClient.Set(query, temp, 10*time.Minute)
		if res.Err() != nil {
			fmt.Println("failed to update in redis", res.Err().Error())
			return nil, res.Err()
		}

	} else {
		err := json.Unmarshal([]byte(result.Val()), &tickets)

		if err != nil {
			log.Println("unable to marshall redis response", err.Error())
			return nil, err
		}
	}
	resTickets := proto.ListTicketResponse{}
	for _, ticket := range tickets {
		ticketres := proto.TicketResponse{
			ID:          int32(ticket.ID),
			Name:        ticket.TicketName,
			Description: ticket.Description,
			AssignedTo:  int32(ticket.AssignedTo),
			Priority:    int32(ticket.Priority),
			Status:      ticket.Status,
			Timestamp:   ticket.Timestamp,
			Labels:      ticket.Labels,
			CreatedBy:   ticket.CreatedBy,
		}
		resTickets.Tickets = append(resTickets.Tickets, &ticketres)
	}

	return &resTickets, nil
}
func (d *DbService) GetTicketsByPriority(ctx context.Context, in *proto.TicketLabel) (*proto.ListTicketResponse, error) {
	log.Println("GetTicketsByPriority method called in dbservice")

	tickets := []internals.Ticket{}
	query := fmt.Sprintf("select * from tickets where assigned_to= %d and  priority = '%s' ;", in.UserID, in.Label)
	result := dbClient.RedisClient.Get(query)

	if result.Err() != nil {
		fmt.Println("key not in redis")
		tx := dbClient.PostgresClient.Raw(query).Scan(&tickets)

		if tx.Error != nil {
			log.Println("unable to query data from db", tx.Error.Error())
			return nil, tx.Error
		}

		temp, err := json.Marshal(tickets)

		if err != nil {
			fmt.Println("failed to marshal json object", err)
			return nil, err
		}

		res := dbClient.RedisClient.Set(query, temp, 10*time.Minute)

		if res.Err() != nil {
			fmt.Println("failed to update in redis", res.Err().Error())
			return nil, res.Err()
		}

	} else {
		err := json.Unmarshal([]byte(result.Val()), &tickets)

		if err != nil {
			log.Println("unable to marshall redis response", err.Error())
			return nil, err
		}

	}

	resTickets := proto.ListTicketResponse{}
	for _, ticket := range tickets {
		ticketres := proto.TicketResponse{
			ID:          int32(ticket.ID),
			Name:        ticket.TicketName,
			Description: ticket.Description,
			AssignedTo:  int32(ticket.AssignedTo),
			Priority:    int32(ticket.Priority),
			Status:      ticket.Status,
			Timestamp:   ticket.Timestamp,
			Labels:      ticket.Labels,
			CreatedBy:   ticket.CreatedBy,
		}
		resTickets.Tickets = append(resTickets.Tickets, &ticketres)
	}

	return &resTickets, nil

}
func (d *DbService) AssignTicketToUser(ctx context.Context, in *proto.AssignTicket) (*emptypb.Empty, error) {
	log.Println("AssignTicketToUser method called in dbservice")

	tx := dbClient.PostgresClient.Model(internals.Ticket{}).Where("id = ?", in.ID).UpdateColumn("assigned_to", in.UserID)

	if tx.Error != nil {
		log.Println("failed to update user details in ticket", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil
}
