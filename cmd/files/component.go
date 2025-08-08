package cmd_files

import "github.com/spf13/cobra"

func NewComponentCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "component",
		Short: "Generate components",
		Long: `Generate component based on the configs and user args.
		The template configuration are the weakest configs. The values should be overriden by the team configuration.
		User configuration (specified through args) will be valid on if they does not interfer with the teams policy.
		Example : 
			- Team component naming convention : PascalCase.
			- User tries to enter a camelCase name.
			-> Terminal : Invalid args, use a PascalCase name for your component. 
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
