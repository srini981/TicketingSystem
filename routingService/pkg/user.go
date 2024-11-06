package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"intelXlabs/routingService/internals"
	tp "intelXlabs/userService/proto"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// function used to validate the emailid of the user
func Validate(email string, password string) string {

	if !strings.Contains(email, "@") {

		return "Email address is required"
	}

	if len(password) < 6 {

		return "Password must have atleast 6 characters"
	}

	details := &tp.UserEmail{Email: email}

	_, err := client.UserClient.GetUserByEmail(context.Background(), details)

	if err == nil {
		return "Email address already in use by another user."
	}

	return ""
}

// handler for logging the user in which request to user service
func Login(c *gin.Context) {
	user := internals.UserLogin{}
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println("unable to read json body")
		Response := internals.Response{Err: err, Message: "unable to read json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = json.Unmarshal(reqBody, &user)

	if err != nil {
		log.Println("unable to parse json body")
		Response := internals.Response{Err: err, Message: "unable to parse json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	userDetails, err := client.UserClient.GetUserByEmail(context.Background(), &tp.UserEmail{Email: user.Name})

	if err != nil {
		Response := internals.Response{Err: err, Message: "user not found invalid credentials"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		Response := internals.Response{Err: err, Message: "user not found invalid credentials"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	tk := &internals.Token{UserId: uint(userDetails.Id)}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, _ := token.SignedString([]byte(internals.Token_key))
	tokenResponse := internals.TokenResponse{UserId: uint(userDetails.Id), Name: user.Name, Token: tokenString}

	log.Println("user logged in successfully using user client ")

	Response := internals.Response{Err: nil, Message: "user logged in  successfully", Response: tokenResponse}
	c.JSON(http.StatusOK, Response)
}

// handler for creating user  which sends request to user service
func CreateUser(c *gin.Context) {
	user := tp.UserRequest{}
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println("unable to read json body")
		Response := internals.Response{Err: err, Message: "unable to read json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = json.Unmarshal(reqBody, &user)

	if err != nil {
		log.Println("unable to parse json body")
		Response := internals.Response{Err: err, Message: "unable to parse json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling create user using  user client")

	check := Validate(user.Email, user.Password)

	if check != "" {
		Response := internals.Response{Err: err, Message: check}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	resuser, err := client.UserClient.CreateUser(ctx, &user)
	fmt.Println(err)
	if err != nil {
		log.Println("unable to  create  user using  user client")

		Response := internals.Response{Err: err, Message: "unable to create user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	tk := &internals.Token{UserId: uint(resuser.Id)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(internals.Token_key))

	if err != nil {
		Response := internals.Response{Err: err, Message: "unable to create token for user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	tokenResponse := internals.TokenResponse{UserId: uint(resuser.Id), Name: user.Name, Token: tokenString}

	log.Println("created user successfully using ticket client ")

	Response := internals.Response{Err: nil, Message: "user created successfully", Response: tokenResponse}
	c.JSON(http.StatusOK, Response)
}

// handler for fetching  user details which sends request to user service
func GetUser(c *gin.Context) {

	userId, exists := c.Params.Get("ID")

	if !exists {
		log.Println("user ID not found in path")
		Response := internals.Response{Err: errors.New("user ID required"), Message: "ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		log.Println("invalid user ID  found in path")
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &tp.UserID{ID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling get user using  user client")
	user, err := client.UserClient.GetUser(ctx, getrequest)

	if err != nil {
		log.Println("failed to get user using user client ", err.Error())
		Response := internals.Response{Err: err, Message: "unable to get user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("fetched user successfully using user client ")
	Response := internals.Response{Err: nil, Message: " successfully fetched user", Response: user}
	c.JSON(http.StatusOK, Response)
}

// handler for fetching all the user details which sends request to user service
func GetAllusers(c *gin.Context) {

	ctx := context.Background()
	defer ctx.Done()
	log.Println("calling get users using  user client")

	users, err := client.UserClient.GetAllusers(ctx, &emptypb.Empty{})

	if err != nil {
		log.Println("failed to get users ", err.Error())
		Response := internals.Response{Err: err, Message: "unable to get users"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("fetched user successfully using user client ")

	Response := internals.Response{Err: nil, Message: " successfully fetched user", Response: users}
	c.JSON(http.StatusOK, Response)
}

// handler for updating user details which sends request to user service
func Updateuser(c *gin.Context) {
	user := tp.UserRequest{}
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println("unable to read json body")
		Response := internals.Response{Err: err, Message: "unable to read json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	err = json.Unmarshal(reqBody, &user)

	if err != nil {
		log.Println("unable to parse json body")
		Response := internals.Response{Err: err, Message: "unable to parse json body"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling update user using  user client")

	_, err = client.UserClient.Updateuser(ctx, &user)

	if err != nil {
		Response := internals.Response{Err: err, Message: "unable to update user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	Response := internals.Response{Err: nil, Message: "user updated successfully"}
	c.JSON(http.StatusOK, Response)
}

// handler for deleting user based on user id which sends request to user service
func Deleteuser(c *gin.Context) {
	userId, exists := c.Params.Get("ID")

	if !exists {
		log.Println("user ID not found in path")
		Response := internals.Response{Err: errors.New("user ID required"), Message: "ID not found in path"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		log.Println("invalid user ID  found in path")
		Response := internals.Response{Err: err, Message: "invalid user ID"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	getrequest := &tp.UserID{ID: int32(id)}
	ctx := context.Background()
	defer ctx.Done()

	log.Println("calling delete user using  user client")

	_, err = client.UserClient.Deleteuser(ctx, getrequest)

	if err != nil {
		log.Println("failed to delete user using user client ", err.Error())
		Response := internals.Response{Err: err, Message: "unable to delete user"}
		c.JSON(http.StatusBadRequest, Response)
		return
	}

	log.Println("deleted user successfully using user client ")

	Response := internals.Response{Err: nil, Message: " successfully deleted user"}
	c.JSON(http.StatusOK, Response)
}
