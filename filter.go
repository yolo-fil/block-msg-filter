package filter

import (
	"fmt"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("yolo-fil")

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
