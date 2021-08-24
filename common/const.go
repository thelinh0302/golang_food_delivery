package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypCategory    = 3
	DbTypeUser       = 4
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
