package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Gist struct {
	Owner       *GithubUser     `json:"owner,omitempty"`
	Id          *string         `json:"id,omitempty"`
	Description *string         `json:"description,omitempty"`
	Files       map[string]File `json:"files,omitempty"`
	CreatedAt   *string         `json:"created_at,omitempty"`
	UpdatedAt   *string         `json:"updated_at,omitempty"`
}

type GithubUser struct {
	Login      *string `json:"login,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	NoneId     *string `json:"node_id,omitempty"`
	AvatarUrl  *string `json:"avatar_url,omitempty"`
	Url        *string `json:"url,omitempty"`
	HtmlUrl    *string `json:"html_url,omitempty"`
	StarredUrl *string `json:"starred_url,omitempty"`
	ReposUrl   *string `json:"repos_url,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type File struct {
	Filename *string `json:"filename,omitempty"`
	Type     *string `json:"type,omitempty"`
	Language *string `json:"language,omitempty"`
	Size     *int    `json:"size,omitempty"`
	Content  *string `json:"content,omitempty"`
}

const API_BASE_URL = "https://api.github.com/"

func Get(id string, token string) (*Gist, error) {
	url := fmt.Sprintf("%sgists/%s", API_BASE_URL, id)
	client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Authorization", "token" + token)
    req.Header.Add("Accept", "application/vnd.github+json")
    if err != nil {
		log.Fatal(err)
	}
    res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	gist := new(Gist)
	json.Unmarshal(body, gist)

	return gist, nil
}
