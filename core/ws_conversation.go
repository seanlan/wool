package core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// makeClientKey 根据appkey和tag生成client唯一标示
func MakeConversationID(appKey, from, to string, event int) string {
	var source string
	if event == 1 {
		keys := []string{from, to}
		sort.Strings(keys)
		source = strings.Join(keys, ":")
	} else {
		source = to
	}
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s:%s", appKey, source)))
	return hex.EncodeToString(h.Sum(nil))
}
