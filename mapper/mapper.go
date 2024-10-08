package mapper

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Mapper struct {
	dic map[string]string
}

func (m *Mapper) BuildMap(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	bytesValue, _ := io.ReadAll(file)

	err = json.Unmarshal(bytesValue, &m.dic)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mapper) InsertInJson(alias, data string) {
	m.dic[alias] = data
	m.saveJson()
}

func (m *Mapper) saveJson() {
	json_data, _ := json.Marshal(m.dic)
	err := os.WriteFile("dictionary.json", json_data, 0644)
	if err != nil {
		panic(err)
	}
}

func (m *Mapper) GetCommand(alias string) (string, error) {
	if command, ok := m.dic[alias]; ok {
		return command, nil
	}

	return "", errors.New("Alias not found")
}

func (m *Mapper) GetDictionary() map[string]string {
	return m.dic
}
