package pkg

import (
	"context"
	"fmt"
	rb "intelXlabs/rbacService/proto"
	"intelXlabs/routingService/internals"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT middleware to check authentication using JWT
func JwtAuthentication(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization") //Grab the token from the header

	if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
		response := "Missing auth token"
		c.Done()
		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusForbidden, response)
		return
	}

	tk := &internals.Token{}
	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		response := "Invalid/Malformed auth token"
		c.AbortWithStatusJSON(http.StatusForbidden, response)
		c.Header("Content-Type", "application/json")
		return
	}

	tokenPart := splitted[1]
	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(internals.Token_key), nil
	})

	if err != nil { //Malformed token, returns with http code 403 as usual
		response := "Invalid/Malformed auth token"
		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusForbidden, response)
		return
	}

	if !token.Valid { //Token is invalid, maybe not signed on this server
		response := "Token is not valid."

		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusForbidden, response)
		return

	}
	c.Set("id", tk.UserId)

	//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
	fmt.Sprintf("User %", tk.UserId) //Useful for monitoring

	c.Next() //proceed in the middleware chain!
}

// func to check if a user is authorized to perform the required action for the api request
func Authorize(c *gin.Context) {
	id, check := c.Get("id")

	if !check {
		c.AbortWithStatusJSON(http.StatusBadRequest, "failed to get user id")
		return
	}
	ctx := context.TODO()
	userID := int32(id.(uint))

	action := getAction(c.Request.Method)
	UserID := &rb.AssignUser{UserID: userID}
	roles, err := client.RbacClient.GetAllUserRoles(ctx, UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to get roles for user id")
		return
	}
	for _, role := range roles.Roles {
		for _, roleaction := range role.Action {
			if roleaction == action {
				log.Println("user authorized ")
				c.Next()
				return
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, "user doest have the necessary permissions")
	return
}

func getAction(action string) string {
	switch action {
	case "GET":
		return "read"
	case "DELETE":
		return "delete"
	default:
		return "write"
	}
}
