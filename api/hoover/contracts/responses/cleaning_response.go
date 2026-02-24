package responses

type CleaningResponse struct {
	Position [2]int `json:"coords"`
	Patches int `json:"patches"`
}
