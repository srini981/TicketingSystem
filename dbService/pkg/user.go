package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"intelXlabs/dbService/internals"
	"intelXlabs/dbService/proto"
	"log"
	"time"
)

func (d *DbService) CreateUser(ctx context.Context, in *proto.UserRequest) (*proto.UserResponse, error) {
	log.Println("CreateUser method called in dbservice")

	user := internals.User{
		Name:      in.Name,
		Phone:     in.Phone,
		Email:     in.Email,
		EmpID:     in.EmpID,
		Timestamp: in.Timestamp,
		Roles:     in.Roles,
		Password:  in.Password,
		Company:   in.Company,
	}

	tx := dbClient.PostgresClient.Create(&user)

	if tx.Error != nil {
		log.Println("failed to create user in db ", tx.Error.Error())
		return nil, tx.Error
	}

	res := &proto.UserResponse{
		Email: user.Email,
		Id:    int32(user.ID),
	}
	fmt.Println("user created in db")
	return res, nil
}
func (d *DbService) GetUser(ctx context.Context, in *proto.UserID) (*proto.UserResponse, error) {
	log.Println("GetUser method called in dbservice")

	user := internals.User{}
	query := fmt.Sprintf("select * from users where id = %d;", in.ID)

	result := dbClient.RedisClient.Get(query)

	if result.Err() != nil {

		fmt.Println("key not in redis")
		tx := dbClient.PostgresClient.Raw(query).Scan(&user)

		if tx.Error != nil {
			log.Println("failed to fetch user detials from db ", tx.Error.Error())
			return nil, tx.Error
		}

		temp, err := json.Marshal(user)

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

		log.Println("got response from redis")
		err := json.Unmarshal([]byte(result.Val()), &user)
		if err != nil {
			log.Println("failed to unmarshal json object", err.Error())
			return nil, err
		}
	}

	ticketres := proto.UserResponse{
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		Roles:     user.Roles,
		Id:        int32(user.ID),
		Timestamp: user.Timestamp,
		EmpID:     user.EmpID,
	}

	return &ticketres, nil
}

func (d *DbService) GetAllusers(context.Context, *emptypb.Empty) (*proto.ListUserResponse, error) {
	log.Println("GetAllusers method called in dbservice")

	users := []internals.User{}
	query := fmt.Sprintf("select * from users;")
	resUsers := &proto.ListUserResponse{}

	result := dbClient.RedisClient.Get(query)
	if result.Err() != nil {
		fmt.Println("key not in redis")

		tx := dbClient.PostgresClient.Raw(query).Scan(&users)
		if tx.Error != nil {
			log.Println("failed to fetch user details from db", tx.Error.Error())
			return nil, tx.Error
		}

		temp, err := json.Marshal(users)

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
		err := json.Unmarshal([]byte(result.Val()), &users)
		if err != nil {
			log.Println("failed to unmarshal json object", err.Error())
			return nil, err
		}
	}

	for _, ticket := range users {
		ticketres := proto.UserResponse{
			Id:        int32(ticket.ID),
			Name:      ticket.Name,
			Email:     ticket.Email,
			EmpID:     ticket.EmpID,
			Timestamp: ticket.Timestamp,
			Phone:     ticket.Phone,
			Roles:     ticket.Roles,
		}
		resUsers.Users = append(resUsers.Users, &ticketres)
	}

	return resUsers, nil
}
func (d *DbService) Updateuser(ctx context.Context, in *proto.UserRequest) (*emptypb.Empty, error) {
	log.Println("Updateuser method called in dbservice")

	user := internals.User{
		ID:        int(in.Id),
		Name:      in.Name,
		Email:     in.Email,
		EmpID:     in.EmpID,
		Timestamp: in.Timestamp,
		Phone:     in.Phone,
		Roles:     in.Roles,
		Password:  in.Password,
		Company:   in.Company,
	}

	tx := dbClient.PostgresClient.Model(internals.User{}).Where("id = ?", user.ID).Updates(&user)

	if tx.Error != nil {
		log.Println("failed to update user details in db ", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil
}
func (d *DbService) Deleteuser(ctx context.Context, in *proto.UserID) (*emptypb.Empty, error) {
	log.Println("Deleteuser method called in dbservice")

	tx := dbClient.PostgresClient.Model(internals.User{}).Where("id = ?", in.ID).Delete(&internals.User{})

	if tx.Error != nil {
		log.Println("failed to delete user from db", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil
}

func (d *DbService) GetUserByEmail(ctx context.Context, in *proto.UserEmail) (*proto.UserResponse, error) {
	log.Println("GetUserByEmail method called in dbservice")

	user := internals.User{}
	_ = dbClient.PostgresClient.Find(&user, "email = ?", in.Email)

	if user.Email == "" {
		log.Println("user not found in db")
		err := errors.New("user not found")
		return nil, err
	}

	ticketres := proto.UserResponse{
		Name:     user.Name,
		Email:    user.Email,
		Id:       int32(user.ID),
		Password: user.Password,
	}
	
	return &ticketres, nil
}
