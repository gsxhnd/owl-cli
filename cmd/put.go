package cmd

import (
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
)

var putCmd = &cli.Command{
	Name:      "put",
	Usage:     "read file then put value",
	UsageText: "owl put [key] [file_path]",
	Action: func(c *cli.Context) error {
		var key = c.Args().Get(0)
		var filePath = c.Args().Get(1)
		owl.SetAddr([]string{endPoint})
		yamlFile, err := ioutil.ReadFile(filePath)

		err = owl.Put(key, string(yamlFile))
		if err != nil {
			log.Panic(err)
		}
		return nil
	},
}
