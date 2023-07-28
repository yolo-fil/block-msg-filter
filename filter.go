package filter

import (
	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("yolo-fil")

// // Hello returns a greeting for the named person.
// func Hello(name string) {
// 	// Return a greeting that embeds the name in a message.
// 	log.Infow("yolo-fil", name)

// }

func Filter[T any](slice []T, predicate func(T) bool) []T {
	log.Infow("yolo-fil filtering")
	result := make([]T, 0, len(slice))

	for _, element := range slice {
		if predicate(element) {
			result = append(result, element)
		}
	}

	return result
}

func PredicateMsgs(msg *types.SignedMessage) bool {
	log.Infow("yolo-fil filtering msg:", msg.Message.From.String())
	return true
}
