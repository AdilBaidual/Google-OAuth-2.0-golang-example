package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
)

func (h Handler) GoogleCallback(c *fiber.Ctx) error {
	// get auth code from the query

	code := c.Query("code")

	// exchange the auth code that retrieved from google via

	// URL query parameter into an access token.

	token, err := h.conf.Exchange(c.Context(), code)

	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)

	}

	// convert token to user data

	profile, err := ConvertToken(token.AccessToken)

	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)

	}
	fmt.Println(profile)
	//создание юзера или аутентиикация

	return c.JSON(profile)
}

type GooglePayload struct {
	SUB           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func ConvertToken(accessToken string) (*GooglePayload, error) {
	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if httpErr != nil {
		return nil, httpErr
	}

	// clean up when this function returns (destroyed)
	defer resp.Body.Close()

	respBody, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	// Unmarshal raw response body to a map
	var body map[string]interface{}
	if err := json.Unmarshal(respBody, &body); err != nil {
		return nil, err
	}

	// if json body containing error,
	// then the token is indeed invalid. return invalid token err
	if body["error"] != nil {
		return nil, errors.New("Invalid token")
	}

	// Bind JSON into struct
	var data GooglePayload
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
