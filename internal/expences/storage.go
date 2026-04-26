package expences

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save() error {
	data, err := json.Marshal(storage)
	if err != nil {
		return fmt.Errorf("save() marshal: %w", err)
	}
	err = os.WriteFile("expences.txt", data, 0644)
	if err != nil {
		return fmt.Errorf("save() write: %w", err)
	}
	return nil
}

func Load() error {
	data, err := os.ReadFile("expences.txt")
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("load() read: %w", err)
	}
	if len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, &storage); err != nil {
		return fmt.Errorf("load() unmarshal: %w", err)
	}
	return nil
}
