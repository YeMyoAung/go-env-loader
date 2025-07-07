# go-env-loader

A simple Go package to load environment variables from a `.env` file or the current environment into a struct, with validation support.

## Installation

```
go get github.com/yemyoaung/go-env-loader
```

## Usage

### 1. Define your config struct

```go
import "github.com/yemyoaung/go-env-loader"

type Config struct {
    Name  string `json:"name" validate:"required"`
    Foods string `json:"foods" validate:"required"`
}
```

> **Note:** Only string fields are supported. If you need to work with other data types (e.g., int, bool, slices), you can add custom Getter methods to your struct to handle conversion.

### 2. Load from environment variables

```go
config, err := go_env_loader.NewGoEnvLoader[Config](nil)
if err != nil {
    // handle error
}
fmt.Println(config.Value.Name)
```

### 3. Load from a .env file

```go
args := &go_env_loader.GoEnvLoaderArgs{
    FileName: ".env",
}
config, err := go_env_loader.NewGoEnvLoader[Config](args)
if err != nil {
    // handle error
}
fmt.Println(config.Value.Name)
```

### 4. Validation

By default, all struct fields with the `validate:"required"` tag are validated. You can provide a custom validator via `GoEnvLoaderArgs` if needed.

## Example .env file

```
NAME=John Doe
FOODS=apple,banana,orange
```

## License

MIT
