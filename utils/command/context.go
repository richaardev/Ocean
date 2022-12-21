package command

import "github.com/disgoorg/disgo/events"

type CommandContext struct {
	Event *events.ApplicationCommandInteractionCreate
}

func (ctx *CommandContext) IsSubCommand() bool {
	return ctx.Event.SlashCommandInteractionData().SubCommandName != nil
}

func (ctx *CommandContext) IsSubCommandGroup() bool {
	return ctx.Event.SlashCommandInteractionData().SubCommandGroupName != nil
}