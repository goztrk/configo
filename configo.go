package configo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"sync"
)

type IConfigo interface{}

type configo struct {
	configs map[string]IConfigo
}

var (
	instance   *configo
	once       sync.Once
	configPath string
)

func Configo() *configo {
	once.Do(func() {
		instance = &configo{}
	})

	return instance
}

func (c *configo) SetPath(p string) (err error) {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(p, 0755)
	}

	if err != nil {
		return err
	}

	configPath = p
	return nil
}

// Add config to the list.
func (c *configo) Add(conf string, s IConfigo) error {
	if _, ok := c.configs[conf]; ok {
		return errors.New("There is already a definition for given config")
	}

	c.configs[conf] = s

	return nil
}

// Get the given config struct. Need to be casted to the original struct. Eg: configo.Get(conf).(*config)
func (c *configo) Get(conf string) (IConfigo, error) {
	if _, ok := c.configs[conf]; !ok {
		return nil, errors.New(fmt.Sprintf("[%s] config could not be found", conf))
	}

	return c.configs[conf], nil
}

// Save the given config.
func (c *configo) Save(conf string) (err error) {
	if _, ok := c.configs[conf]; !ok {
		return errors.New(fmt.Sprintf("[%s] config could not found", conf))
	}

	return c.save(conf)
}

// SaveAll configs to their files
func (c *configo) SaveAll() (err error) {
	for conf := range c.configs {
		err = c.save(conf)
		if err != nil {
			return err
		}
	}

	return err
}

func (c *configo) save(conf string) (err error) {
	data, err := json.MarshalIndent(c.configs[conf], "", "  ")

	if err != nil {
		return err
	}

	filename := path.Join(configPath, conf+".json")

	err = os.WriteFile(filename, []byte(data), 0644)

	return err
}

// Load given config from file
func (c *configo) Load(conf string) (err error) {
	if _, ok := c.configs[conf]; !ok {
		return errors.New(fmt.Sprintf("[%s] config could not found", conf))
	}

	return c.load(conf)
}

func (c *configo) load(conf string) (err error) {
	filename := path.Join(configPath, conf+".json")

	if _, err := os.Stat(filename); err == nil {
		file, err := os.ReadFile(filename)

		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(file), c.configs[conf])

		return err
	} else if errors.Is(err, os.ErrNotExist) {
		return c.save(conf)
	}

	return err
}
