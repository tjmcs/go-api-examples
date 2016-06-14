package lib

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ProcessCmd(cmdMap map[string]string) string {
	// output message received
	fmt.Printf("Command Received: %+v\n", cmdMap)
	// sample process for string received
	jb, _ := json.Marshal(cmdMap)
	return strings.ToUpper(string(jb))
}
