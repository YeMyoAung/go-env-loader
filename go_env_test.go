package goenv

import (
	"log"
	"strings"
	"testing"
)

type Config struct {
	Name  string `json:"name" validate:"required"`
	Foods string `json:"foods" validate:"required"`
}

func (c *Config) FoodArray() []string {
	return strings.Split(c.Foods, ",")
}

func TestGoEnvLoader(t *testing.T) {
	config, err := NewGoEnv[Config](nil)
	if err != nil {
		t.Fatal(err)
	}
	if config == nil {
		t.Fatal("Config was null")
	}

	log.Println(config.Value)
}
