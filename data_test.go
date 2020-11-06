package legionhq

import (
	"encoding/json"
	"reflect"
	"testing"
)

const darthVaderCommanderJSON = `{
		"cardType": "unit",
		"defense": "red",
		"surges": [],
		"speed": 1,
		"wounds": 8,
		"courage": -1,
		"cardSubtype": "trooper",
		"cardName": "Darth Vader",
		"title": "Dark Lord of the Sith",
		"isUnique": true,
		"rank": "commander",
		"cost": 190,
		"faction": "empire",
		"imageName": "Darth Vader.jpeg",
		"keywords": [
			"Deflect",
			"Immune",
			"Master of the Force",
			"Relentless",
			"Impact",
			"Pierce"
		],
		"upgradeBar": ["force", "force", "force"],
		"history": [
			{
				"date": "12 September 2019",
				"description": "Cost reduced from 200 to 190 points."
			}
		],
		"id": "at"
	}`

func Test_GetData(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "get data", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			dataJSON, _ := json.Marshal(&got)
			t.Log(string(dataJSON))
		})
	}
}

func Test_execCardsJS(t *testing.T) {
	tests := []struct {
		name               string
		wantLenGreaterThan int
		wantErr            bool
	}{
		{name: "data js", wantLenGreaterThan: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execCardsJS()
			if (err != nil) != tt.wantErr {
				t.Errorf("execCardsJS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) <= tt.wantLenGreaterThan {
				t.Errorf("execCardsJS() got = %v, want > %v", len(got), tt.wantLenGreaterThan)
			}
		})
	}
}

func Test_jsonToCard(t *testing.T) {
	type args struct {
		cardJSON string
	}
	tests := []struct {
		name string
		args args
		want Card
	}{
		{
			name: "Darth Vader Commander",
			args: struct{ cardJSON string }{cardJSON: darthVaderCommanderJSON},
			want: Card{
				ID:          "at",
				CardName:    "Darth Vader",
				CardType:    "unit",
				CardSubType: "trooper",
				Cost:        190,
				Faction:     "empire",
				IsUnique:    true,
				Keywords: []string{
					"Deflect",
					"Immune",
					"Master of the Force",
					"Relentless",
					"Impact",
					"Pierce",
				},
				Rank:          "commander",
				UpgradeBar:    []string{"force", "force", "force"},
				ImageLocation: "Darth Vader.jpeg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsonToCard(tt.args.cardJSON); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonToCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
