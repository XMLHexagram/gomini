package commands

import "github.com/spf13/cobra"

func Execute(args []string) (resp Response) {
	gominiCmd := newCommandsBuilder().addAll().build()
	cmd := gominiCmd.getCommand()
	cmd.SetArgs(args)

	_, err := cmd.ExecuteC()
	if err != nil {
		resp.Err = err
		return
	}

	return
}

type commandsBuilder struct {
	commands []cmder
}

func newCommandsBuilder() *commandsBuilder {
	return &commandsBuilder{}
}

func (b *commandsBuilder) addAll() *commandsBuilder {
	b.addCommands(
		b.newServerCmd(),
		newVersionCmd(),
		b.newNewCmd(),
		b.newNewBuild(),
		b.newNewSiteCmd(),
		b.newNewExec(),
	)
	return b
}

func (b *commandsBuilder) addCommands(commands ...cmder) *commandsBuilder {
	b.commands = append(b.commands, commands...)
	return b
}

func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

func (b *commandsBuilder) newBuilderCmd(cmd *cobra.Command) *baseBuilderCmd {
	bcmd := &baseBuilderCmd{commandsBuilder: b, baseCmd: &baseCmd{cmd: cmd}}
	//bcmd.hugoBuilderCommon.handleFlags(cmd)
	return bcmd
}

func (b *commandsBuilder) build() *gominiCmd {
	h := b.newGominiCmd()
	addCommands(h.getCommand(), b.commands...)
	return h
}

func addCommands(root *cobra.Command, commands ...cmder) {
	for _, command := range commands {
		cmd := command.getCommand()
		if cmd == nil {
			continue
		}
		root.AddCommand(cmd)
	}
}

type baseBuilderCmd struct {
	*baseCmd
	*commandsBuilder
}

type baseCmd struct {
	cmd *cobra.Command
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}

type cmder interface {
	getCommand() *cobra.Command
}

type gominiCmd struct {
	*baseBuilderCmd

	// Need to get the sites once built.
	//c *commandeer
}

type Response struct {
	Err error
}
