package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

const defaultAmount = 1

func main() {
	var amount int
	var err error
	if len(os.Args) > 1 {
		amount, err = strconv.Atoi(os.Args[1])
		switch {
		case err != nil:
			fmt.Printf("can't get random uuid amount: %s", err.Error())

			os.Exit(1)
		case amount <= 0:
			fmt.Printf("guid amount to generate could not be <= 0")

			os.Exit(1)
		}

		amount--
	}

	for amount += defaultAmount; amount > 0; amount-- {
		uuid, err := uuid.NewRandom()
		if err != nil {
			fmt.Printf("can't generate random uuid: %s", err.Error())

			os.Exit(1)
		}

		fmt.Println(uuid.String())
	}
}
