package cmd

import (
	"fmt"
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"log"
)

var getCmd = &cli.Command{
	Name:        "get",
	Usage:       "get value by key",
	UsageText:   "owl get [key]",
	Description: "the [key] what you want value at the etcd",
	Action: func(c *cli.Context) error {
		var key = c.Args().Get(0)
		owl.SetRemoteAddr([]string{endPoint})

		v, err := owl.GetRemote(key)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("value: ", v)
		return nil
	},
}
