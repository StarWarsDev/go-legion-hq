package legionhq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"rogchap.com/v8go"
)

// RootDataURL is the base URL for all data files
const RootDataURL = "https://raw.githubusercontent.com/NicholasCBrown/legion-hq-web/master/src/constants"

// CardDataJS is the JavaScript file that defines the card data for LegionHQ
const CardDataJS = "cards.js"

// KeywordDataJS is the JavaScript file that defines the keyword data for LegionHQ
const KeywordDataJS = "keywords.js"

// GetData retrieves all data from files in the master branch
func GetData() (Data, error) {
	data := Data{
		AllCards:       map[string]Card{},
		KeywordDict:    map[string]string{},
		CommunityLinks: map[string][]Link{},
	}

	// cards
	cardsJSON, err := execCardsJS()
	if err != nil {
		return data, err
	}
	var cards map[string]interface{}
	_ = json.Unmarshal([]byte(cardsJSON), &cards)
	for key := range cards {
		cardMap := cards[key]
		cardBytes, _ := json.Marshal(&cardMap)
		card := jsonToCard(string(cardBytes))
		data.AllCards[card.ID] = card
	}

	// TODO: keywords
	keywordsJSON, err := execKeywordsJS()
	if err != nil {
		return data, err
	}
	var keywords map[string]string
	_ = json.Unmarshal([]byte(keywordsJSON), &keywords)
	data.KeywordDict = keywords

	return data, nil
}

func getScript(script string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", RootDataURL, script)
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("non-Ok response code %s (%d)", resp.Status, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func execJS(scriptBytes []byte, varName string) (*v8go.Value, error) {
	// prepare the script
	script := strings.Replace(
		string(scriptBytes),
		fmt.Sprintf("export default %s;", varName),
		"",
		1,
	)

	ctx, _ := v8go.NewContext(nil)
	ctx.RunScript(script, "data.js")
	val, err := ctx.RunScript(fmt.Sprintf("JSON.stringify(%s)", varName), "value.js")
	return val, err
}

func execCardsJS() (string, error) {
	var data string
	scriptBytes, err := getScript(CardDataJS)
	if err != nil {
		return data, err
	}

	val, err := execJS(scriptBytes, "cards")
	if err != nil {
		return data, err
	}

	data = val.String()

	return data, nil
}

func execKeywordsJS() (string, error) {
	var data string
	scriptBytes, err := getScript(KeywordDataJS)
	if err != nil {
		return data, err
	}

	val, err := execJS(scriptBytes, "keywords")
	if err != nil {
		return data, err
	}

	data = val.String()

	return data, nil
}

func jsonToCard(cardJSON string) Card {
	var card Card

	cardBytes := []byte(cardJSON)
	_ = json.Unmarshal(cardBytes, &card)

	return card
}
