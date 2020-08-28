package cmd

import (
	"fmt"
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"log"
)

var getCmd = &cli.Command{
	Name:      "get",
	Usage:     "get value by key",
	UsageText: "owl get [key]",
	Action: func(c *cli.Context) error {
		var key = c.Args().Get(0)
		owl.SetAddr([]string{endPoint})
		owl.SetKey(key)

		v, err := owl.Get()
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("value: ", v)
		return nil
	},
}
