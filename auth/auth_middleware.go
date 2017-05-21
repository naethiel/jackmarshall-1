package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/elwinar/token"
	"github.com/julienschmidt/httprouter"
	"github.com/pborman/uuid"
)

// A AuthMiddleware ensure each request is given an initialized context.
type AuthMiddleware struct {
	ctx    AuthContext
	secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{
		ctx:    AuthContext{},
		secret: secret,
	}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	// Generate an ID for the request.
	requestID := uuid.New()

	// Generate a custom context for the request.
	ctx := m.ctx
	ctx.RequestID = requestID

	// Check that the token is valid
	claims, err := token.ParseHS256(r.Header.Get("Authorization"), []byte(m.secret))
	if err != nil {
		log.Println("invalid token:", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		return
	}

	// Get the requester from the claims
	requester, err := NewUserFromClaims(claims["user"])
	if err != nil {
		log.Println("invalid claim:", claims)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		return
	}

	ctx.User = requester

	// Generate a new context for the request.
	clock.Lock()
	cmap[r] = &ctx
	clock.Unlock()

	// Execute the next handlers.
	n(w, r)

	// Remove the context from the map.
	clock.Lock()
	delete(cmap, r)
	clock.Unlock()
}

func NewAuthHandler(next httprouter.Handle, authorizedRole []string, secret string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// Check that the token is valid
		claims, err := token.ParseHS256(r.Header.Get("Authorization"), []byte(secret))
		if err != nil {
			log.Println("invalid token:", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		// Get the requester from the claims
		requester, err := NewUserFromClaims(claims["user"])
		if err != nil {
			log.Println("invalid claim:", claims)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		// Check if the user is roor or has a valid role
		authorization := false
		if authorizedRole == nil || len(authorizedRole) == 0 {
			authorization = true
		}

		for _, role := range requester.Roles {
			if role == "root" {
				authorization = true
				p = append(p, httprouter.Param{"root", "ok"})
				break
			}
			if contains(authorizedRole, role) {
				authorization = true
				break
			}
		}

		if authorization == false {
			log.Println("invalid roles:", requester.Roles)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}
		fmt.Println(requester)
		p = append(p, httprouter.Param{"userId", strconv.Itoa(int(requester.ID))})

		next(w, r, p)
		return
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
