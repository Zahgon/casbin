package config

import (
	"bufio"
	"bytes"
)

var (
	DEFAULT_SECTION = "default"

	DEFAULT_COMMENT = []byte{'#'}

	DEFAULT_COMMENT_SEM = []byte{';'}

	DEFAULT_MULTI_LINE_SEPARATOR = []byte{'\\'}
)

type ConfigInterface interface {
	String(key string) string
	Strings(key string) []string
	Bool(key string) (bool, error)
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Float64(key string) (float64, error)
	Set(key string, value string) error
}

type Config struct {
	data map[string]map[string]string
}

func NewConfig(confName string) (ConfigInterface, error) {
	_ = "STUB: not implemented"
	return *new(ConfigInterface), nil
}

func NewConfigFromText(text string) (ConfigInterface, error) {
	_ = "STUB: not implemented"
	return *new(ConfigInterface), nil
}

func (c *Config) AddConfig(section string, option string, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Config) parse(fname string) (err error) { _ = "STUB: not implemented"; return nil }

func (c *Config) parseBuffer(buf *bufio.Reader) error { _ = "STUB: not implemented"; return nil }

func (c *Config) write(section string, lineNum int, b *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) Bool(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c *Config) Int(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Config) Int64(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Config) Float64(key string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Config) String(key string) string { _ = "STUB: not implemented"; return "" }

func (c *Config) Strings(key string) []string { _ = "STUB: not implemented"; return nil }

func (c *Config) Set(key string, value string) error { _ = "STUB: not implemented"; return nil }

func (c *Config) get(key string) string { _ = "STUB: not implemented"; return "" }
