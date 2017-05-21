package auth

import (
	"net/http"
	"sync"
)

// AuthContext enable request-associated data and structures to cross middelwares
// and handlers boundaries.
type AuthContext struct {
	RequestID string
	User      *User
}

var cmap = map[*http.Request]*AuthContext{}
var clock sync.RWMutex

// Context returns the context associated with the given request, or nil if
// none.
func Context(r *http.Request) *AuthContext {
	clock.RLock()
	defer clock.RUnlock()
	return cmap[r]
}
