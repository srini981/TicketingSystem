package internals

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// key for JWT token
var Token_key = "test_key"

// user struct for handling user DATA
type User struct {
	Name     string
	Phone    string
	Email    string
	EmpID    string
	roles    []int
	Password string `json:"password"`
	Token    string `json:"token" sql:"-"`
}

//ticket struct for handling ticket DATA

type Ticket struct {
	ID          int
	TicketName  string
	Description string
	AssignedTo  int32
	Priority    int32
	Status      string
	Timestamp   time.Time
}

//role struct for handling role DATA

type Role struct {
	ID          int
	Name        string
	Description string
	Action      []string
	Timestamp   time.Time
}

//api response struct for handling reponse from handlers

type Response struct {
	Err      error
	Response interface{}
	Message  string
}

// token struct for JWT
type Token struct {
	UserId uint
	jwt.StandardClaims
}

type UserLogin struct {
	Name     string
	Password string
}

// Token response for JWT
type TokenResponse struct {
	UserId uint
	Name   string
	Token  string
}
