package cmd

import (
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
)

var putCmd = &cli.Command{
	Name:  "put",
	Usage: "put",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint,e",
			Value:       "http://127.0.0.1:2379",
			Usage:       "",
			Destination: &endPoint,
		},
	},
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
