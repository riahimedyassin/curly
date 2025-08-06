package cmd

import (
	"fmt"

	"github.com/riahimedyassin/curly/internal/config"
	"github.com/spf13/cobra"
)

func NewRootCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "rg",
		Short: "Curly - V0.0.1",
		Long:  "Curly - Ultimate React toolkit.",
		RunE: func(cmd *cobra.Command, args []string) error {
			configLoader := config.NewConfigLoader()
			config, err := configLoader.Load()
			if err != nil {
				return err
			}
			fmt.Printf("%v", config)
			return nil
		},
		Aliases: []string{
			"reactgenerator",
		},
	}
}
