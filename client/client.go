package client

import (
	//"encoding/json"
	"fmt"
	//"io"
	"encoding/json"
	"github.com/thebitmonk/storedb"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	_email       string = "email"
	_password    string = "password"
	_endpoint    string = "endpoint"
	_command     string = "command"
	_options     string = "options"
	_description string = "description"
	_prefix      string = "prefix"
)

type HttpResponse struct {
	response *http.Response
	err      error
}

type Client struct {
	endPoint   string
	email      string
	password   string
	httpClient *http.Client
}

func NewClient(path string) (*Client, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error occurred while trying to read file. Exiting")
		panic(err)
	}

	config := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(string(line), "=")
		if len(parts) == 2 {
			config[string(parts[0])] = string(parts[1])
		}
	}
	client := Client{
		endPoint:   config[_endpoint],
		email:      config[_email],
		password:   config[_password],
		httpClient: http.DefaultClient,
	}

	return &client, nil
}

func (c *Client) SearchCommand(prefix string) (bool, string, []storedb.Value) {

	form := url.Values{}
	form.Set(_prefix, prefix)
	u, _ := url.ParseRequestURI(c.endPoint)
	u.Path = "/search"
	urlStr := fmt.Sprintf("%v", u)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))
	if err != nil {
		return false, "Failed to search in Yo DB :(", nil
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return false, "Failed to search in Yo DB :(", nil
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, "Got some spooky responses from Yo DB. Something is fishy !", nil
	}
	var subcommands = make([]storedb.Value, 10)
	json.Unmarshal(contents, &subcommands)
	return true, "Yay found results in Yo DB !", subcommands

}

func (c *Client) AddCommand(command, options, description string) (bool, string) {
	if c.email == "" {
		return false, "Email is not set. Kindly add email=varunjain003@gmail.com at location /tmp/hack.config file"
	}

	if c.password == "" {
		return false, "Password is not set. Kindly add password=******* at location /tmp/hack.config file"
	}

	form := url.Values{}
	form.Set(_command, command)
	form.Add(_options, options)
	form.Add(_description, description)
	form.Add(_email, c.email)
	u, _ := url.ParseRequestURI(c.endPoint)
	u.Path = "/add"
	urlStr := fmt.Sprintf("%v", u)

	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := c.httpClient.Do(req)
	if resp.StatusCode == http.StatusOK {
		return true, fmt.Sprintf("Successfully added command '%s %s' to Yo DB", command, options)
	} else {
		return false, fmt.Sprintf("Failed to add command to Yo DB :(")
	}
}
