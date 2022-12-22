package command

import (
    "github.com/disgoorg/disgo/bot"
    "github.com/disgoorg/disgo/discord"
    "github.com/disgoorg/disgo/events"
)

type Context struct {
	Event *events.ApplicationCommandInteractionCreate
}

func (ctx *Context) GetSubCommandName() string {
    return *ctx.Event.SlashCommandInteractionData().SubCommandName
}

func (ctx *Context) IsSubCommand() bool {
	return ctx.Event.SlashCommandInteractionData().SubCommandName != nil
}

func (ctx *Context) IsSubCommandGroup() bool {
	return ctx.Event.SlashCommandInteractionData().SubCommandGroupName != nil
}

func (ctx *Context) GetSelf() (discord.OAuth2User, bool) {
    return ctx.Event.Client().Caches().GetSelfUser()
}

func (ctx *Context) Client() bot.Client {
    return ctx.Event.Client()
}