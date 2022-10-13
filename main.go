package main

import (
	"github.com/derickr/go-litra-driver"
	"os"
	"strconv"
	"strings"
)
import "log"

func main() {
	ld, err := litra.New()
	if err != nil {
		log.Fatalf("Couldn't open light %v", err)
	}
	defer ld.Close()

	if len(os.Args) > 1 {
		cmd := strings.ToLower(os.Args[1])
		switch cmd {
		case "on":
			ld.TurnOn()
			break
		case "off":
			ld.TurnOff()
			break
		case "level":
		case "bright":
			level, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatalf("Couldn't parse brightness %v", err)
			}
			ld.SetBrightness(level)
			break
		case "temp":
			temp, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatalf("Couldn't parse temperature %v", err)
			}
			ld.SetTemperature(int16(temp))
			break
		}
	}
}
