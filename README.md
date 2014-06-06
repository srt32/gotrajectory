gotrajectory
============

A golang wrapper for [apptrajectory](https://www.apptrajectory.com/)

### Installation

```
go get github.com/srt32/gotrajectory
```

### Example

`gotrajectory` expects an ENV var called `TRAJECTORY_API_KEY` to be set with your API key.

```
package main

import(
  trajectory "github.com/srt32/gotrajectory"
  "fmt"
)

func main(){
  accountName := "myAccount"
  projectName := "myProject"

  stories, err := trajectory.GetStories(accountName, projectName)

  if err != nil {
    fmt.Println(err)
  } else {
    for _, story := range stories {
      fmt.Println(story)
    }
  }
}

```


### Development

Feel free to jump in. Always run `go test`
