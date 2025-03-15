package hibana

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"slices"
)

type channel struct {
	ID string `json:"id"`
}

type notifiedChannels struct {
	Channels []channel `json:"channels"`
}

func (nc *notifiedChannels) removeChannel(channel channel) {
	var index int
	var found bool
	for i, c := range nc.Channels {
		if channel == c {
			index = i
			found = true
			break
		}
	}
	if found {
		nc.Channels = slices.Delete(nc.Channels, index, index+1)
	}
}

func (nc *notifiedChannels) readFromFile(file string) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(data, &nc)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (nc *notifiedChannels) writeToFile(file string) error {
	data, err := json.Marshal(nc)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
