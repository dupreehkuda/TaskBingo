package interfaces

import "net/http"

type Handlers interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type Stored interface {
	Ping(userID string) ([]byte, error)
}
