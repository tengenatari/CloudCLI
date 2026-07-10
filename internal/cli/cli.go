package cli

import (
	"CloudCLI/config"
	"fmt"

	"github.com/spf13/cobra"
)

type ServiceInterface interface {
	ProfileCreate(args map[string]string) error
	ProfileDelete(args map[string]string) error
	ProfileList(args map[string]string) error
	ProfileGet(args map[string]string) error
}

type CLI struct {
	profileService ServiceInterface
	execRoot       *cobra.Command
	root           *cobra.Command
	args           map[string]*string
	handlers       map[string]func(args map[string]string) error
}

func NewCLI(config *config.Config, profileService ServiceInterface) (*CLI, error) {
	execRoot := &cobra.Command{
		Use:   "manage",
		Short: "Cloud manage CLI tool",
		Long:  "A CLI tool for managing cloud",
	}

	root := &cobra.Command{
		Use:   "profile",
		Short: "managing profiles",
		Long:  "you can manage profiles: delete, create, list, get",
	}

	cli := &CLI{profileService: profileService, execRoot: execRoot, root: root, args: make(map[string]*string)}

	cli.registerHandlers()

	err := cli.registerCommands(config)

	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (cli *CLI) registerHandlers() {
	cli.handlers = map[string]func(map[string]string) error{
		"create": cli.profileService.ProfileCreate,
		"delete": cli.profileService.ProfileDelete,
		"list":   cli.profileService.ProfileList,
		"get":    cli.profileService.ProfileGet,
	}
}

func (cli *CLI) registerCommands(config *config.Config) error {

	for argName := range config.Args {
		var flagValue string
		cli.args[argName] = &flagValue
	}

	for commandName, commandStruct := range config.Commands {
		handler, exists := cli.handlers[commandName]
		if !exists {
			return fmt.Errorf("handler for command '%s' not found", commandName)
		}
		command := &cobra.Command{
			Use:   commandName,
			Short: commandStruct.Description,
			Long:  commandStruct.Description,
			Run: func(cmd *cobra.Command, args []string) {
				flags := make(map[string]string)
				for name, ptr := range cli.args {
					if ptr != nil {
						flags[name] = *ptr
					}
				}
				if err := handler(flags); err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			},
		}

		for _, argName := range commandStruct.ArgNames {

			argStruct, ok := config.Args[argName]

			if !ok {
				return fmt.Errorf("argument '%s' not found", argName)
			}

			command.Flags().StringVarP(cli.args[argName], argName, argStruct.ShortName, "", argStruct.Description)
			err := command.MarkFlagRequired(argName)

			if err != nil {
				return err
			}

		}
		cli.root.AddCommand(command)
	}
	cli.execRoot.AddCommand(cli.root)
	return nil
}

func (cli *CLI) Run() error {
	return cli.execRoot.Execute()
}
