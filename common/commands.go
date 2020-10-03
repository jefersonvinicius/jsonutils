package common

import "encoding/json"

// Command interface
type Command interface {
	Execute() string
}

// MinifyCommand minify json
type MinifyCommand struct {
	InputData []byte
}

func (command MinifyCommand) Execute() string {
	var unmarshaled interface{}
	json.Unmarshal(command.InputData, &unmarshaled)
	content, _ := json.Marshal(unmarshaled)
	return string(content)
}

// FormatCommand format json
type FormatCommand struct {
	InputData []byte
}

func (command FormatCommand) Execute() string {
	var jsonObject map[string]interface{}
	json.Unmarshal(command.InputData, &jsonObject)
	content, _ := json.MarshalIndent(jsonObject, "", "	")
	return string(content)
}
