package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	endPoint string
	rootCmd  = cli.NewApp()
	tasks    = []string{"get", "put", "version"}
)

func init() {

	rootCmd.Usage = "owl"
	rootCmd.EnableBashCompletion = true
	rootCmd.Version = ""
	rootCmd.HideVersion = true
	rootCmd.Commands = []*cli.Command{
		getCmd,
		putCmd,
		versionCmd,
	}
	rootCmd.BashComplete = func(c *cli.Context) {
		if c.NArg() > 0 {
			fmt.Println("123")
			return
		}
		for _, v := range tasks {
			fmt.Println(v)
		}
	}

}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Run(os.Args); err != nil {
		panic(err)
	}
}
