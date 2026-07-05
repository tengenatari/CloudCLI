package cli

import (
	"CloudCLI/config"
	"fmt"

	"github.com/spf13/cobra"
)

type ServiceInterface interface {
	ProfileCreate(name string, user string, profile string) error
	ProfileDelete(name string) error
	ProfileList() error
	ProfileGet(name string) error
	Help() error
}

type CLI struct {
	profileService ServiceInterface
	root           *cobra.Command
	args           map[string]*string
}

func NewCLI(config *config.Config, profileService ServiceInterface) (*CLI, error) {
	root := &cobra.Command{
		Use:   "cloudcli",
		Short: "Cloud CLI tool",
		Long:  "A CLI tool for managing cloud profiles",
	}

	cli := &CLI{profileService: profileService, root: root, args: make(map[string]*string)}
	for commandName, commandStruct := range config.Commands {

		command := &cobra.Command{
			Use:   commandName,
			Short: commandStruct.Description,
			Long:  commandStruct.Description,
		}

		for _, argName := range commandStruct.ArgNames {

			argStruct, ok := config.Args[argName]

			if !ok {
				return nil, fmt.Errorf("argument '%s' not found", argName)
			}

			var flagValue string
			cli.args[argName] = &flagValue

			command.Flags().StringVarP(cli.args[argName], argName, argStruct.ShortName, "", argStruct.Description)
			err := command.MarkFlagRequired(argName)

			if err != nil {
				return nil, err
			}
		}
		root.AddCommand(command)
	}
	return cli, nil
}
