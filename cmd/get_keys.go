package cmd

import (
	"fmt"
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"log"
)

var getKeysCmd = &cli.Command{
	Name:      "get_keys",
	Usage:     "get keys by prefix",
	UsageText: "owl get_keys [prefix]",
	Action: func(c *cli.Context) error {
		var key = c.Args().Get(0)
		owl.SetAddr([]string{endPoint})
		owl.SetKey(key)

		v, err := owl.GetKeys(key)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("value: ", v)
		return nil
	},
}
