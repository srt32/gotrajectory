package gotrajectory

import (
  "os"
  "testing"
)

func TestUrlGeneration(t *testing.T) {
  accountName := "Ralph"
  projectName := "P"
  url := StoriesUrl(accountName, projectName)
  expectedUrl := "https://www.apptrajectory.com/api/" + os.Getenv("TRAJECTORY_API_KEY") + "/accounts/" +
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

func TestGetStories(t *testing.T) {
  accountName := "simon"
  projectName := "test"

  stories, err := GetStories(accountName, projectName)

  storyTitle := "Whoohoo!  My first trajectory"
  if err != nil || stories[0].Title != storyTitle {
    t.Error(stories[0].Title)
  } else {
    t.Log("Passed")
  }
}
