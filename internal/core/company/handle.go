package company

import (
	"encoding/json"
	"net/http"
	"rest-go/internal/infra/httpx"
	"rest-go/internal/shared/errs"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CNPJ string `json:"cnpj"`
}

var companies = make(map[int]Company)
var nextCompanyID = 1

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) error {
	var c Company

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return errs.BadRequestException("Erro ao decodificar o JSON", "invalid_json")
	}

	c.ID = nextCompanyID
	nextCompanyID++
	companies[c.ID] = c

	return json.NewEncoder(w).Encode(c)
}

func ListCompaniesHandler(w http.ResponseWriter, r *http.Request) error {
	var list []Company
	for _, c := range companies {
		list = append(list, c)
	}
	return json.NewEncoder(w).Encode(list)
}

func GetCompanyHandler(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}
	c, ok := companies[id]
	if !ok {
		return errs.NotFoundException("Empresa não encontrada", "company_not_found")
	}
	return json.NewEncoder(w).Encode(c)
}

func UpdateCompanyHandler(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}

	c, ok := companies[id]
	if !ok {
		return errs.NotFoundException("Empresa não encontrada", "company_not_found")
	}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return errs.BadRequestException("Erro ao decodificar o JSON", "invalid_json")
	}

	companies[id] = c
	return json.NewEncoder(w).Encode(c)
}

func DeleteCompanyHandler(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}

	_, ok := companies[id]
	if !ok {
		return errs.NotFoundException("Empresa não encontrada", "company_not_found")
	}
	delete(companies, id)
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func SetupRoutes(r *chi.Mux) {
	r.Route("/companies", func(r chi.Router) {
		r.Post("/", httpx.Wrap(CreateCompanyHandler))
		r.Get("/", httpx.Wrap(ListCompaniesHandler))
		r.Get("/{id}", httpx.Wrap(GetCompanyHandler))
		r.Put("/{id}", httpx.Wrap(UpdateCompanyHandler))
		r.Delete("/{id}", httpx.Wrap(DeleteCompanyHandler))
	})
}
