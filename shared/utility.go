package shared

import (
	"eskimoe-client/api"
)

func GetRoomNameFromId(categories []api.Category, roomId int) string {
	for _, category := range categories {
		for _, room := range category.Rooms {
			if room.ID == roomId {
				return room.Name
			}
		}
	}
	return ""
}

func GetTotalRoomCount(categories []api.Category) int {
	count := 0
	for _, category := range categories {
		count += len(category.Rooms)
	}
	return count
}

// Returns the next room ID in the list of rooms if there is a next room, otherwise returns the first room ID
func GetNextRoomId(categories []api.Category, currentRoomId int) int {
	for i, category := range categories {
		for j, room := range category.Rooms {
			if room.ID == currentRoomId {
				if j+1 < len(category.Rooms) {
					return category.Rooms[j+1].ID
				} else if i+1 < len(categories) {
					return categories[i+1].Rooms[0].ID
				} else {
					return categories[0].Rooms[0].ID
				}
			}
		}
	}
	return -1
}

// Returns the previous room ID in the list of rooms if there is a previous room, otherwise returns the last room ID
func GetPreviousRoomId(categories []api.Category, currentRoomId int) int {
	for i, category := range categories {
		for j, room := range category.Rooms {
			if room.ID == currentRoomId {
				if j-1 >= 0 {
					return category.Rooms[j-1].ID
				} else if i-1 >= 0 {
					return categories[i-1].Rooms[len(categories[i-1].Rooms)-1].ID
				} else {
					return categories[len(categories)-1].Rooms[len(categories[len(categories)-1].Rooms)-1].ID
				}
			}
		}
	}
	return -1
}
