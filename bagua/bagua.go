package bagua

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	//	"sort"
	"time"
)

func BaguaGenerator() (hexa string, signature int) {
	// show coin flips and what thing they represent
	rand.Seed(int64(time.Now().Second()))

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
	hexSignatureMap := map[string]int{
		"Yin changing into yang": 0,
		"Yang changing into yin": 1,
		"Yin unchanging":         0,
		"Yang unchanging":        1,
	}
	// Note binary is counter top to bottom
	// ex:
	// ______
	// ______
	// __  __
	// is 1 + 3 + 0

	signatureToChapter := map[string]int{
		// first row
		"000000": 2,
		"000001": 23,
		"000010": 8,
		"000011": 20,
		"000100": 16,
		"000101": 35,
		"000110": 45,
		"000111": 12,
		// second row
		"001000": 15,
		"001001": 52,
		"001010": 39,
		"001011": 53,
		"001100": 62,
		"001101": 56,
		"001110": 31,
		"001111": 33,
		// third row
		"010000": 7,
		"010001": 4,
		"010010": 29,
		"010011": 59,
		"010100": 40,
		"010101": 64,
		"010110": 47,
		"010111": 6,
		// fourth row
		"011000": 46,
		"011001": 18,
		"011010": 48,
		"011011": 47,
		"011100": 32,
		"011101": 50,
		"011110": 28,
		"011111": 44,
		// fifth row
		"100000": 24,
		"100001": 27,
		"100010": 3,
		"100011": 42,
		"100100": 51,
		"100101": 21,
		"100110": 17,
		"100111": 25,
		// sixth row
		"101000": 36,
		"101001": 22,
		"101010": 63,
		"101011": 37,
		"101100": 55,
		"101101": 30,
		"101110": 49,
		"101111": 13,
		// seventh row
		"110000": 19,
		"110001": 41,
		"110010": 60,
		"110011": 61,
		"110100": 54,
		"110101": 38,
		"110110": 58,
		"110111": 10,
		// eigth row
		"111000": 11,
		"111001": 26,
		"111010": 5,
		"111011": 9,
		"111100": 34,
		"111101": 14,
		"111110": 43,
		"111111": 1,
	}

	hexagram := []string{}
	hexSignature := []string{}

	for i := 0; i < 6; i++ {
		// This only works on Linux and MacOS
		cmd := exec.Command("clear")
		// cmd := exec.Command("cls") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
		firstline := possibleResults[rand.Intn(len(possibleResults))]
		hexagram = append(hexagram, firstline)
		if i != 5 {
			fmt.Println("Results in: ", resultsMap[firstline])
		}
		//time.Sleep(time.Second)
		_ = hexSignatureMap[firstline]
		hexSignature = append(hexSignature, strconv.Itoa((hexSignatureMap[firstline])))
	}

	res := []string{}
	for i := len(hexagram) - 1; i >= 0; i-- {
		//write to response here
		state := resultsMap[hexagram[i]]

		res = append(res, state)
		res = append(res, "\n")
	}
	hexa = strings.Join(res, "")
	signature = signatureToChapter[strings.Join(hexSignature, "")]
	return
}
