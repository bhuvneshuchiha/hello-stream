package hub

import (
	"errors"

	"github.com/bhuvneshuchiha/hello-stream/internal/clients"
	"github.com/google/uuid"
)

type Hub struct {
	HubId        uuid.UUID
	TotalClients int
	Clients      map[uuid.UUID]*clients.Clients
	messages     chan string
}

func CreateHub() *Hub {
	return &Hub{
		HubId: uuid.New(),
		TotalClients: 0,
		Clients : make(map[uuid.UUID]*clients.Clients),
		messages : make(chan string),
	}
}

func (h *Hub) AddClient(client *clients.Clients) error {
	newClient, _ := client.CreateClient()
	h.Clients[newClient.ClientId] = newClient
	return nil
}

func (h *Hub) RemoveClient(clientId uuid.UUID) error {
	_, ok := h.Clients[clientId]
	if !ok {
		return errors.New("Client does not exist")
	}
	delete(h.Clients, clientId)
	return nil
}

func (h *Hub) BroadcastMessage() error {
	return nil
}





