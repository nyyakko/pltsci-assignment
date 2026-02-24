package requests

type CleaningRequest struct {
	RoomSize [2]int `json:"roomSize"`
	Position [2]int `json:"coords"`
	Patches [][2]int `json:"patches"`
	Instructions string `json:"instructions"`
}
