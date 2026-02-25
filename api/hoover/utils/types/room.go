package hoover_types

type Dimensions struct { Width int; Height int }

func IsPositionBoundedWithinDimensions(dimensions Dimensions, position Position) bool {
	return (position.X >= 0 && position.X < dimensions.Width) && (position.Y >= 0 && position.Y < dimensions.Height)
}

type Room struct
{
	Dimensions Dimensions;
	Patches []Position
}

func MakeRoom(size [2]int, patches [][2]int) Room {
	return Room {
		Dimensions: Dimensions {
			Width: size[0],
			Height: size[1],
		},
		Patches: func () []Position {
			var _patches []Position
			for _, v := range patches {
				_patches = append(_patches, MakePosition(v))
			}
			return _patches
		}(),
	}
}
