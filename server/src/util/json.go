package util

import (
	"encoding/json"
	"project/zj"
)

// DumpJSON ...
func DumpJSON(a any) {

	if s, ok := a.([]byte); ok {
		json.Unmarshal(s, &a)
	}

	ab, err := json.MarshalIndent(a, ``, "\t")
	if err != nil {
		zj.WF(`json marshal err: %s`, err)
		return
	}
	zj.J(string(ab))
}
