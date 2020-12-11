package handler

import (
	"errors"
	"time"

	"github.com/daniilsolovey/http-service/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/reconquest/karma-go"
)

type Handler struct {
	config *config.Config
	users  []User
}

type User struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func NewHandler(
	config *config.Config,
	users []User,
) *Handler {
	return &Handler{
		config: config,
		users:  users,
	}
}

func (handler *Handler) CreateRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Timeout(60 * time.Second))
	router.Route("/", func(r chi.Router) {
		r.Get("/users", handler.PrintUsers)
		r.Get("/users/{id}", handler.GetUserByID)
		r.Put("/users/{id}", handler.UpdateUserByID)
	})

	return router
}

func CreateUsers() []User {
	var users []User
	users = append(
		users,
		User{
			ID:    1,
			Token: "12345",
			Name:  "Andrey",
			Age:   25,
		},

		User{
			ID:    2,
			Token: "5678910",
			Name:  "Dmitry",
			Age:   30,
		},

		User{
			ID:    3,
			Token: "1112131415",
			Name:  "Daniel",
			Age:   26,
		},
	)

	return users
}

func (handler *Handler) findUser(id int) (*User, error) {
	var user User
	for _, item := range handler.users {
		if id == item.ID {
			user = item
			break
		}
	}

	if user.ID == 0 {
		return nil, karma.Format(errors.New("user doesn't exist"), "")
	}

	return &user, nil
}

func (handler *Handler) updateUser(id int, newFields User) {
	for i, item := range handler.users {
		if item.ID == id {
			if newFields.Age != 0 {
				item.Age = newFields.Age
			}

			if newFields.Name != "" {
				item.Name = newFields.Name
			}

			handler.users[i] = item
			break
		}
	}
}
