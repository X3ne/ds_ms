package repositories

import (
	"context"

	"github.com/X3ne/ds_ms/guilds_service/internal/models"
	"gorm.io/gorm"
)

type GuildRepository struct {
	db *gorm.DB
}

func NewGuildRepository(db *gorm.DB) *GuildRepository {
	return &GuildRepository{db: db}
}

func (gr *GuildRepository) CreateGuild(ctx context.Context, guild *models.Guild) error {
	if err := gr.db.Create(guild).Error; err != nil {
		return err
	}
	return nil
}

func (gr *GuildRepository) GetGuildByID(ctx context.Context, guildID string) (*models.Guild, error) {
	var guild models.Guild
	if err := gr.db.First(&guild, guildID).Error; err != nil {
		return nil, err
	}
	return &guild, nil
}

func (gr *GuildRepository) UpdateGuild(ctx context.Context, guild *models.Guild) error {
	if err := gr.db.Save(guild).Error; err != nil {
		return err
	}
	return nil
}

func (gr *GuildRepository) DeleteGuild(ctx context.Context, guildID string) error {
	if err := gr.db.Delete(&models.Guild{}, guildID).Error; err != nil {
		return err
	}
	return nil
}
