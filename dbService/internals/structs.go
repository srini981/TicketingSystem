package internals

import (
	"github.com/lib/pq"
)

type User struct {
	ID        int
	Name      string
	Phone     string
	Email     string
	EmpID     string
	Roles     pq.StringArray `gorm:"type:text[]"`
	Timestamp string
	Password  string
	Company   string
}

type Ticket struct {
	ID          int
	TicketName  string
	Description string
	AssignedTo  int
	Priority    int
	Status      string
	Timestamp   string
	Labels      pq.StringArray `gorm:"type:text[]"`
	CreatedBy   int32
}

type Role struct {
	ID          int
	Name        string
	Description string
	Action      pq.StringArray `gorm:"type:text[]"`
	Timestamp   string
	Users       pq.Int32Array `gorm:"type:integer[]"`
	CreatedBy   int32
}
