package utils

import (
	"fmt"
	"slices"

	"assignment/api/hoover/utils/types"
)

func Clean(instructions string, room *types.Room, hoover *types.Hoover) ([][2]int, error) {
	var patches [][2]int

	isBoundedWithinRoom := func(room types.Room, position types.Position) bool {
		return (position.X >= 0 && position.X < room.Dimensions.Width) && (position.Y >= 0 && position.Y < room.Dimensions.Height)
	}

	for _, v := range []byte(instructions) {
		if !isBoundedWithinRoom(*room, hoover.Position) {
			return nil, fmt.Errorf("HOOVER_WENT_OUT_OF_BOUNDS")
		}

		if slices.IndexFunc(room.Patches, func(value [2]int) bool {
			shouldPatch := hoover.Position.X == value[0] && hoover.Position.Y == value[1]
			hasPatched  := slices.Index(patches, value) >= 0
			return shouldPatch && !hasPatched
		}) != -1 {
			patches = append(patches, [2]int{ hoover.Position.X, hoover.Position.Y })
		}

		err := hoover.Advance(v)
		if err != nil {
			return nil, err
		}
	}

	return patches, nil
}
