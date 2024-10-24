package pkg

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	cp "intelXlabs/chatService/proto"
	dp "intelXlabs/dbService/proto"

	"intelXlabs/userService/proto"
	"log"
	"time"
)

type UserService struct {
	proto.UserServiceServer
}

func (u *UserService) CreateUser(ctx context.Context, in *proto.UserRequest) (*proto.UserResponse, error) {
	log.Println("CreateUser ,ethod called form ticket service")

	user := &dp.UserRequest{
		Name:      in.Name,
		Phone:     in.Phone,
		Email:     in.Email,
		EmpID:     in.EmpID,
		Company:   in.Company,
		Roles:     in.Roles,
		Password:  in.Password,
		Timestamp: time.Now().String(),
	}

	log.Println("calling CreateUser method in dbservice")

	userResponse, err := client.DbServiceClient.CreateUser(ctx, user)

	if err != nil {
		log.Println("calling CreateUser method failed:", err)
		return nil, err
	}

	userDetails := &cp.UserDetails{
		UserName: in.Email,
		Password: in.Password,
	}

	_, err = client.ChatClient.AddUser(ctx, userDetails)

	if err != nil {
		log.Println("calling AddUser method failed:", err)
		return nil, err
	}

	response := &proto.UserResponse{
		Id:    userResponse.Id,
		Email: userResponse.Email,
	}

	log.Println("got response from CreateUser method in dbservice")
	return response, nil
}

func (u *UserService) GetUser(ctx context.Context, in *proto.UserID) (*proto.UserResponse, error) {
	log.Println("GetUser method called form ticket service")

	user := &dp.UserID{
		ID: in.ID,
	}

	log.Println("calling CreateUser method in dbservice")

	resUser, err := client.DbServiceClient.GetUser(ctx, user)

	if err != nil {
		log.Println("calling CreateUser method failed:", err)
		return nil, err
	}

	log.Println("got response from CreateUser method in dbservice")

	userResponse := &proto.UserResponse{Name: resUser.Name, Phone: resUser.Phone, Email: resUser.Email, EmpID: resUser.EmpID, Company: resUser.Company, Roles: resUser.Roles, Timestamp: resUser.Timestamp}
	return userResponse, nil
}

func (u *UserService) GetAllusers(ctx context.Context, in *emptypb.Empty) (*proto.ListUserResponse, error) {
	log.Println("GetUser method called form ticket service")

	log.Println("calling CreateUser method in dbservice")

	users, err := client.DbServiceClient.GetAllusers(ctx, &emptypb.Empty{})

	if err != nil {
		log.Println("calling GetAllusers method failed:", err)
		return nil, err
	}

	response := &proto.ListUserResponse{}
	for _, resUser := range users.Users {
		userResponse := &proto.UserResponse{Name: resUser.Name, Phone: resUser.Phone, Email: resUser.Email, EmpID: resUser.EmpID, Company: resUser.Company, Timestamp: resUser.Timestamp}
		response.Users = append(response.Users, userResponse)
	}

	log.Println("got response from GetAllusers method in dbservice")
	return response, nil
}

func (u *UserService) Updateuser(ctx context.Context, in *proto.UserRequest) (*emptypb.Empty, error) {
	log.Println("Updateuser method called form ticket service")

	user := &dp.UserRequest{
		Name:      in.Name,
		Phone:     in.Phone,
		Email:     in.Email,
		EmpID:     in.EmpID,
		Company:   in.Company,
		Timestamp: in.Timestamp,
		Password:  in.Password,
		Roles:     in.Roles,
		Id:        in.Id,
	}

	log.Println("calling Updateuser method in dbservice")

	_, err := client.DbServiceClient.Updateuser(ctx, user)

	if err != nil {
		log.Println("calling Updateuser method failed:", err)
		return nil, err
	}

	log.Println("got response from Updateuser method in dbservice")
	return nil, nil
}

func (u *UserService) Deleteuser(ctx context.Context, in *proto.UserID) (*emptypb.Empty, error) {

	log.Println("Deleteuser method called form ticket service")

	user := &dp.UserID{
		ID: in.ID,
	}

	log.Println("calling Deleteuser method in dbservice")

	_, err := client.DbServiceClient.Deleteuser(ctx, user)

	if err != nil {
		log.Println("calling Deleteuser method failed:", err)
		return nil, err
	}

	log.Println("got response from Deleteuser method in dbservice")
	return nil, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, in *proto.UserEmail) (*proto.UserResponse, error) {
	log.Println("GetUserByEmail method called form ticket service")

	user := &dp.UserEmail{
		Email: in.Email,
	}

	log.Println("calling GetUserByEmail method in dbservice")

	resUser, err := client.DbServiceClient.GetUserByEmail(ctx, user)

	if err != nil {
		log.Println("calling GetUserByEmail method failed:", err)
		return nil, err
	}
	
	userResponse := &proto.UserResponse{Id: resUser.Id, Name: resUser.Name, Email: resUser.Email, Password: resUser.Password}

	log.Println("got response from GetUserByEmail method in dbservice")
	return userResponse, nil
}
