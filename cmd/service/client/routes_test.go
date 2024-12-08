package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/roman-haidarov/go-rest-api/types"
)

func TestClientServiceHandlers(t *testing.T) {
	clientStore := &mockClientStore{}
	handler := NewHandler(clientStore)

	t.Run("should fail if the client payload is invalid", func(t *testing.T) {
		payload := types.RegisterClientPayload{
			IdentificationNo: "960517351807",
			Phone:            "+77018732323",
			Email:            "MacDonald@gmail.com",
			Password:  			  "123JdsaW",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(
			http.MethodPost, "/register", bytes.NewBuffer(marshalled),
		)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockClientStore struct {}

func (m *mockClientStore) GetClientByIin(iin string) (*types.Client, error) {
	return nil, fmt.Errorf("client not found")
}

func (m *mockClientStore) GetClientById(id int) (*types.Client, error) {
	return nil, fmt.Errorf("client not found")
}

func (m *mockClientStore) CreateClient(types.Client) error {
	return nil
}
