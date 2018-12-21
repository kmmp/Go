package main

import (
	"flag"
	"log"
	"strconv"
)

var botNumber = flag.Int("n", 2, "bot number")

type bot struct {
	name  string
	botID int
}

var bots = make(map[bot]bool)
var broadcast = make(chan string)

func main() {
	flag.Parse()

	for i := 1; i <= *botNumber; i++ {
		bot := bot{
			name:  "bot" + strconv.Itoa(i),
			botID: i,
		}
		bots[bot] = true
		log.Printf("%s dołączył do chatu", bot.name)
	}
}

func letsChat() {

}
