package internals

import "time"

// ticket struct to handle ticket details
type Ticket struct {
	ID          int
	TicketName  string
	Description string
	AssignedTo  int
	Priority    int
	Status      string
	Timestamp   time.Time
}
