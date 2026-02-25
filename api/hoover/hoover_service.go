package hoover

import (
	"assignment/api/hoover/contracts/requests"
	"assignment/api/hoover/contracts/responses"
	"assignment/api/hoover/utils"
	"assignment/api/hoover/utils/types"
	"assignment/utils"
	"fmt"
)

type service struct {}

var Service service

func (service) CleaningSessions(request requests.CleaningRequest) (*responses.CleaningResponse, *http_utils.HttpError) {
	room := hoover_types.MakeRoom(request.RoomSize, request.Patches)

	if !(room.Dimensions.Width > 0 && room.Dimensions.Height > 0) {
		return nil, http_utils.MakeBadRequestError("VALUE_OUT_OF_RANGE", "The room size must be a positive integer")
	}

	for _, patch := range room.Patches {
		if !hoover_types.IsPositionBoundedWithinDimensions(room.Dimensions, patch) {
			return nil, http_utils.MakeBadRequestError(
				"VALUE_OUT_OF_BOUNDS", fmt.Sprintf("The given patch coords (%d, %d) would place it out of bounds", patch.X, patch.Y),
			)
		}
	}

	hoover := hoover_types.MakeHoover(request.Position)

	if !hoover_types.IsPositionBoundedWithinDimensions(room.Dimensions, hoover.Position) {
		return nil, http_utils.MakeBadRequestError(
			"VALUE_OUT_OF_BOUNDS", fmt.Sprintf("The given hoover coords (%d, %d) would place it out of bounds", hoover.Position.X, hoover.Position.Y),
		)
	}

	patches, err := hoover_utils.Clean(request.Instructions, &room, &hoover)
	if err != nil {
		if err.Error() == "HOOVER_WENT_OUT_OF_BOUNDS" {
			return nil, http_utils.MakeBadRequestError("HOOVER_WENT_OUT_OF_BOUNDS", "The given hoover instructions made it go out of bounds")
		}
		if err.Error() == "HOOVER_RECEIVED_UNKNOWN_INSTRUCTION" {
			return nil, http_utils.MakeBadRequestError("HOOVER_RECEIVED_UNKNOWN_INSTRUCTION", "Hoover tried to proccess an unknown instruction")
		}
		panic("UNREACHABLE")
	}

	return &responses.CleaningResponse {
		Position: [2]int{hoover.Position.X, hoover.Position.Y},
		Patches: len(patches),
	}, nil
}
