package hoover_utils

import (
	"fmt"
	"slices"

	"assignment/api/hoover/utils/types"
)

func Clean(instructions string, room *hoover_types.Room, hoover *hoover_types.Hoover) ([][2]int, error) {
	var patches [][2]int

	cleanFn := func () {
		if slices.IndexFunc(room.Patches, func(value hoover_types.Position) bool {
			shouldPatch := hoover.Position.X == value.X && hoover.Position.Y == value.Y
			hasCleaned  := slices.IndexFunc(patches, func (patch [2]int) bool {
				return patch[0] == value.X && patch[1] == value.Y
			}) >= 0
			return shouldPatch && !hasCleaned
		}) != -1 {
			patches = append(patches, [2]int{ hoover.Position.X, hoover.Position.Y })
		}
	}

	cleanFn()

	for _, v := range []byte(instructions) {
		cleanFn()

		err := hoover.Advance(v)
		if err != nil {
			return nil, err
		}

		if !hoover_types.IsPositionBoundedWithinDimensions((*room).Dimensions, hoover.Position) {
			return nil, fmt.Errorf("HOOVER_WENT_OUT_OF_BOUNDS")
		}
	}

	return patches, nil
}
