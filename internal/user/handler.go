package user

import (
	"encoding/json"
	"net/http"
	"rest-go/internal/errs"
	"rest-go/internal/httpx"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = make(map[int]User)
var nextID = 1

func CreateUserHandler(w http.ResponseWriter, r *http.Request) error {
	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {

		return errs.BadRequestException("Erro ao decodificar o JSON", "invalid_json")
	}

	u.ID = nextID
	nextID++
	users[u.ID] = u

	return json.NewEncoder(w).Encode(u)

}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) error {
	var list []User

	for _, u := range users {
		list = append(list, u)
	}

	return json.NewEncoder(w).Encode(list)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")

	}

	u, ok := users[id]
	if !ok {
		return errs.NotFoundException("Usuário não encontrado", "user_not_found")
	}

	return json.NewEncoder(w).Encode(u)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}

	u, ok := users[id]
	if !ok {
		return errs.NotFoundException("Usuário não encontrado", "user_not_found")
	}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return errs.BadRequestException("Erro ao decodificar o JSON", "invalid_json")

	}

	users[id] = u

	return json.NewEncoder(w).Encode(u)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}

	_, ok := users[id]
	if !ok {
		return errs.NotFoundException("Usuário não encontrado", "user_not_found")

	}

	delete(users, id)

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func SetupRoutes(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", httpx.Wrap(CreateUserHandler))
		r.Get("/", httpx.Wrap(ListUsersHandler))
		r.Get("/{id}", httpx.Wrap(GetUserHandler))
		r.Put("/{id}", httpx.Wrap(UpdateUserHandler))
		r.Delete("/{id}", httpx.Wrap(DeleteUserHandler))
	})
}
