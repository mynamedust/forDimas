package main

type PlayerType struct {
	inventory   []string
	isInventory bool
	wearing     []string
	rooms       *RoomsType
	currentRoom *RoomType
}

type RoomsType struct {
	kitchen, bedroom, korridor, street RoomType
}

type RoomType struct {
	name         string
	entranceText string
	toGo         []string
	toWatch      string
	onTable      []string
	onChair      []string
	isLocked     bool
}

var (
	NewPlayer = PlayerType{
		inventory:   []string{},
		isInventory: false,
		wearing:     []string{},
	}
	Kitchen RoomType = RoomType{
		name:         "кухня",
		entranceText: "кухня, ничего интересного. ",
		toGo:         []string{"коридор"},
		toWatch:      "ты находишься на кухне, ",
		onTable:      []string{"чай"},
		onChair:      []string{},
		isLocked:     false,
	}
	Korridor RoomType = RoomType{
		name:         "коридор",
		entranceText: "ничего интересного. ",
		toGo:         []string{"кухня", "комната", "улица"},
		toWatch:      "ничего интересного",
		onTable:      []string{},
		onChair:      []string{},
		isLocked:     false,
	}
	Bedroom RoomType = RoomType{
		name:         "комната",
		entranceText: "ты в своей комнате. ",
		toGo:         []string{"коридор"},
		toWatch:      "",
		onTable:      []string{"ключи", "конспекты"},
		onChair:      []string{"рюкзак"},
		isLocked:     false,
	}
	Street RoomType = RoomType{
		name:         "улица",
		entranceText: "на улице весна. ",
		toGo:         []string{"домой"},
		toWatch:      "на улице весна",
		onTable:      []string{},
		onChair:      []string{},
		isLocked:     true,
	}
)
