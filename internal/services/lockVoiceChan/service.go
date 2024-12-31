package lockVoiceChan

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

type Service struct {
	discord discord
}

func New(discord discord) *Service {
	return &Service{
		discord: discord,
	}
}

func (s *Service) Exec(
	ctx context.Context, i *discordgo.Interaction, limit *int,
) error {
	voiceState, err := s.discord.UserVoiceState(
		ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
	)
	if err != nil {
		return fmt.Errorf("failed to get user voice channel: %w", err)
	}

	if voiceState == nil {
		err = s.discord.Respond(ctx, i, responses.UserNotInVoiceChannel())
		if err != nil {
			return fmt.Errorf("failed to respond: %w", err)
		}

		return nil
	}

	limit, err = s.verifyLimit(
		ctx, limit, models.ID(i.GuildID), models.ID(voiceState.ChannelID),
	)
	if err != nil {
		return err
	}

	channel, err := s.discord.Channel(ctx, models.ID(voiceState.ChannelID))
	if err != nil {
		return fmt.Errorf("failed to get voice channel: %w", err)
	}

	if *limit > commands.LockOptionLimitMaxValue {
		err = s.discord.Respond(ctx, i, responses.LockSpecifyLimit(channel))
		if err != nil {
			return fmt.Errorf("failed to respond: %w", err)
		}

		return nil
	}

	err = s.discord.EditChannel(
		ctx,
		models.ID(voiceState.ChannelID),
		&discordgo.ChannelEdit{UserLimit: *limit},
	)
	if err != nil {
		return fmt.Errorf("failed to edit channel user limit: %w", err)
	}

	err = s.discord.Respond(ctx, i, responses.LockDone(channel, *limit))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}

func (s *Service) verifyLimit(
	ctx context.Context, limit *int, guildID models.ID, channelID models.ID,
) (*int, error) {
	if limit != nil {
		return limit, nil
	}

	users, err := s.discord.VoiceChannelUsers(ctx, guildID, channelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get voice channel users: %w", err)
	}

	return ptr.To(len(users)), nil
}