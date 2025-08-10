package clients

import "github.com/google/uuid"

type Clients struct {
	ClientId uuid.UUID
	message string
	HubId string
}

func (c *Clients) CreateClient() (*Clients,error) {
	return &Clients{
		ClientId: uuid.New(),
		message: "",
		HubId : "",
	}, nil
}

