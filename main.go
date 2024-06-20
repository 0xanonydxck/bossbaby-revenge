package main

import (
	"bossbaby/module"
	"fmt"
	"log"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func inputSequenceOfEvents() string {
	answer, err := prompt.New().
		Ask("Enter a sequence of events ex(SRSSRRR):").
		Input("SRSSRRR", input.WithValidateFunc(func(s string) error {
			if len(s) < 1 {
				return fmt.Errorf("please enter a sequence of events")
			}

			return nil
		}))

	if err != nil {
		log.Panic(err)
	}

	return answer
}

func main() {
	sequence := inputSequenceOfEvents()
	result := module.BossBabyRevenge(sequence)
	fmt.Printf("Boss Baby is %q\n", result)
}
