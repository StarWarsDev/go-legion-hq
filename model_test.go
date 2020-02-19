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

func TestData_UpgradeCards(t *testing.T) {
	data, err := GetData()
	if err != nil {
		t.Fatal(err)
	}

	cards := data.UpgradeCards()

	if len(cards) == 0 {
		t.Fatal("expected more than 0 upgrade")
	}

	for _, card := range cards {
		if card.CardType != "upgrade" {
			t.Fatal("expected card type \"upgrade\" but got", card.CardType)
		}
	}
}

func TestData_UnitCards(t *testing.T) {
	data, err := GetData()
	if err != nil {
		t.Fatal(err)
	}

	cards := data.UnitCards()
	if len(cards) == 0 {
		t.Fatal("expected more than 0 unit cards")
	}

	for _, card := range cards {
		if card.CardType != "unit" {
			t.Fatal("expected card type \"unit\" but got", card.CardType)
		}
	}
}

func TestData_CommandCards(t *testing.T) {
	data, err := GetData()
	if err != nil {
		t.Fatal(err)
	}

	cards := data.CommandCards()
	if len(cards) == 0 {
		t.Fatal("expected more than 0 command cards")
	}

	for _, card := range cards {
		if card.CardType != "command" {
			t.Fatal("expected card type \"command\" but got", card.CardType)
		}
	}
}
