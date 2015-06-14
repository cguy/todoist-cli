package main

import (
  "fmt"
  "os"
  "crypto/rand"
  "encoding/base64"
  "github.com/mitchellh/go-homedir"
)

func getTodoistConfigurationFilePath() (string) {
  homeDir, err := homedir.Dir()
  check(err)
  return homeDir + "/.config/todoist/todoist.conf"
}

func getProjectsFilePath() (string) {
  homeDir, err := homedir.Dir()
  check(err)
  return homeDir + "/.config/todoist/projects.json"
}

func getTasksFilePath() (string) {
  homeDir, err := homedir.Dir()
  check(err)
  return homeDir + "/.config/todoist/tasks.json"
}

func getLabelsFilePath() (string) {
  homeDir, err := homedir.Dir()
  check(err)
  return homeDir + "/.config/todoist/labels.json"
}

func createConfigurationDirectory() {
  homeDir, err := homedir.Dir()
  check(err)
  homeDir = homeDir + "/.config/todoist/"
  os.MkdirAll(homeDir, 0744)
}

func randomString(size int) (string) {
  rb := make([]byte,size)
  _, err := rand.Read(rb)


  if err != nil {
    fmt.Println(err)
  }

  rs := base64.URLEncoding.EncodeToString(rb)
  return rs
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
