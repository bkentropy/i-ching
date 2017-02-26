package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	//	"sort"
	"time"
)

func main() {
	// show coin flips and what thing they represent
	rand.Seed(int64(time.Now().Minute()))
	possibleResults := []string{
		"Yin changing into yang",
		"Yin unchanging",
		"Yin unchanging",
		"Yin unchanging",
		"Yin unchanging",
		"Yin unchanging",
		"Yin unchanging",
		"Yin unchanging",
		"Yang changing into yin",
		"Yang changing into yin",
		"Yang changing into yin",
		"Yang unchanging",
		"Yang unchanging",
		"Yang unchanging",
		"Yang unchanging",
		"Yang unchanging",
	}
	resultsMap := map[string]string{
		"Yin changing into yang": "__  __ x",
		"Yang changing into yin": "______ x",
		"Yin unchanging":         "__  __",
		"Yang unchanging":        "______",
	}
	hexagram := []string{}

	for i := 0; i < 6; i++ {
		// This only works on Linux and MacOS
		cmd := exec.Command("clear")
		// cmd := exec.Command("cls") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
		firstline := possibleResults[rand.Intn(len(possibleResults))]
		hexagram = append(hexagram, firstline)
		fmt.Println("Your first line is:")
		fmt.Println(hexagram)
		if i != 5 {
			fmt.Println("Results in: ", resultsMap[firstline])
		}
		time.Sleep(time.Second)
	}
	for i := len(hexagram) - 1; i >= 0; i-- {
		fmt.Println(resultsMap[hexagram[i]])
	}
}
