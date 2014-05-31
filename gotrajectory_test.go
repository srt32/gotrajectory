package gotrajectory

import (
  "testing"
)

func TestUrlGeneration(t *testing.T)  {
  accountName := "Ralph"
  projectName := "P"
  url := StoriesUrl(accountName, projectName)
  expectedUrl := "https://www.apptrajectory.com/api/API_KEY/accounts/" +
    accountName +
    "/projects/" +
    projectName +
    "/stories.json"

  if url != expectedUrl {
    t.Error("Failed. URL is: " + url + ". Expected: " + expectedUrl)
  } else {
    t.Log("Passed")
  }
}
