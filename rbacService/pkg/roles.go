package pkg

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	dp "intelXlabs/dbService/proto"
	"intelXlabs/rbacService/proto"
	"log"
	"time"
)

type RbactService struct {
	proto.RbacServiceServer
}

func (s *RbactService) CreateRole(ctx context.Context, in *proto.RoleRequest) (*proto.RoleResponse, error) {

	log.Println("CreateRole method called form rbac service")

	role := &dp.RoleRequest{
		Name:        in.Name,
		Description: in.Description,
		Timestamp:   time.Now().String(),
		Action:      in.Action,
		Id:          in.Id,
		CreatedBy:   in.CreatedBy,
		Users:       in.Users,
	}

	log.Println("calling CreateRole method in dbservice")

	_, err := client.DbServiceClient.CreateRole(ctx, role)

	if err != nil {
		log.Println("calling CreateRole method failed:", err)
		return nil, err
	}

	log.Println("got response from CreateRole method in dbservice")
	return &proto.RoleResponse{}, nil
}

func (s *RbactService) GetRole(ctx context.Context, in *proto.RoleID) (*proto.RoleResponse, error) {

	log.Println("GetRole method called form rbac service")

	role := &dp.RoleID{
		ID: in.ID,
	}

	log.Println("calling GetRole method in dbservice")

	resRole, err := client.DbServiceClient.GetRole(ctx, role)

	if err != nil {
		log.Println("calling GetRole method failed:", err)
		return nil, err
	}

	log.Println("got response from GetRole method in dbservice")

	response := &proto.RoleResponse{
		Name:        resRole.Name,
		Action:      resRole.Action,
		Description: resRole.Description,
		Timestamp:   resRole.Timestamp,
		Id:          resRole.Id,
		Users:       resRole.Users,
		CreatedBy:   resRole.CreatedBy,
	}

	return response, nil
}

func (s *RbactService) UpdateRole(ctx context.Context, in *proto.RoleRequest) (*proto.RoleResponse, error) {

	log.Println("UpdateRole method called form rbac service")

	role := &dp.RoleRequest{
		Name:        in.Name,
		Description: in.Description,
		Timestamp:   in.Timestamp,
		Action:      in.Action,
		Id:          in.Id,
		Users:       in.Users,
		CreatedBy:   in.CreatedBy,
	}

	log.Println("calling UpdateRole method in dbservice")

	_, err := client.DbServiceClient.UpdateRole(ctx, role)

	if err != nil {
		log.Println("calling UpdateRole method failed:", err)
		return nil, err
	}

	log.Println("got response from UpdateRole method in dbservice")
	return &proto.RoleResponse{}, nil
}

func (s *RbactService) DeleteRole(ctx context.Context, in *proto.RoleID) (*emptypb.Empty, error) {
	log.Println("DeleteRole method called form rbac service")

	role := &dp.RoleID{
		ID: in.ID,
	}

	log.Println("calling DeleteRole method in dbservice")

	_, err := client.DbServiceClient.DeleteRole(ctx, role)

	if err != nil {
		log.Println("calling DeleteRole method failed:", err)
		return nil, err
	}

	log.Println("got response from DeleteRole method in dbservice")
	return nil, nil
}

func (s *RbactService) AssignRole(ctx context.Context, in *proto.AssignUser) (*proto.RoleID, error) {
	log.Println("AssignRole method called form rbac service")

	log.Println("calling AssignRole method in dbservice")
	roleDetails := &dp.AssignUser{
		UserID: in.UserID,
		RoleId: in.RoleId,
	}

	roleid, err := client.DbServiceClient.AssignRole(ctx, roleDetails)

	if err != nil {
		log.Println("calling AssignRole method failed:", err)
		return nil, err
	}

	response := &proto.RoleID{ID: roleid.ID}

	log.Println("got response from AssignRole method in dbservice")
	return response, nil
}

func (s *RbactService) UnAssignRole(ctx context.Context, in *proto.AssignUser) (*proto.RoleID, error) {
	log.Println("UnAssignRole method called form rbac service")

	log.Println("calling UnAssignRole method in dbservice")
	roleDetails := &dp.AssignUser{
		UserID: in.UserID,
		RoleId: in.RoleId,
	}

	roleid, err := client.DbServiceClient.AssignRole(ctx, roleDetails)

	if err != nil {
		log.Println("calling UnAssignRole method failed:", err)
		return nil, err
	}

	response := &proto.RoleID{ID: roleid.ID}

	log.Println("got response from UnAssignRole method in dbservice")
	return response, nil
}

func (s *RbactService) GetAllRoles(ctx context.Context, in *emptypb.Empty) (*proto.Roles, error) {
	log.Println("GetAllRoles method called form rbac service")

	log.Println("calling GetAllRoles method in dbservice")

	resRoles, err := client.DbServiceClient.GetAllRoles(ctx, &emptypb.Empty{})

	if err != nil {
		log.Println("calling GetAllRoles method failed:", err)
		return nil, err
	}

	log.Println("got response from GetAllRoles method in dbservice")

	response := &proto.Roles{}
	for _, userrole := range resRoles.Roles {
		role := proto.RoleRequest{
			Name:        userrole.Name,
			Description: userrole.Description,
			Action:      userrole.Action,
			Id:          userrole.Id,
			Timestamp:   userrole.Timestamp,
			Users:       userrole.Users,
			CreatedBy:   userrole.CreatedBy,
		}
		response.Roles = append(response.Roles, &role)
	}

	return response, nil
}
func (s *RbactService) GetAllUserRoles(ctx context.Context, in *proto.AssignUser) (*proto.Roles, error) {
	log.Println("GetAllUserRoles method called form rbac service")

	log.Println("calling GetAllUserRoles method in dbservice")

	user := &dp.AssignUser{
		RoleId: in.RoleId,
		UserID: in.UserID,
	}

	resRoles, err := client.DbServiceClient.GetAllUserRoles(ctx, user)

	if err != nil {
		log.Println("calling GetAllUserRoles method failed:", err)
		return nil, err
	}

	log.Println("got response from GetAllUserRoles method in dbservice")

	response := &proto.Roles{}
	for _, userrole := range resRoles.Roles {
		role := proto.RoleRequest{
			Name:        userrole.Name,
			Description: userrole.Description,
			Action:      userrole.Action,
			Id:          userrole.Id,
			Timestamp:   userrole.Timestamp,
			Users:       userrole.Users,
			CreatedBy:   userrole.CreatedBy,
		}
		response.Roles = append(response.Roles, &role)
	}

	return response, nil
}
