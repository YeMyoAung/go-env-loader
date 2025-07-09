package goenv

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type GoEnv[T any] struct {
	Value *T
}

type Args struct {
	FileName string
	Validate *validator.Validate
}

// parseFromFile reads the env file and parses it into a struct
// by splitting this function, we can also implement other loaders like NewConfigLoaderFromFlags
func parseFromFile[T any](
	source string,
) (*T, error) {
	file, err := os.Open(source)

	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file", err)
		}
	}(file)

	env, err := godotenv.Parse(file)

	if err != nil {
		return nil, err
	}
	buf, err := json.Marshal(env)

	if err != nil {
		return nil, err
	}

	var config T

	if err := json.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	err = godotenv.Load(source)

	if err != nil {
		return nil, err
	}
	return &config, nil
}

func parseFromEnv[T any]() (*T, error) {
	envMap := make(map[string]string)
	for _, v := range os.Environ() {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}
	buf, err := json.Marshal(envMap)
	if err != nil {
		return nil, err
	}

	var config T

	if err := json.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func NewGoEnv[T any](
	args *Args,
) (*GoEnv[T], error) {
	if args == nil {
		args = &Args{}
	}

	if args.Validate == nil {
		args.Validate = validator.New(
			validator.WithRequiredStructEnabled(),
		)
	}

	var err error
	var config *T

	if args.FileName != "" {
		log.Println("Loading env from ", args.FileName)
		config, err = parseFromFile[T](args.FileName)
	} else {
		log.Println("Loading env from env")
		config, err = parseFromEnv[T]()
	}

	if err != nil {
		return nil, err
	}

	if err := args.Validate.Struct(config); err != nil {
		return nil, err
	}

	return &GoEnv[T]{
		Value: config,
	}, nil
}
