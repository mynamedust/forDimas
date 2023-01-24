package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	Player PlayerType
	Rooms  RoomsType
)

func main() {
	initGame()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println(handleCommand(scanner.Text()))
	}
}

func initGame() {
	Player = NewPlayer
	Rooms = RoomsType{kitchen: Kitchen, bedroom: Bedroom, street: Street, korridor: Korridor}
	Player.rooms = &Rooms
	Player.currentRoom = &Rooms.kitchen
}

func handleCommand(command string) string {
	var cmdSplit []string

	cmdSplit = strings.Split(command, " ")
	if !sizeValidation(cmdSplit) {
		return "неизвестная команда"
	}
	switch cmdSplit[0] {
	case "идти":
		return Player.toGo(cmdSplit[1])
	case "взять":
		return Player.toTake(cmdSplit[1])
	case "применить":
		return Player.toUse(cmdSplit[1], cmdSplit[2])
	case "надеть":
		return Player.toWear(cmdSplit[1])
	case "осмотреться":
		return Player.toWatch()
	default:
		return "неизвестная команда"
	}
}

func sizeValidation(strSlice []string) bool {
	switch strSlice[0] {
	case "идти", "взять", "надеть":
		if len(strSlice) == 2 {
			return true
		}
	case "осмотреться":
		if len(strSlice) == 1 {
			return true
		}
	case "применить":
		if len(strSlice) == 3 {
			return true
		}
	}
	return false
}

func (player *PlayerType) toGo(dir string) string {
	var message string
	room := player.currentRoom
	for _, elem := range player.currentRoom.toGo {
		if elem == dir {
			switch elem {
			case "кухня":
				room = &Rooms.kitchen
			case "коридор":
				room = &Rooms.korridor
			case "комната":
				room = &Rooms.bedroom
			case "улица":
				room = &Rooms.street
			case "домой":
				room = &Rooms.korridor
			}
			if room.isLocked {
				return "дверь закрыта"
			}
			player.currentRoom = room
			message = player.currentRoom.entranceText + "можно пройти -"
			break
		}
	}
	if message == "" {
		return ("нет пути в " + dir)
	}
	for idx, elem := range player.currentRoom.toGo {
		message += " " + elem
		if idx == len(player.currentRoom.toGo)-1 {
			break
		}
		message += ","
	}
	return message
}

func (player *PlayerType) toWatch() string {
	room := player.currentRoom
	if room.name == "комната" && len(room.onChair) == 0 && len(room.onTable) == 0 {
		player.currentRoom.toWatch = "пустая комната"
	}
	message := player.currentRoom.toWatch
	if len(player.currentRoom.onTable) != 0 {
		message += "на столе: "
		for idx, elem := range player.currentRoom.onTable {
			message += elem
			if idx == len(player.currentRoom.onTable)-1 && len(player.currentRoom.onChair) == 0 && player.currentRoom.name != "кухня" {
				break
			}
			message += ", "
		}
	}
	if len(player.currentRoom.onChair) != 0 {
		message += "на стуле: "
		for idx, elem := range player.currentRoom.onChair {
			message += elem
			if idx == len(player.currentRoom.onChair)-1 && player.currentRoom.name != "кухня" {
				break
			}
			message += ", "
		}
	}
	if player.currentRoom.name == "кухня"  {
		message += "надо "
		if !inventoryCheck(player.inventory) {
			message += "собрать рюкзак и "
		}
		message += "идти в универ"
	}
	message += ". можно пройти -"
	for idx, elem := range player.currentRoom.toGo {
		message += " " + elem
		if idx == len(player.currentRoom.toGo)-1 {
			break
		}
		message += ","
	}
	return message
}

func inventoryCheck(inventory []string) bool {
	count := 0
	for _, elem := range inventory {
		if elem == "ключи" || elem == "конспекты" {
			count++
		}
	}
	return count == 2
}

func (player *PlayerType) toWear(item string) string {
	if item != "рюкзак" {
		return "неизвестная команда"
	}
	isFind := false
	for idx, elem := range player.currentRoom.onTable {
		if elem == item {
			arr := player.currentRoom.onTable
			arr[idx] = arr[len(arr)-1]
			arr[len(arr)-1] = ""
			player.currentRoom.onTable = arr[:len(arr)-1]
			isFind = true
			break
		}
	}
	for idx, elem := range player.currentRoom.onChair {
		if elem == item {
			arr := player.currentRoom.onChair
			arr[idx] = arr[len(arr)-1]
			arr[len(arr)-1] = ""
			player.currentRoom.onChair = arr[:len(arr)-1]
			isFind = true
			break
		}
	}
	if isFind {
		player.wearing = append(player.wearing, item)
		player.isInventory = true
		return ("вы надели: " + item)
	} else {
		return "нет такого"
	}
}

func (player *PlayerType) toTake(item string) string {
	if item == "рюкзак" {
		return "неизвестная команда"
	}
	if !player.isInventory {
		return "некуда класть"
	}
	isFind := false
	for idx, elem := range player.currentRoom.onTable {
		if elem == item {
			arr := player.currentRoom.onTable
			arr[idx] = arr[len(arr)-1]
			arr[len(arr)-1] = ""
			player.currentRoom.onTable = arr[:len(arr)-1]
			isFind = true
			break
		}
	}
	for idx, elem := range player.currentRoom.onChair {
		if elem == item {
			arr := player.currentRoom.onChair
			arr[idx] = arr[len(arr)-1]
			arr[len(arr)-1] = ""
			player.currentRoom.onChair = arr[:len(arr)-1]
			isFind = true
			break
		}
	}
	if isFind {
		player.inventory = append(player.inventory, item)
		player.isInventory = true
		return ("предмет добавлен в инвентарь: " + item)
	} else {
		return ("нет такого")
	}
}

func (player *PlayerType) toUse(item, dir string) string {
	var find bool
	for _, elem := range player.inventory {
		if elem == item {
			find = true
			break
		}
	}
	if !find {
		return ("нет предмета в инвентаре - " + item)
	}
	if dir != "дверь" {
		return ("не к чему применить")
	}
	for _, elem := range player.currentRoom.toGo {
		switch elem {
		case "кухня":
			if player.rooms.kitchen.isLocked {
				player.rooms.kitchen.isLocked = false
				return ("дверь открыта")
			}
		case "коридор":
			if player.rooms.korridor.isLocked {
				player.rooms.korridor.isLocked = false
				return ("дверь открыта")
			}
		case "комната":
			if player.rooms.bedroom.isLocked {
				player.rooms.bedroom.isLocked = false
				return ("дверь открыта")
			}
		case "улица":
			if player.rooms.street.isLocked {
				player.rooms.street.isLocked = false
				return ("дверь открыта")
			}
		}
	}
	return "неизвестная команда"
}
