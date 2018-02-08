package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	sawDamageRecord bool
	command         string
)

func main() {
	write("You enter the forest. In the front of you stay wild wolf. He's very hungry.")
	write("Wolf: \"Hrrrrr!\"")
	write("He try to bite you. What would you do?")
	write("[You now may use commands \"punch\" and \"step back\"]")

	readCommand()
	if command == "punch" {
		write("You punched the wolf in the face. From it appeared, and immediately disappeared number -5")
		sawDamageRecord = true
		write("You: \"What is it?!\"")
		write("Wolf: \"Hrrrrr!\"")
		readCommand()
		if command == "punch" {
			write("You punched wolf again. You saw \"-5\" again.")
			write("You don't know, what it mean. But you know, what wolf madly wants to eat you")
			write("But you really angry too")
			write("[Now you can use \"heavy punch\"]")
			readCommand()
			if command == "heavy punch" {
				write("You: Aaaaaaargh! [smash sound]")
				write("From the head of a wolf splashed some blood")
				write("Now you seen number \"-15\"")
				write("You: What's the hell...")
				write("Wolf: \"Hrrrrr! WOW! WOW!\"")
				write("From the forest:\"WOW! WOW! Ahoooooooooooo!\"")
			}
		} else if command == "step back" {

		}
	} else if command == "step back" {
	}

	write("To be continued...")
	readCommandWithQuestion("[Press Enter]")
}

func write(message string) {
	fmt.Println(message)
	time.Sleep(1 * time.Second)
}

func readCommand() {
	fmt.Print("You: ")
	command = scanLn()
}

func readCommandWithQuestion(text string) {
	fmt.Print(text + " ")
	command = scanLn()
}

func scanLn() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
