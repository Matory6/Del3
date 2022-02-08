package Rivercrossing

import "fmt"

var state "Nothing is in the boat"

func ViewState() string {
    return state
}

func PutInFox() {
    state = "---\\ \\_Fox_/ ____________________/---"
}

func PutInCorn() {
    state = "---\\ \\_______\\_Corn_/ __________/---"
}

func PutInChicken() {
    state = "---\\ \\______________\\_Chicken_/ /---"
}