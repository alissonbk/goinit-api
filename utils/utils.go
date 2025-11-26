package utils

import (
	"os"
	"os/exec"
	"reflect"

	"github.com/charmbracelet/bubbles/list"
)

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// This function is considering that the interface implementation has the item name in the first position of the struct
func ExtractStringFromListItem(i list.Item) string {
	el := reflect.ValueOf(i)
	return el.Field(0).String()
}
