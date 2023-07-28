package filter

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
)

type cfgFormat map[string][]uint64

var log = logging.Logger("yolo-fil")
var timestmp = time.Unix(0, 0)
var cfg cfgFormat

func FilterMsgList(msgs []*types.SignedMessage) []*types.SignedMessage {
	log.Infow("yolo-fil filtering")
	file, err := os.Stat(os.Getenv("YOLO_FIL_CONFIG_PATH"))
	if err != nil {
		log.Errorf("yolo-fil: error stat'ing file")
		return msgs
	}
	modifiedtime := file.ModTime()
	if modifiedtime.After(timestmp) {
		log.Infow("yolo-fil filtering: loading config file")
		content, err := ioutil.ReadFile(os.Getenv("YOLO_FIL_CONFIG_PATH"))
		if err != nil {
			log.Errorf("yolo-fil: error opening file")
			return msgs
		}
		cfg := cfgFormat{}
		err = json.Unmarshal(content, &cfg)
		if err != nil {
			log.Errorf("yolo-fil: error unmarshalling file")
			return msgs
		}
		timestmp = modifiedtime
	}
	result := make([]*types.SignedMessage, 0, len(msgs))
	for _, element := range msgs {
		if val, ok := cfg[element.Message.To.String()]; ok {
			if val[0] == 0 {
				break
			} else {
				for _, m := range val {
					if uint64(element.Message.Method) == m {
						break
					} else {
						result = append(result, element)
					}
				}
			}
		} else {
			result = append(result, element)
		}
	}
	return result
}

func DefaultFilter(msg *types.SignedMessage, cfg cfgFormat) bool {
	log.Infow("yolo-fil DefaultFilter Cid:", msg.Message.Cid())
	return true
}
