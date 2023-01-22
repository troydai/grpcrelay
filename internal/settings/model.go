package settings

import (
	"encoding/json"
)

// The logic of content loading and config model parsing are separated.

// parseConfig parses the input byte array into a Config struct.
func parseConfig(input []byte) (Config, error) {
	var c Config
	if err := json.Unmarshal(input, &c); err != nil {
		return c, err
	}

	return c, nil
}
