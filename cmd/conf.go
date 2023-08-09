package cmd

import (
	"fmt"
	"goCommand/conf"

	"github.com/spf13/cobra"
)

type Conf struct {
	Config *conf.Config
}

func NewConf(c *conf.Config) (*Conf, error) {
	a := &Conf{
		Config: c,
	}
	return a, nil
}

func (c Conf) Conf() {
	fmt.Printf("%#v", c.Config)
}

func (c Conf) Build() *cobra.Command {

	cc := &cobra.Command{
		Use:   "conf",
		Short: "Print the configuration of bookCmd",
		Run: func(cmd *cobra.Command, args []string) {
			c.Conf()
		},
	}
	return cc

}
