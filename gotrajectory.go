package gotrajectory

import()

const (
  baseUrl = "https://www.apptrajectory.com/api"
  apiKey = "API_KEY"
)

func StoriesUrl(accountName, projectName string) (url string) {
  projectPath := "/accounts/" + accountName + "/projects/" + projectName
  endPoint := "/stories"
  format := ".json"

  url = baseUrl + "/" + apiKey + projectPath + endPoint + format

  return url
}
