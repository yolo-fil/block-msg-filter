package filter

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("yolo-fil")
var timestmp = time.Unix(0, 0)
var cfg map[string][]uint64

func init() {
	log.Warnf("yolo-fil: active")
}

func FilterMsgList(msgs []*types.SignedMessage) []*types.SignedMessage {
	log.Infow("yolo-fil: filter msg list")
	file, err := os.Stat(os.Getenv("YOLO_FIL_CONFIG_PATH"))
	if err != nil {
		log.Errorf("yolo-fil: error stat'ing file")
		return msgs
	}
	modifiedtime := file.ModTime()
	if modifiedtime.After(timestmp) {
		log.Infow("yolo-fil: loading config file")
		content, err := ioutil.ReadFile(os.Getenv("YOLO_FIL_CONFIG_PATH"))
		if err != nil {
			log.Errorf("yolo-fil: error opening file")
			return msgs
		}
		cfg := map[string][]uint64{}
		err = json.Unmarshal(content, &cfg)
		if err != nil {
			log.Errorf("yolo-fil: error unmarshalling file")
			return msgs
		}
		timestmp = modifiedtime
	}
	result := make([]*types.SignedMessage, 0, len(msgs))
	log.Infow("yolo-fil: cfg loaded")
	for _, msg := range msgs {
		log.Infow("yolo-fil: check msg", "cid", msg.Message.Cid().String())
		if val, ok := cfg[msg.Message.To.String()]; ok {
			log.Infow("yolo-fil: filter msg", "cid", msg.Message.Cid().String(), "to", msg.Message.To.String())
			if val[0] == 0 {
				log.Warnw("yolo-fil: filter applied", "cid", msg.Message.Cid().String(), "to", msg.Message.To.String(), "method", 0)
				break
			} else {
				okm := true
				for _, m := range val {
					if ok && uint64(msg.Message.Method) == m {
						log.Warnw("yolo-fil: filter applied", "cid", msg.Message.Cid().String(), "to", msg.Message.To.String(), "method", m)
						okm = false
					}
				}
				if okm {
					result = append(result, msg)
				}
			}
		} else {
			result = append(result, msg)
		}
	}
	return result
}
