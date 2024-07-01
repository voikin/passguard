package password

import (
	"encoding/json"
	"os"
)

func loadCommonPatterns(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var patterns struct {
		Patterns []string `json:"patterns"`
	}
	err = json.Unmarshal(data, &patterns)
	if err != nil {
		return nil, err
	}

	return patterns.Patterns, nil
}
