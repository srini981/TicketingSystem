package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	rp "intelXlabs/rbacService/proto"
	"intelXlabs/routingService/internals"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// handler for role creation which sends request to rbac service
func CreateRole(c *gin.Context) {
	role := rp.RoleRequest{}
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println("unable to read json body")
		Response := internals.Response{Err: err, Message: "unable to read json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = json.Unmarshal(reqBody, &role)

	if err != nil {
		log.Println("unable to parse json body")
		Response := internals.Response{Err: err, Message: "unable to parse json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ctx := context.Background()

	defer ctx.Done()

	userId, check := c.Get("id")

	if !check {
		log.Println("unable to find user id")
		Response := internals.Response{Err: err, Message: "unable to find user id"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	role.CreatedBy = int32(userId.(uint))

	log.Println("calling create role using  rbac client")

	_, err = client.RbacClient.CreateRole(ctx, &role)

	if err != nil {
		log.Println("unable to  create role using  rbac client")
		Response := internals.Response{Err: err, Message: "unable to create role"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("created role successfully using ticket client")

	Response := internals.Response{Err: nil, Message: "role created successfully"}
	c.JSON(http.StatusOK, Response)
}

// handler for fetching role based on role ID which sends request to rbac service
func GetRole(c *gin.Context) {
	roleid, exists := c.Params.Get("ID")

	if !exists {
		log.Println("user ID not found in path")
		Response := internals.Response{Err: errors.New("user ID required"), Message: "ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(roleid)

	if err != nil {
		log.Println("invalid user ID  found in path")
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &rp.RoleID{ID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get role using  rbac client")

	user, err := client.RbacClient.GetRole(ctx, getrequest)

	if err != nil {
		log.Println("failed to get role using rbac client ", err.Error())
		Response := internals.Response{Err: err, Message: "unable to get role"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("fetched role successfully using rbac client ")
	Response := internals.Response{Err: nil, Message: " successfully fetched user", Response: user}
	c.JSON(http.StatusOK, Response)
}

// handler for role updation which sends request to rbac service
func UpdateRole(c *gin.Context) {
	role := rp.RoleRequest{}
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println("unable to read json body")
		Response := internals.Response{Err: err, Message: "unable to read json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = json.Unmarshal(reqBody, &role)

	if err != nil {
		log.Println("unable to parse json body")
		Response := internals.Response{Err: err, Message: "unable to parse json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling update role using  rbac client")

	_, err = client.RbacClient.UpdateRole(ctx, &role)

	if err != nil {
		Response := internals.Response{Err: err, Message: "unable to update role"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	Response := internals.Response{Err: nil, Message: "role updated successfully"}
	c.JSON(http.StatusOK, Response)
}

// handler for role deletion which sends request to rbac service
func DeleteRole(c *gin.Context) {
	roleID, exists := c.Params.Get("ID")

	if !exists {
		log.Println("role ID not found in path")
		Response := internals.Response{Err: errors.New("role ID required"), Message: "ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(roleID)

	if err != nil {
		log.Println("invalid role ID  found in path")
		Response := internals.Response{Err: err, Message: "invalid role ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &rp.RoleID{ID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling delete role using  role client")

	_, err = client.RbacClient.DeleteRole(ctx, getrequest)

	if err != nil {
		log.Println("failed to delete role using rbac client ", err.Error())
		Response := internals.Response{Err: err, Message: "unable to delete role"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("deleted role successfully using rbac client")

	Response := internals.Response{Err: nil, Message: "successfully deleted role"}
	c.JSON(http.StatusOK, Response)
}

// handler for fethcing all roles which sends request to rbac service
func GetAllRoles(c *gin.Context) {

	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get roles using  rbac client")

	roles, err := client.RbacClient.GetAllRoles(ctx, &emptypb.Empty{})

	if err != nil {
		log.Println("failed to get rbac ", err.Error())
		Response := internals.Response{Err: err, Message: "unable to get rbac"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("fetched user roles successfully using rbac client ")

	Response := internals.Response{Err: nil, Message: " successfully fetched roles", Response: roles}
	c.JSON(http.StatusOK, Response)
}

// handler for fethcing all user roles which sends request to rbac service
func GetAllUserRoles(c *gin.Context) {
	userID, exists := c.Params.Get("ID")

	if !exists {
		log.Println("user ID not found in path")

		Response := internals.Response{Err: errors.New("user ID required"), Message: "user ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(userID)

	if err != nil {
		log.Println("invalid user ID", err.Error())
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &rp.AssignUser{UserID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get all user roles method using  rbac client")

	userRoles, err := client.RbacClient.GetAllUserRoles(ctx, getrequest)

	if err != nil {
		log.Println("failed to  get all roles for user using  rbac client ", err.Error())

		Response := internals.Response{Err: nil, Message: "failed to get roles for user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("successfully fetched all roles for user using  rbac client ")

	Response := internals.Response{Err: nil, Message: " successfully fetched tickets for user", Response: userRoles}
	c.JSON(http.StatusOK, Response)
}
