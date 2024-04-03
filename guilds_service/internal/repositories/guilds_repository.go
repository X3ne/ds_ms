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

func (gr *GuildRepository) GetGuildMember(ctx context.Context, guildID string, userID string) (*models.GuildMember, error) {
	var member models.GuildMember
	if err := gr.db.Where("guild_id = ? AND user_id = ?", guildID, userID).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (gr *GuildRepository) GetGuildMembers(ctx context.Context, guildID string) ([]*models.GuildMember, error) {
	var members []*models.GuildMember
	if err := gr.db.Where("guild_id = ?", guildID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (gr *GuildRepository) SearchGuildMembers(ctx context.Context, guildID string, query string) ([]*models.GuildMember, error) {
	var members []*models.GuildMember
	if err := gr.db.Where("guild_id = ? AND username LIKE ? OR nickname LIKE ?", guildID, "%"+query+"%", "%"+query+"%").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (gr *GuildRepository) AddGuildMember(ctx context.Context, member *models.GuildMember) error {
	if err := gr.db.Create(member).Error; err != nil {
		return err
	}
	return nil
}

func (gr *GuildRepository) UpdateGuildMember(ctx context.Context, member *models.GuildMember) error {
	if err := gr.db.Save(member).Error; err != nil {
		return err
	}
	return nil
}

func (gr *GuildRepository) DeleteGuildMember(ctx context.Context, guildID string, userID string) error {
	if err := gr.db.Where("guild_id = ? AND user_id = ?", guildID, userID).Delete(&models.GuildMember{}).Error; err != nil {
		return err
	}
	return nil
}

func (gr *GuildRepository) GetGuildBans(ctx context.Context, guildID string) ([]*models.GuildBan, error) {
	var bans []*models.GuildBan
	if err := gr.db.Where("guild_id = ?", guildID).Find(&bans).Error; err != nil {
		return nil, err
	}
	return bans, nil
}

func (gr *GuildRepository) GetGuildBan(ctx context.Context, guildID string, userID string) (*models.GuildBan, error) {
	var ban models.GuildBan
	if err := gr.db.Where("guild_id = ? AND user_id = ?", guildID, userID).First(&ban).Error; err != nil {
		return nil, err
	}
	return &ban, nil
}

func (gr *GuildRepository) AddGuildBan(ctx context.Context, ban *models.GuildBan) error {
	if err := gr.db.Create(ban).Error; err != nil {
		return err
	}
	return nil
}

func (gr *GuildRepository) DeleteGuildBan(ctx context.Context, guildID string, userID string) error {
	if err := gr.db.Where("guild_id = ? AND user_id = ?", guildID, userID).Delete(&models.GuildBan{}).Error; err != nil {
		return err
	}
	return nil
}
