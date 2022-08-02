package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// function to handle errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// create a list to store usernames
var userList []string

// function to get a list of users on system
func getUsers() []string {
	// determine host os
	hostOS := runtime.GOOS

	// get the list of users for each operating system
	if hostOS == "linux" {

		// /etc/passwd is where users are stored in linux
		file, err := os.Open("/etc/passwd")
		check(err)

		// Set up scanner to read each line
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)

		// loop thorough each line of file
		for fileScanner.Scan() {

			line := strings.Split(fileScanner.Text(), ":")
			uid, err := strconv.Atoi(line[2])
			check(err)
			if line[0] == "nobody" {
				continue
			} else if uid == 0 {
				userList = append(userList, line[0])
			} else if uid >= 1000 {
				userList = append(userList, line[0])
			}
		}

		file.Close()
	} else {
		fmt.Printf("Error: OS %s not supported", hostOS)
	}
	return userList
}

// function to generate random pass
func randPass(passLen int) string {
	// define the alphabet to use for passwords
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"

	// seed to get different result every time
	rand.Seed(time.Now().UnixNano())

	// define some variables for loop
	var finPass string

	// loop until pass len is reached
	for i := 1; i < passLen; i++ {
		// pick random letter from alphabet and add to fin pass
		finPass += string([]rune(alphabet)[rand.Intn(len([]rune(alphabet)))])
	}

	// finish function by returning random pass at given len
	return finPass
}

func main() {
	// get a list of users on the system
	getUsers()
	fmt.Println(userList)

	loginMap := make(map[string]string)
	f, err := os.Create("uudogz.out")
	w := bufio.NewWriter(f)
	check(err)
	for _, element := range userList {
		loginMap[element] = randPass(20)
		_, err := w.WriteString(element + ":" + loginMap[element])
		check(err)
		_, err2 := w.WriteString("\n")
		check(err2)
	}

	w.Flush()
	fmt.Printf("The logins are: %v", loginMap)
}
