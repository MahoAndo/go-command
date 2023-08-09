package cmd

import (
	"goCommand/conf"

	"github.com/spf13/cobra"
)

type Root struct {
	Config *conf.Config
}

func NewRoot() (*Root, error) {
	// Load a config
	c, err := conf.Load()
	if err != nil {
		return nil, err
	}
	a := &Root{
		Config: c,
	}
	return a, nil
}

func (r Root) Build() *cobra.Command {

	cc := &cobra.Command{
		Use:   "bookCmd",
		Short: "bookCmd's root doesn't have any funcion",
	}
	return cc

}
