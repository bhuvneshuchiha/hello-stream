package hub

import (
	"errors"
	"sync"

	"github.com/bhuvneshuchiha/hello-stream/internal/clients"
	"github.com/google/uuid"
)

type Hub struct {
	HubId        uuid.UUID                      `json:"hubId"`
	TotalClients int                            `json:"totalClients"`
	Clients      map[uuid.UUID]*clients.Clients `json:"clients"`
	HubMessages  chan string                    `json:"hubMessages"`
	Mu *sync.Mutex
}

func CreateHub() *Hub {
	return &Hub{
		HubId:        uuid.New(),
		TotalClients: 0,
		Clients:      make(map[uuid.UUID]*clients.Clients),
		HubMessages:  make(chan string),
		Mu : &sync.Mutex{},
	}
}

func (h *Hub) CountTotalClients() int {
	return len(h.Clients)
}

func (h *Hub) AddClient(client *clients.Clients) error {
	newClient, _ := client.CreateClient()
	_, ok := h.Clients[newClient.ClientId]
	if !ok {
		h.Mu.Lock()
		h.Clients[newClient.ClientId] = newClient
		h.Mu.Unlock()
		return nil
	}
	return errors.New("Client already exists")
}

func (h *Hub) RemoveClient(clientId uuid.UUID) error {
	_, ok := h.Clients[clientId]
	if !ok {
		return errors.New("Client does not exist")
	}
	h.Mu.Lock()
	delete(h.Clients, clientId)
	h.Mu.Unlock()
	return nil
}

func (h *Hub) BroadcastMessage() error {
	// take all the client messages and put them in the hub message channel
	// this should also run forever and should push in all the messages of the clients
	return nil
}
