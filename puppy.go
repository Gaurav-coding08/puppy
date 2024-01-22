package puppy

import (
	"github.com/Gaurav-coding08/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}
