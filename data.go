package legionhq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RepoURL is the URL to the github repo
const RepoURL = "https://raw.githubusercontent.com/NicholasCBrown/legion-HQ-2.0"

// GetData retrieves a copy of src/data.json from the master branch
func GetData() (Data, error) {
	dataURL := RepoURL + "/master/src/data.json"
	var data Data
	client := http.Client{}
	resp, err := client.Get(dataURL)
	if err != nil {
		return data, err
	}

	if resp.StatusCode != 200 {
		return data, fmt.Errorf("non-Ok response code %s (%d)", resp.Status, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &data)

	return data, nil
}
