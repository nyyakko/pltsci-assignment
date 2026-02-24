package types

type Dimensions struct { Width int; Height int }

type Room struct
{
	Dimensions Dimensions;
	Patches [][2]int
}

func MakeRoom(size [2]int, patches [][2]int) Room {
	return Room {
		Dimensions: Dimensions {
			Width: size[0],
			Height: size[1],
		},
		Patches: patches,
	}
}
