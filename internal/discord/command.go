package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	Handle(ctx context.Context, i *discordgo.InteractionCreate) error
}

type Command struct {
	Description *discordgo.ApplicationCommand
	Handler     Handler
}