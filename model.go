package legionhq

// Data wraps the entire structure of src/data.json
type Data struct {
	KeywordDict        map[string]string `json:"keywordDict"`
	AdditionalKeywords map[string]string `json:"additionalKeywords"`
	CommunityLinks     map[string][]Link `json:"communityLinks"`
	AllCards           map[string]Card   `json:"allCards"`
}

// Link refers to a URL and contains a label
type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Card represents all card types
type Card struct {
	AdditionalTags []string `json:"additionalTags,omitempty"`
	ID             string   `json:"id"`
	CardName       string   `json:"cardName,omitempty"`
	CardType       string   `json:"cardType,omitempty"`
	CardSubType    string   `json:"cardSubType,omitempty"`
	Commander      string   `json:"commander,omitempty"`
	Cost           int      `json:"cost,omitempty"`
	DisplayName    string   `json:"displayName,omitempty"`
	Faction        string   `json:"faction,omitempty"`
	IconLocation   string   `json:"iconLocation,omitempty"`
	ImageLocation  string   `json:"imageLocation,omitempty"`
	IsUnique       bool     `json:"isUnique,omitempty"`
	Keywords       []string `json:"keywords,omitempty"`
	Products       []string `json:"products,omitempty"`
	Rank           string   `json:"rank,omitempty"`
	Requirements   []string `json:"requirements,omitempty"`
	UpgradeBar     []string `json:"upgradeBar,omitempty"`
}
