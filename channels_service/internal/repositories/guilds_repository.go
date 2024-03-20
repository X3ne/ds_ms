package repositories

import (
	"context"

	"github.com/X3ne/ds_ms/channels_service/internal/models"
	"gorm.io/gorm"
)

type ChannelsRepository struct {
	db *gorm.DB
}

func NewChannelsRepository(db *gorm.DB) *ChannelsRepository {
	return &ChannelsRepository{db: db}
}

func (gr *ChannelsRepository) CreateChannel(ctx context.Context, guild *models.Channel) error {
	if err := gr.db.Create(guild).Error; err != nil {
		return err
	}
	return nil
}

func (gr *ChannelsRepository) GetChannelByID(ctx context.Context, guildID int64) (*models.Channel, error) {
	var guild models.Channel
	if err := gr.db.First(&guild, guildID).Error; err != nil {
		return nil, err
	}
	return &guild, nil
}

func (gr *ChannelsRepository) UpdateChannel(ctx context.Context, guild *models.Channel) error {
	if err := gr.db.Save(guild).Error; err != nil {
		return err
	}
	return nil
}

func (gr *ChannelsRepository) DeleteChannel(ctx context.Context, guildID int64) error {
	if err := gr.db.Delete(&models.Channel{}, guildID).Error; err != nil {
		return err
	}
	return nil
}
