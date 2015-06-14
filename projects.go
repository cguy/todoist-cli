package main

import (
  "fmt"
  "os"
  "sort"
  "errors"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/daviddengcn/go-colortext"
)

func listProjects() {
  configuration, err := getSavedToken()
  check(err)
  url := "https://todoist.com/API/v6/sync?token=" + configuration.Token + "&seq_no=0&seq_no_global=0&resource_types=[\"projects\"]"
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
  }

  // Cast to error ?
  var dataError TodoistError
  err = json.Unmarshal(body, &dataError)
  if len(dataError.Error) > 0 {
    ct.ChangeColor(ct.Red, true, ct.None, false)
    fmt.Println("Error getting account information, cause: " + dataError.Error)
    ct.ResetColor()
    os.Exit(1)
  }

  // Cast to Data
  var data Data
  err = json.Unmarshal(body, &data)
  check(err)
  projects := data.Projects
  if projects != nil {
    sort.Sort(Projects(projects))
    for index,element := range projects {
      ct.ChangeColor(ct.Cyan, false, ct.None, false)
      fmt.Print(element.Id)
      ct.ChangeColor(ct.Yellow, false, ct.None, false)
      fmt.Print(" [")
      fmt.Print(index)
      fmt.Print("]\t")
      ct.ChangeColor(ct.Blue, false, ct.None, false)
      fmt.Println(element.Name)
    }
    ct.ResetColor()

    err = ioutil.WriteFile(getProjectsFilePath(), body, 0644)
    check(err)
  }
}

func getProjectFromOrder(order int) (*Project, error) {
  bytes, err := ioutil.ReadFile(getProjectsFilePath())
  check(err)
  var data *Data
  err = json.Unmarshal(bytes, &data)
  check(err)

  projects := data.Projects
  for _,element := range projects {
    if element.Order == order {
      return &element, nil
    }
  }
  return nil, errors.New("Unable to retrieve project from project order")
}

func getProject(id int) (*Project, error) {
  bytes, err := ioutil.ReadFile(getProjectsFilePath())
  check(err)
  var data *Data
  err = json.Unmarshal(bytes, &data)
  check(err)

  projects := data.Projects
  for _,element := range projects {
    if element.Id == id {
      return &element, nil
    }
  }
  return nil, errors.New(fmt.Sprintf("Unable to retrieve project from the id: %d", id))
}
