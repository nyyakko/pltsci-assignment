package types

import "fmt"

type Position struct { X int; Y int }

type Hoover struct
{
	Position Position
}

func MakeHoover(position [2]int) Hoover {
	return Hoover {
		Position: Position {
			X: position[0],
			Y: position[1],
		},
	}
}

func (self *Hoover) Advance(instruction byte) error {
	switch (instruction) {
	case 'N': { self.Position.Y += 1; break }
	case 'S': { self.Position.Y -= 1; break }
	case 'E': { self.Position.X += 1; break }
	case 'W': { self.Position.X -= 1; break }
	default: {
		return fmt.Errorf("UNKNOWN_INSTRUCTION_RECEIVED")
	}
	}

	return nil
}
