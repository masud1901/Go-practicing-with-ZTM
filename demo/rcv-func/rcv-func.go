package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) freeSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{
		spaces: make([]Space, 5),
	}
	fmt.Println("Initial parking lot state:")
	for i, space := range lot.spaces {
		fmt.Printf("Space %d occupied: %t\n", i+1, space.occupied)
	}

	fmt.Println("\nAfter occupying spaces 2 and 3, and freeing space 2:")
	occupySpace(&lot, 2)

	lot.occupySpace(3)
	lot.freeSpace(2)

	for i, space := range lot.spaces {
		fmt.Printf("Space %d occupied: %t\n", i+1, space.occupied)
	}

}
