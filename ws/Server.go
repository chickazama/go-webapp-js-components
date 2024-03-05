package ws

import (
	"fmt"
	"log"
	"matthewhope/go-webapp-js-components/services"
	"time"
)

type Server struct {
	Clients map[*Client]bool
}

func NewServer() *Server {
	ret := new(Server)
	ret.Clients = make(map[*Client]bool)
	return ret
}

func (s *Server) Run() {
	for {
		time.Sleep(5 * time.Second)
		n, err := services.GetMaxItemID()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		msg := fmt.Sprintf("%d", n)
		for c := range s.Clients {
			err = c.Conn.WriteMessage(1, []byte(msg))
			if err != nil {
				log.Println(err.Error())
			}
		}
	}

}
