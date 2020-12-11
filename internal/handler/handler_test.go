package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/daniilsolovey/http-service/internal/config"
	"github.com/reconquest/karma-go"
)

const (
	CONFIG_PATH = "../../testdata/test_config.toml"
	USERS       = "/users/"
)

type Server struct {
	testHandler *Handler
}

func (server Server) createServerWithUsers() *httptest.Server {
	testServer := httptest.NewServer(server.testHandler.CreateRouter())
	return testServer
}

func TestHandler_GetUserByID_ReturnsRequestedUserByID(
	t *testing.T,
) {
	config, err := config.Load(CONFIG_PATH)
	assert.NoError(t, err)

	users := CreateUsers()
	handler := NewHandler(config, users)
	client := Server{
		testHandler: handler,
	}

	testServer := client.createServerWithUsers()
	defer testServer.Close()
	url, err := url.Parse(testServer.URL)
	assert.NoError(t, err)

	// validate ids and name of requested users
	user, err := createGetRequest(url.String(), "1")
	assert.NoError(t, err)

	assert.Equal(
		t, users[0].ID, user.ID, "should contain user with id: %d", users[0].ID,
	)
	assert.Equal(
		t, users[0].Name, user.Name, "should contain user with name: %s", users[0].Name,
	)

	user, err = createGetRequest(url.String(), "2")
	assert.NoError(t, err)

	assert.Equal(
		t, users[1].ID, user.ID, "should contain user with id: %d", users[1].ID,
	)
	assert.Equal(
		t, users[1].Name, user.Name, "should contain user with name: %s", users[1].Name,
	)

	user, err = createGetRequest(url.String(), "3")
	assert.NoError(t, err)

	assert.Equal(
		t, users[2].ID, user.ID, "should contain user with id: %d", users[2].ID,
	)
	assert.Equal(
		t, users[2].Name, user.Name, "should contain user with name: %s", users[2].Name,
	)

	// validate that function returns error
	user, err = createGetRequest(url.String(), "4")
	assert.Error(t, err)
}

func TestHandler_UpdateUserByID_UpdateOnlyNecessaryFieldsForRequestedUserByID(
	t *testing.T,
) {
	config, err := config.Load(CONFIG_PATH)
	assert.NoError(t, err)

	users := CreateUsers()
	handler := NewHandler(config, users)
	client := Server{
		testHandler: handler,
	}

	testServer := client.createServerWithUsers()
	defer testServer.Close()
	url, err := url.Parse(testServer.URL)
	assert.NoError(t, err)

	userID := "1"
	user, err := createGetRequest(url.String(), userID)
	assert.NoError(t, err)

	assert.Equal(
		t, users[0].ID, user.ID, "should contain user with id: %d", users[0].ID,
	)
	assert.Equal(
		t, users[0].Age, user.Age, "should contain user with age: %d", users[0].Age,
	)
	assert.Equal(
		t, users[0].Token, user.Token, "should contain user with token: %s", users[0].Token,
	)
	assert.Equal(
		t, users[0].Name, user.Name, "should contain user with name: %s", users[0].Name,
	)

	newAge := 777
	data := User{
		Age: newAge,
	}

	err = createPutRequest(url.String(), userID, data)
	assert.NoError(t, err)

	user, err = createGetRequest(url.String(), userID)
	assert.NoError(t, err)

	assert.Equal(
		t, users[0].ID, user.ID, "should contain user with id: %d", users[0].ID,
	)
	assert.Equal(
		t, newAge, user.Age, "should contain user with age: %d", newAge,
	)
	assert.Equal(
		t, users[0].Token, user.Token, "should contain user with token: %s", users[0].Token,
	)
	assert.Equal(
		t, users[0].Name, user.Name, "should contain user with name: %s", users[0].Name,
	)

	user, err = createGetRequest(url.String(), userID)
	assert.NoError(t, err)

	assert.Equal(
		t, users[0].ID, user.ID, "should contain user with id: %d", users[0].ID,
	)
	assert.Equal(
		t, users[0].Age, user.Age, "should contain user with age: %d", users[0].Age,
	)
	assert.Equal(
		t, users[0].Token, user.Token, "should contain user with token: %s", users[0].Token,
	)
	assert.Equal(
		t, users[0].Name, user.Name, "should contain user with name: %s", users[0].Name,
	)

	newName := "newNameTest"
	data = User{
		Name: newName,
	}

	err = createPutRequest(url.String(), userID, data)
	assert.NoError(t, err)

	user, err = createGetRequest(url.String(), userID)
	assert.NoError(t, err)

	assert.Equal(
		t, users[0].ID, user.ID, "should contain user with id: %d", users[0].ID,
	)
	assert.Equal(
		t, users[0].Age, user.Age, "should contain user with age: %d", users[0].Age,
	)
	assert.Equal(
		t, users[0].Token, user.Token, "should contain user with token: %s", users[0].Token,
	)
	assert.Equal(
		t, newName, user.Name, "should contain user with name: %s", newName,
	)

	user, err = createGetRequest(url.String(), userID)
	assert.NoError(t, err)

	assert.Equal(
		t, users[0].ID, user.ID, "should contain user with id: %d", users[0].ID,
	)
	assert.Equal(
		t, users[0].Age, user.Age, "should contain user with age: %d", users[0].Age,
	)
	assert.Equal(
		t, users[0].Token, user.Token, "should contain user with token: %s", users[0].Token,
	)
	assert.Equal(
		t, users[0].Name, user.Name, "should contain user with name: %s", users[0].Name,
	)

	data = User{
		Name: newName,
		Age:  newAge,
	}

	err = createPutRequest(url.String(), userID, data)
	assert.NoError(t, err)

	user, err = createGetRequest(url.String(), userID)
	assert.NoError(t, err)

	assert.Equal(
		t, users[0].ID, user.ID, "should contain user with id: %d", users[0].ID,
	)
	assert.Equal(
		t, newAge, user.Age, "should contain user with age: %d", newAge,
	)
	assert.Equal(
		t, users[0].Token, user.Token, "should contain user with token: %s", users[0].Token,
	)
	assert.Equal(
		t, newName, user.Name, "should contain user with name: %s", newName,
	)

}

func createGetRequest(baseURL, id string) (*User, error) {
	var user User
	baseURL = baseURL + USERS + id
	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to create request to url: %s",
			baseURL,
		)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to send an http request",
		)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to decode response by url: %s",
			baseURL,
		)
	}

	return &user, nil
}

func createPutRequest(url, id string, data User) error {
	url = url + USERS + id
	body, err := json.Marshal(data)
	if err != nil {
		return karma.Format(
			err,
			"unable to prepare request body",
		)
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return karma.Format(
			err,
			"unable to create request to url: %s",
			url,
		)
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return karma.Format(
			err,
			"unable to send an http request",
		)
	}

	return nil
}
