package legionhq

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestDataModel(t *testing.T) {
	jsonStr, err := ioutil.ReadFile("./sample-data.json")
	if err != nil {
		t.Fatal(err)
	}

	var data Data
	_ = json.Unmarshal([]byte(jsonStr), &data)
}

func TestGetData(t *testing.T) {
	_, err := GetData()
	if err != nil {
		t.Fatal(err)
	}
}
