package clients

import "github.com/google/uuid"

type Clients struct {
	ClientId uuid.UUID `json:"clientId"`
	ClientMessage  chan string    `json:"clientMessage"`
	HubId    string    `json:"hubId"`
}

func (c *Clients) CreateClient() (*Clients, error) {
	return &Clients{
		ClientId: uuid.New(),
		ClientMessage: make(chan string),
		HubId:    "",
	}, nil
}
