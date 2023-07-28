package filter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
)

type cfgFormat map[string][]int

var log = logging.Logger("yolo-fil")
var timestmp = time.Unix(0, 0)
var cfg cfgFormat

func Filter[T any](slice []T, predicate func(T, cfgFormat) bool) []T {
	log.Infow("yolo-fil filtering")
	file, err := os.Stat(os.Getenv("YOLO_FIL_CONFIG_PATH"))
	if err != nil {
		log.Errorf("yolo-fil: error stat'ing file")
		return slice
	}
	modifiedtime := file.ModTime()
	if modifiedtime.After(timestmp) {
		log.Infow("yolo-fil filtering: loading config file")
		content, err := ioutil.ReadFile(os.Getenv("YOLO_FIL_CONFIG_PATH"))
		if err != nil {
			log.Errorf("yolo-fil: error opening file")
			return slice
		}
		cfg := cfgFormat{}
		err = json.Unmarshal(content, &cfg)
		if err != nil {
			log.Errorf("yolo-fil: error unmarshalling file")
			return slice
		}
		timestmp = modifiedtime
	}
	result := make([]T, 0, len(slice))
	for _, element := range slice {
		if predicate(element, cfg) {
			result = append(result, element)
		}
	}
	return result
}

func PredicateMsgs(msg *types.SignedMessage, cfg cfgFormat) bool {
	log.Infow("yolo-fil filtering msg:", msg.Message.From.String())

	fmt.Printf("hey: %x", cfg)
	return true
}
