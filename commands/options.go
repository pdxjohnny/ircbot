package commands

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

var created = false

var sharedOptions = map[string]interface{}{
	"server": map[string]interface{}{
		"value": "irc.freenode.net:6697",
		"help":  "IRC server to connect to",
	},
	"tls": map[string]interface{}{
		"value": true,
		"help":  "Connect to the server using tls",
	},
	"nick": map[string]interface{}{
		"value": "golangirc_:id",
		"help":  "Nick to join with",
	},
	"name": map[string]interface{}{
		"value": "golangirc_:id",
		"help":  "Name to be refered to as",
	},
}

// ConfigOptionsMap are command line options
var ConfigOptionsMap = map[string]interface{}{
	"echo": sharedOptions,
}

// ConfigOptions generates the command line options
func ConfigOptions() map[string]interface{} {
	if !created {
		created = true
	} else {
		return ConfigOptionsMap
	}
	id := ""
	size := 6
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		return nil
	}
	id = base64.URLEncoding.EncodeToString(rb)
	replaceID(string(id), &ConfigOptionsMap)
	return ConfigOptionsMap
}

func replaceID(id string, optPointer interface{}) {
	options := *(optPointer.(*map[string]interface{}))
	for index, item := range options {
		switch value := item.(type) {
		case map[string]interface{}:
			replaceID(id, &value)
		case string:
			options[index] = strings.Replace(
				value,
				":id",
				id,
				-1,
			)
		}
	}
}
