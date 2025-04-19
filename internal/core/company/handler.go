package company

import (
	"encoding/json"
	"net/http"
	"rest-go/internal/shared/errs"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	var c Company
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return errs.BadRequestException("JSON inválido", "invalid_json")
	}
	c, err := h.svc.CreateCompany(c)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(c)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) error {
	companies, err := h.svc.ListCompanies()
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(companies)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}
	c, err := h.svc.GetCompany(id)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(c)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}
	var c Company
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return errs.BadRequestException("JSON inválido", "invalid_json")
	}
	c.ID = id
	c, err = h.svc.UpdateCompany(c)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(c)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errs.BadRequestException("ID inválido", "invalid_id")
	}
	err = h.svc.DeleteCompany(id)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
