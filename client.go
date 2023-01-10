package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname   string
	port       int
	authToken  string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, port int, token string) *Client {
	return &Client{
		hostname:   hostname,
		port:       port,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) getAllProfiles() (*map[string]server.Profile, error) {
	body, err := c.httpRequest("profiles", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	profiles := map[string]server.Profile{}
	err = json.NewDecoder(body).Decode(&profiles)
	if err != nil {
		return nil, err
	}
	return &profiles, nil
}

// GetItem gets an item with a specific name from the server
func (c *Client) getProfile(name string) (*server.Profile, error) {
	body, err := c.httpRequest(fmt.Sprintf("profile/%v", name), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	Profile := &server.Profile{}
	err = json.NewDecoder(body).Decode(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// NewItem creates a new Profie
func (c *Client) additem(item *server.Profile) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(profile)
	if err != nil {
		return err
	}
	_, err = c.httpRequest("profile", "POST", buf)
	if err != nil {
		return err
	}
	return nil
}

// UpdateItem updates the values of an item
func (c *Client) updateProfile(item *server.Profile) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("profile/%s", profile.Name), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem removes an item from the server
func (c *Client) deleteProfile(profileName string) error {
	_, err := c.httpRequest(fmt.Sprintf("profile/%s", profileName), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.hostname, c.port, path)
}
