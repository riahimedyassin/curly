package args

import (
	"github.com/riahimedyassin/curly/internal/dto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ComponentArgsParser struct {
	cmd  *cobra.Command
	args []string
}

func NewComponentArgsParser(cmd *cobra.Command, args []string) *ComponentArgsParser {
	return &ComponentArgsParser{
		cmd:  cmd,
		args: args,
	}
}

func (p *ComponentArgsParser) Parse() (*dto.ComponentArgs, error) {
	v := viper.New()
	v.BindPFlags(p.cmd.Flags())
	v.Set("name", p.args[0])
	var value dto.ComponentArgs
	if err := v.Unmarshal(&value); err != nil {
		return nil, err
	}
	return &value, nil
}
