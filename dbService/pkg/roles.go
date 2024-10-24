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

func (d *DbService) CreateRole(ctx context.Context, in *proto.RoleRequest) (*proto.RoleResponse, error) {
	log.Println("createRole method called in dbservice")

	ticket := internals.Role{
		Name:        in.Name,
		Description: in.Description,
		Action:      in.Action,
		Timestamp:   in.Timestamp,
		Users:       in.Users,
		CreatedBy:   in.CreatedBy,
	}

	tx := dbClient.PostgresClient.Create(&ticket)

	if tx.Error != nil {
		log.Println("failed to create role", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil
}
func (d *DbService) GetRole(ctx context.Context, in *proto.RoleID) (*proto.RoleResponse, error) {
	log.Println("GetRole method called in dbservice")

	role := internals.Role{}
	query := fmt.Sprintf("select * from roles where id = %d;", in.ID)

	result := dbClient.RedisClient.Get(query)

	if result.Err() != nil {
		fmt.Println("key not in redis")

		tx := dbClient.PostgresClient.Raw(query).Scan(&role)

		if tx.Error != nil {
			return nil, tx.Error
		}

		temp, err := json.Marshal(role)

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

		err := json.Unmarshal([]byte(result.Val()), &role)

		if err != nil {
			return nil, err
		}
	}

	roleRes := proto.RoleResponse{
		Name:        role.Name,
		Description: role.Description,
		Action:      role.Action,
		Timestamp:   role.Timestamp,
		Users:       role.Users,
		CreatedBy:   role.CreatedBy,
	}

	return &roleRes, nil
}
func (d *DbService) UpdateRole(ctx context.Context, in *proto.RoleRequest) (*proto.RoleResponse, error) {
	log.Println("UpdateRole method called in dbservice")

	role := internals.Role{
		ID:          int(in.Id),
		Name:        in.Name,
		Description: in.Description,
		Action:      in.Action,
		Timestamp:   in.Timestamp,
		Users:       in.Users,
		CreatedBy:   in.CreatedBy,
	}

	tx := dbClient.PostgresClient.Model(internals.Role{}).Where("id = ?", in.Id).Updates(&role)

	if tx.Error != nil {
		log.Println("failed to update role", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil
}
func (d *DbService) DeleteRole(ctx context.Context, in *proto.RoleID) (*emptypb.Empty, error) {
	log.Println("DeleteRole method called in dbservice")

	tx := dbClient.PostgresClient.Model(internals.Role{}).Where("id = ?", in.ID).Delete(&internals.Role{})

	if tx.Error != nil {
		log.Println("failed to delete role", tx.Error.Error())
		return nil, tx.Error
	}

	return nil, nil
}

func (d *DbService) AssignRole(context.Context, *proto.AssignUser) (*proto.RoleID, error) {
	panic("implement me")
}
func (d *DbService) UnAssignRole(context.Context, *proto.AssignUser) (*proto.RoleID, error) {
	panic("implement me")
}

func (d *DbService) GetAllRoles(context.Context, *emptypb.Empty) (*proto.Roles, error) {
	log.Println("GetAllRoles method called in dbservice")

	roles := []internals.Role{}
	resRoles := &proto.Roles{}
	query := fmt.Sprintf("select * from roles;")
	result := dbClient.RedisClient.Get(query)

	if result.Err() != nil {
		fmt.Println("key not in redis")
		tx := dbClient.PostgresClient.Raw(query).Scan(&roles)

		if tx.Error != nil {
			return nil, tx.Error
		}

		temp, err := json.Marshal(roles)

		if err != nil {
			fmt.Println("failed to marshal json object", err)
			return nil, err
		}

		res := dbClient.RedisClient.Set(query, temp, 5*time.Minute)
		if res.Err() != nil {
			fmt.Println("failed to update in redis", res.Err().Error())
			return nil, res.Err()
		}
	} else {
		err := json.Unmarshal([]byte(result.Val()), &roles)

		if err != nil {
			return nil, err
		}

	}

	for _, role := range roles {
		tempRole := proto.RoleRequest{
			Name:        role.Name,
			Description: role.Description,
			Action:      role.Action,
			Id:          int32(role.ID),
			Timestamp:   role.Timestamp,
			Users:       role.Users,
			CreatedBy:   role.CreatedBy,
		}
		resRoles.Roles = append(resRoles.Roles, &tempRole)
	}

	return resRoles, nil
}
func (d *DbService) GetAllUserRoles(ctx context.Context, in *proto.AssignUser) (*proto.Roles, error) {
	log.Println("GetAllUserRoles method called in dbservice")

	roles := []internals.Role{}
	query := fmt.Sprintf("select * from roles,unnest(users) as uid where uid= %d;;", in.UserID)
	result := dbClient.RedisClient.Get(query)
	resRoles := proto.Roles{}

	if result.Err() != nil {
		fmt.Println("key not in redis")

		err := dbClient.PostgresClient.Raw(query).Scan(&roles).Error

		if err != nil {
			log.Println("failed to get roles from db", err.Error())
			return nil, err
		}

		temp, err := json.Marshal(roles)

		if err != nil {
			fmt.Println("failed to marshal json object", err)
			return nil, err
		}

		res := dbClient.RedisClient.Set(query, temp, 5*time.Minute)

		if res.Err() != nil {

			fmt.Println("failed to update in redis", res.Err().Error())
			return nil, res.Err()
		}
	} else {
		err := json.Unmarshal([]byte(result.Val()), &roles)
		if err != nil {
			log.Println("failed to unmarshal redis response", err.Error())
			return nil, err
		}
	}

	for _, role := range roles {
		roleres := proto.RoleRequest{
			Id:          int32(role.ID),
			Action:      role.Action,
			Name:        role.Name,
			Description: role.Description,
			Timestamp:   role.Timestamp,
			Users:       role.Users,
			CreatedBy:   role.CreatedBy,
		}
		resRoles.Roles = append(resRoles.Roles, &roleres)
	}

	return &resRoles, nil
}
