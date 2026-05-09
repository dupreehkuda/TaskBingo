package game_realtime

import (
	"context"
	"sync"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type roomLoader interface {
	LoadRoom(ctx context.Context, gameID string) (*models.Room, error)
}

type Hub struct {
	mu     sync.Mutex
	rooms  map[string]*models.Room
	loader roomLoader
}

func NewHub(loader roomLoader) *Hub {
	return &Hub{rooms: make(map[string]*models.Room), loader: loader}
}

// GetOrCreate returns the room for gameID, loading it on first connect.
func (h *Hub) GetOrCreate(ctx context.Context, gameID string) (*models.Room, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if r, ok := h.rooms[gameID]; ok {
		return r, nil
	}
	r, err := h.loader.LoadRoom(ctx, gameID)
	if err != nil {
		return nil, err
	}
	h.rooms[gameID] = r
	return r, nil
}

// Leave clears the slot held by userID (Player1 or Player2) on the room.
// If both slots are empty afterwards, the room is removed from the hub.
func (h *Hub) Leave(gameID, userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	r, ok := h.rooms[gameID]
	if !ok {
		return
	}
	if r.Player1 != nil && r.Player1.Id == userID {
		r.Player1 = nil
	}
	if r.Player2 != nil && r.Player2.Id == userID {
		r.Player2 = nil
	}
	if r.Player1 == nil && r.Player2 == nil {
		delete(h.rooms, gameID)
	}
}
