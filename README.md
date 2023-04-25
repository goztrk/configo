# configo - Lightweight Config Manaer

configo is a lightweight configuration manager that requires no external library.

## Install
```shell
go get github.com/goztrk/configo
```

## Why configo?
It is very lightweight and focused on saving and/or loading configuration files.
You can update your configuration in runtime and save it when closing the program
or whenever you want. It is perfect if you have multiple packages that uses
seperate config files. Currently it only supports JSON. YAML might be added
also in the future.

## Usage
```go
package main

import (
    "os"

    "github.com/goztrk/configo"
)

type config struct {
    addr string `json:"server-address"`
    port string `json:"server-port"`
}

var conf *config

func main() {
    // These are default values
    conf = &config{
        addr: "127.0.0.1",
        port: "80",
    }

    _ = configo.Add("server", conf, false)

    // Loads config from file and if file does not exists,
    // it saves the file with default values.
    _ = configo.Load("server")

    // now conf has updated values from `server.json` file
    // config/server.json source:
    // {
    //   "server-address": "127.0.0.1",
    //   "server-port": "80",
    // }
    
    // Get configuration struct from configo:
    c, _ := configo.Get("server").(*config)
}

```

## Methods
### `configo.Add(conf string, s IConfigo) error`
Adds given struct to the map. The defined values of struct can be count as default values.

### `configo.Load(conf string) error`
Loads given configuration group from file. If the file does not exists, it creates and saves with the
current values.

### `configo.Save(conf string) error`
Saves given configuration group to a file. Overrides existing file with new one.

### `configo.SaveAll() error`
Saves all existing configuration groups to their files.

### `configo.SetPath(path string) error`
Sets the folder that holds configuration files. Default path is `./config/`

## Planned Features
- [ ] Write tests
- [ ] Prefix for file names
- [ ] Ability to change file extensions
- [ ] Ability to choose between JSON and YAML formats
- [ ] Create configuration folder if it not exists
