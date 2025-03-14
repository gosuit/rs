# RS

rs is a simple and convenient wrapper around the <a href="github.com/redis/go-redis/v9">go-redis</a>, providing an easy way to connect to a Redis server in your Go applications. This library abstracts the connection setup and allows you to focus on using Redis without dealing with boilerplate code.

## Installation

```zsh
go get github.com/gosuit/rs
```

## Features

• Easy configuration using environment variables or YAML files.

• Simplified client creation with automatic connection testing.

• Compatible with Redis v9.

## Usage

```golang
package main

import (
    "context"
    "log"

    "github.com/gosuit/rs"
)

func main() {
    ctx := context.Background()

    // Create a configuration
    cfg := &rs.Config{
        Host:     "localhost",
        Port:     6379,
        DBNumber: 0,
        Password: "password",
    }

    // Connect to Redis
    client, err := rs.New(ctx, cfg)
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    log.Println("Connected to Redis successfully!")

    // Use client...
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.