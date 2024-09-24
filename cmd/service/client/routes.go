package client

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/roman-haidarov/go-rest-api/cmd/service/auth"
	"github.com/roman-haidarov/go-rest-api/types"
	"github.com/roman-haidarov/go-rest-api/utils"
)

type Handler struct {
	store types.ClientStore
}

func NewHandler(store types.ClientStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var client types.RegisterClientPayload
	if err := utils.ParseJSON(r, client); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetClientByIin(client.IdentificationNo)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("client with IIN %s already exists", client.IdentificationNo))
	}

	hashedPassword, err := auth.HashPassword(client.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateClient(types.Client{
    IdentificationNo: client.IdentificationNo,
    Phone:            client.Phone,
    Email:            client.Email,
    APIAccessToken:   hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteError(w, http.StatusCreated, nil)
}
