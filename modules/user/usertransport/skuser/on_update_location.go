package skuser

import (
	"Tranning_food/common"
	"Tranning_food/component"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx component.AppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		log.Println("User update location", requester.GetUserId(), location)
	}
}
