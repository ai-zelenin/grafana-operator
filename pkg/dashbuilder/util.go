package dashbuilder

import (
	"encoding/json"
	"fmt"
)

func DumpJSON(obj any) string {
	data, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return fmt.Errorf("dump json error %v", err).Error()
	}
	return string(data)
}
