package hoover

import (
	"assignment/api/hoover/contracts/requests"
	"assignment/api/hoover/contracts/responses"
	"assignment/api/hoover/utils"
	"assignment/api/hoover/utils/types"
)

type service struct {}

var Service service

func (service) CleaningSessions(request requests.CleaningRequest) (*responses.CleaningResponse, error) {
	room := types.MakeRoom(request.RoomSize, request.Patches)
	hoover := types.MakeHoover(request.Position)

	patches, err := utils.Clean(request.Instructions, &room, &hoover)
	if err != nil {
		return nil, err
	}

	return &responses.CleaningResponse {
		Position: [2]int{hoover.Position.X, hoover.Position.Y},
		Patches: len(patches),
	}, nil
}
