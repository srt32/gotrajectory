package gotrajectory

import(
  "errors"
  "encoding/json"
  "net/http"
  "os"
)

const (
  BASEURL = "https://www.apptrajectory.com/api"
)

type Story struct {
  Title string `json:"title"`
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
