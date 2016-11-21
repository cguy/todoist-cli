package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/daviddengcn/go-colortext"
	"github.com/twinj/uuid"
)

func addLabel(name string) (*Label, error) {
	// Get user token
	savedToken, err := getSavedToken()
	check(err)
	token := savedToken.Token

	// Generate some UUID to avoid duplicates
	uuid.SwitchFormat(uuid.FormatCanonical)
	generatedUUID := uuid.NewV4().String()
	tmpID := uuid.NewV4().String()

	// Create Label instance
	label := Label{}
	label.Name = url.QueryEscape(name)

	bytes, _ := json.Marshal(label)

	// Launch the HTTP request
	url := fmt.Sprintf("https://todoist.com/API/v6/sync?token=%s&commands=[{\"type\":\"label_add\",\"temp_id\":\"%s\",\"uuid\":\"%s\",\"args\":%s}]", token, tmpID, generatedUUID, string(bytes))
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	// Error case
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("An error occured during the label creation \"" + name + "\": " + resp.Status + "; response: " + string(body) + ")")
	}

	// Success
	var status SyncStatus
	err = json.Unmarshal(body, &status)
	check(err)
	label.Id = status.TempIdMapping[tmpID]

	return &label, nil
}

// ListLabels return a list of the user's todoist labels
func ListLabels() {
	configuration, err := getSavedToken()
	check(err)
	url := "https://todoist.com/API/v6/sync?token=" + configuration.Token + "&seq_no=0&seq_no_global=0&resource_types=[\"labels\"]"
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
	labels := data.Labels
	if labels != nil {
		for _, element := range labels {
			ct.ChangeColor(ct.Cyan, false, ct.None, false)
			fmt.Print(element.Id)
			ct.ChangeColor(ct.Blue, false, ct.None, false)
			fmt.Println(" " + element.Name)
		}
		ct.ResetColor()

		err = ioutil.WriteFile(getLabelsFilePath(), body, 0644)
		check(err)
	}
}

func getLabelFromName(name string) (*Label, error) {
	bytes, err := ioutil.ReadFile(getLabelsFilePath())
	check(err)
	var data *Data
	err = json.Unmarshal(bytes, &data)
	check(err)

	labels := data.Labels
	for _, element := range labels {
		if element.Name == name {
			return &element, nil
		}
	}
	return nil, errors.New("Unable to retrieve label from given name: " + name)
}

func getLabel(id int) (*Label, error) {
	bytes, err := ioutil.ReadFile(getLabelsFilePath())
	check(err)
	var data *Data
	err = json.Unmarshal(bytes, &data)
	check(err)

	labels := data.Labels
	for _, element := range labels {
		if element.Id == id {
			return &element, nil
		}
	}
	return nil, fmt.Errorf("Unable to retrieve label from the id: %d", id)
}
