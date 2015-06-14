package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/daviddengcn/go-colortext"
)

func isTokenValid(token string) (bool) {
  url := "https://todoist.com/API/v6/get_redirect_link?token=" + token
  resp, err := http.Get(url)
  return err == nil && resp.StatusCode == http.StatusOK
}

func checkSavedToken() (bool) {
  data, err := getSavedToken()
  if err != nil {
    return false
  }
  return isTokenValid(data.Token)
}

func getSavedToken() (*UserToken, error) {
  bytes, err := ioutil.ReadFile(getTodoistConfigurationFilePath())
  if err != nil {
    return nil, err
  }
  var userToken *UserToken
  err = json.Unmarshal(bytes, &userToken)
  if err != nil {
    return nil, err
  }
  return userToken, nil
}

func login() {
  fmt.Println("Todoist authentication ...")

  var state = randomString(16)
  resp, err := http.Get("https://todoist.com/oauth/authorize?client_id=&scope=data:read_write,data:delete,project:delete&state="+state)
  check(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  check(err)
  fmt.Printf("%s\n", string(body))

  var data TodoistError
  err = json.Unmarshal(body, &data)
  if len(data.Error) > 0 {
    ct.ChangeColor(ct.Red, false, ct.None, false)
    fmt.Println("Error during authentication, cause: " + data.Error)
    ct.ResetColor()
  }
}

func askUserForToken(parameter string) {
  if (parameter == "") {
    fmt.Println("Please paste here an authentication token that you can find in your Todoist account (Settings -> Todoist Settings -> Account -> API token):")

    in := bufio.NewReader(os.Stdin)
    token, err := in.ReadString('\n')
    check(err)
    token = strings.Trim(token, "\n")
    fmt.Print("\n")
    validateToken(token)
  } else {
    validateToken(parameter)
  }
}

func validateToken(token string) {
  fmt.Print("Checking token... ")

  if isTokenValid(token) {
    ct.ChangeColor(ct.Green, false, ct.None, false)
    fmt.Println("Success")
    ct.ResetColor()

    fmt.Print("Saving token... ")

    data := UserToken{token}
    bytes, err := json.Marshal(data)
    if err == nil {
      createConfigurationDirectory()
      err = ioutil.WriteFile(getTodoistConfigurationFilePath(), bytes, 0644)
      if err == nil {
        ct.ChangeColor(ct.Green, false, ct.None, false)
        fmt.Println("Success")
        fmt.Println("You can now fully use todoist command line (launch `todoist help` to display all available commands)")
        ct.ResetColor()
      } else {
        ct.ChangeColor(ct.Red, false, ct.None, false)
        fmt.Println("Failed", err)
        ct.ResetColor()
      }
    } else {
      ct.ChangeColor(ct.Red, false, ct.None, false)
      fmt.Println("Failed", err)
      ct.ResetColor()
    }
  } else {
    ct.ChangeColor(ct.Red, false, ct.None, false)
    fmt.Println("Failed")
    ct.ResetColor()
  }
}
