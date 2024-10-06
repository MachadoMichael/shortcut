package mapper

import (
	"encoding/json"
	"io"
	"os"
)

type Mapper struct{}

func (m *Mapper) BuildMap(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytesValue, _ := io.ReadAll(file)

	var dic map[string]string
	err = json.Unmarshal(bytesValue, &dic)
	if err != nil {
		return nil, err
	}
	return dic, nil
}

func (m *Mapper) InsertInJson(alias, data string, dic map[string]string) {
	dic[alias] = data
	m.saveJson(dic)
}

func (m *Mapper) saveJson(dic map[string]string) {
	json_data, _ := json.Marshal(dic)
	err := os.WriteFile("dictionary.json", json_data, 0644)
	if err != nil {
		panic(err)
	}
}
