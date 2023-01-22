package settings

import (
	"context"
	"fmt"
	"io"
	"os"
)

const _configFilePath = "settings/config.json"

type (
	// Implement a file based config loader
	FileConfigLoader struct {
		filepath     string
		allowMissing bool
	}

	FileConfigOpt func(*FileConfigLoader)
)

var _ ConfigReader = (*FileConfigLoader)(nil)

// NewFileConfigLoader creates a new config loader for file based configuration
func NewFileConfigLoader(opts ...FileConfigOpt) (*FileConfigLoader, error) {
	retval := &FileConfigLoader{
		filepath: _configFilePath,
	}

	for _, o := range opts {
		o(retval)
	}

	return retval, nil
}

// WithAllowFileMissing allows the config loader to return a default config when the file is missing
func WithAllowFileMissing() FileConfigOpt {
	return func(l *FileConfigLoader) {
		l.allowMissing = true
	}
}

// WithFilePath sets the file path for the config loader
func WithFilePath(filepath string) FileConfigOpt {
	return func(l *FileConfigLoader) {
		l.filepath = filepath
	}
}

// Load implements ConfigReader
func (l *FileConfigLoader) Load(context.Context) (Config, error) {
	// check if the file exists
	_, err := os.Stat(l.filepath)
	if os.IsNotExist(err) {
		if !l.allowMissing {
			return Config{}, fmt.Errorf("file %s does not exist: %w", l.filepath, err)
		}

		// when missing file is allowed, returns a default config when the file is missing
		return Config{}, nil
	}

	file, err := os.Open(l.filepath)
	if err != nil {
		return Config{}, fmt.Errorf("fail to open file %s: %w", l.filepath, err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return Config{}, fmt.Errorf("fail to read file %s: %w", l.filepath, err)
	}

	c, err := parseConfig(content)
	if err != nil {
		return Config{}, fmt.Errorf("fail to unmarshal file %s: %w", l.filepath, err)
	}

	return c, nil
}
