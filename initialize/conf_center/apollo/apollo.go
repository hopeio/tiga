package apollo

import (
	"encoding/json"
	"github.com/hopeio/tiga/utils/configor/apollo"
)

type Apollo struct {
	apollo.Config
}

func (e *Apollo) HandleConfig(handle func([]byte)) error {
	client, err := e.NewClient()
	if err != nil {
		return err
	}
	config, err := client.GetDefaultConfig()
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	handle(data)
	return nil
}
