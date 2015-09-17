package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/daviddengcn/go-colortext"
	"github.com/twinj/uuid"
)

// AddTask adds a task to a user todoist project
func AddTask(taskContent string, projectOrder int, date string, rawLabels string, priority int) {
	// Get user token
	savedToken, err := getSavedToken()
	check(err)
	token := savedToken.Token

	// Retrieve the project
	project, err := getProjectFromOrder(projectOrder)
	check(err)

	// Generate some UUID to avoid duplicates
	uuid.SwitchFormat(uuid.CleanHyphen)
	generatedUUID := uuid.NewV4().String()
	tmpID := uuid.NewV4().String()

	// Get labels id
	strLabels := strings.Split(rawLabels, ",")
	labels := make([]int, len(strLabels))
	for index, labelName := range strLabels {
		label, err := getLabelFromName(labelName)
		if err != nil {
			createdLabel, _ := addLabel(labelName)
			if err != nil {
				labels[index] = createdLabel.Id
			}
			continue
		}
		labels[index] = label.Id
	}

	// Ensure the task is URL-safe
	taskContent = url.QueryEscape(taskContent)
	date = url.QueryEscape(date)

	task := Task{}
	task.Content = taskContent
	task.Priority = priority
	task.DateLang = "en"
	task.Date = date
	task.ProjectId = project.Id
	if len(labels) > 0 {
		task.Labels = labels
	}

	bytes, _ := json.Marshal(task)

	// Launch the HTTP request
	url := fmt.Sprintf("https://todoist.com/API/v6/sync?token=%s&commands=[{\"type\":\"item_add\",\"temp_id\":\"%s\",\"uuid\":\"%s\",\"args\":%s}]", token, tmpID, generatedUUID, string(bytes))
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	// Error case
	if resp.StatusCode != http.StatusOK {
		ct.ChangeColor(ct.Red, false, ct.None, false)
		fmt.Println("An error occured during the task adding (HTTP status: " + resp.Status + "; response: " + string(body) + ")")
		ct.ResetColor()
		return
	}

	// Success
	ct.ChangeColor(ct.Green, false, ct.None, false)
	fmt.Println("Task successfully added to the project: " + project.Name)
	ct.ResetColor()
}

// ListTasks returns a list of the user's todoist tasks
func ListTasks() {
	configuration, err := getSavedToken()
	check(err)
	url := "https://todoist.com/API/v6/sync?token=" + configuration.Token + "&seq_no=0&seq_no_global=0&resource_types=[\"items\"]"
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
	tasks := data.Tasks
	if tasks != nil {
		sort.Sort(Tasks(tasks))
		for _, element := range tasks {
			project, err := getProject(element.ProjectId)
			if err != nil {
				continue
			}

			ct.ChangeColor(ct.Cyan, false, ct.None, false)
			fmt.Print(element.Id)
			fmt.Print(" ")
			ct.ChangeColor(ct.Yellow, false, ct.None, false)
			fmt.Print("[" + project.Name + "]")
			fmt.Print("\t")
			ct.ChangeColor(ct.Blue, false, ct.None, false)
			fmt.Println(element.Content)
		}
		ct.ResetColor()

		err = ioutil.WriteFile(getTasksFilePath(), body, 0644)
		check(err)
	}
}

// MarkTaskAsDone marks a user's todoist taks as done
func MarkTaskAsDone(taskID int) {
	fmt.Println(taskID)
}

// MarkTaskAsUndone marks a user's todoist taks as undone
func MarkTaskAsUndone(taskID int) {
	fmt.Println(taskID)
}
