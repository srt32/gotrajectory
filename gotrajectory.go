package gotrajectory

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

const (
	BASEURL = "https://www.apptrajectory.com/api"
)

type Story struct {
	Archived          bool     `json:"archived"`
	AssigneeId        int      `json:"assignee_id"`
	CommentsCount     int      `json:"comments_count"`
	CreatedAt         string   `json:"created_at"`
	Deleted           bool     `json:"deleted"`
	DesignNeeded      bool     `json:"design_needed"`
	DevelopmentNeeded bool     `json:"development_needed"`
	Id                int      `json:"id"`
	IterationId       int      `json:"iteration_id"`
	Points            int      `json:"points"`
	Position          int      `json:"position"`
	State             string   `json:"state"`
	StateEvents       []string `json:"state_events"`
	TaskType          string   `json:"task_type"`
	Title             string   `json:"title"`
	UpdatedAt         string   `json:"updated_at"`
	UserId            int      `json:"user_id"`
}

type response struct {
	Stories []Story
}

func GetStories(accountName, projectName string) ([]Story, error) {
	url := StoriesUrl(accountName, projectName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r.Stories, nil
}

func StoriesUrl(accountName, projectName string) (url string) {
	apiKey := os.Getenv("TRAJECTORY_API_KEY")
	projectPath := "/accounts/" + accountName + "/projects/" + projectName
	endPoint := "/stories"
	format := ".json"

	url = BASEURL + "/" + apiKey + projectPath + endPoint + format

	return url
}
