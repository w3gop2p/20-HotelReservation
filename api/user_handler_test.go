package api

import (
	"20-HotelReservation/types"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"testing"
)

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	app := fiber.New()
	userHandler := NewUserHandler(tdb.User)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "James",
		LastName:  "Bond",
		Password:  "Parola",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("expected user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected the EncryptedPassword no to be in json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected username %s but got %s", params.FirstName, user.FirstName)
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected username %s but got %s", params.LastName, user.LastName)
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected username %s but got %s", params.Email, user.Email)
	}

}
