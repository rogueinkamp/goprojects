package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"strings"
	"unicode"
)

func isAlphanumeric(s string) bool {
	for _, r := range s {
		// var isLetter bool = unicode.IsLetter(r)
		// var isDigit bool = unicode.IsDigit(r)
		// fmt.Printf("CHECK %c | isletter: %t isDigit: %t\n", r, isLetter, isDigit)

		if !unicode.IsLetter(r) {
			if !unicode.IsDigit(r) {
				return false
			}
		}
	}
	return true
}

func generateMacAddress() string {
	buf := make([]byte, 6)
	var mac net.HardwareAddr

	_, err := rand.Read(buf)
	if err != nil {
	}

	// Set the local bit
	buf[0] |= 2

	mac = append(mac, buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	// fmt.Printf("Here is the generated mac address string %s\n", mac.String())
	return mac.String()
}

// I do not follow convention because convention is confusing, sir!
func convertMacAddress(maybeMacAddress string) {
	// fmt.Print("Enter the maybe mac address> ")

	var inputIsAlphanumeric bool
	// var maybeMacAddress string
	var maybeMacAddressClean string

	// fmt.Scanln(&maybeMacAddress)

	replacer := strings.NewReplacer("[", "", ".", "", ":", "", "]", "")
	maybeMacAddressClean = replacer.Replace(maybeMacAddress)

	// fmt.Println("Here is the user input with out undesirable punct: ", maybeMacAddressClean)

	var maybeMacAddressCleanLower = strings.ToLower(maybeMacAddressClean)
	// fmt.Println("Here is the user input with out undesirable punct and in lowercase: ", maybeMacAddressCleanLower)

	inputIsAlphanumeric = isAlphanumeric(maybeMacAddressCleanLower)
	if !inputIsAlphanumeric == true {
		log.Fatal("User input is not all alphanumeric")
	}
	if len(maybeMacAddressClean) > 12 {
		log.Fatal("User input is longer than 12 characters")
	}

	var add int
	var delimCounter int
	var buffer bytes.Buffer
	var delim string = "-"

	for _, c := range maybeMacAddressClean {
		if delimCounter == 2 {
			buffer.WriteString(delim)
			add++
			buffer.WriteString(string(c))
			add++
			delimCounter = 1
		} else {
			buffer.WriteString(string(c))
			add++
			delimCounter++
		}
	}

	fmt.Printf("CONVERT_OK | original: %s converted: %s\n", maybeMacAddress, buffer.String())

}

func doStuff() {
	var mac_address string = generateMacAddress()
	convertMacAddress(mac_address)
}

func main() {
	for i := 0; i < 10000; i++ {
		go doStuff()
	}

}
