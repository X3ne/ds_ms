package repositories

import (
	"context"

	"github.com/X3ne/ds_ms/channels_service/internal/models"
	"gorm.io/gorm"
)

type ChannelsRepository struct {
	db *gorm.DB
}

type SearchRequest struct {
	Limit  int32
	Before string
	After  string
	Around string
}

func constructSearchQuery(query *gorm.DB, opts ...SearchRequest) *gorm.DB {
	if len(opts) == 0 {
		return query
	}

	if opts[0].Limit > 0 {
		query = query.Limit(int(opts[0].Limit))
	} else {
		query = query.Limit(50)
	}

	if opts[0].Before != "" {
		query = query.Where("id < ?", opts[0].Before)
	}

	if opts[0].After != "" {
		query = query.Where("id > ?", opts[0].After)
	}

	if opts[0].Around != "" {
		query = query.Where("id = ?", opts[0].Around)
	}

	return query
}

func NewChannelsRepository(db *gorm.DB) *ChannelsRepository {
	return &ChannelsRepository{db: db}
}

func (r *ChannelsRepository) CreateChannel(ctx context.Context, channel *models.Channel) error {
	if err := r.db.Create(channel).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) GetChannelByID(ctx context.Context, channelID string) (*models.Channel, error) {
	var channel models.Channel
	if err := r.db.First(&channel, channelID).Error; err != nil {
		return nil, err
	}
	return &channel, nil
}

func (r *ChannelsRepository) UpdateChannel(ctx context.Context, channel *models.Channel) error {
	if err := r.db.Save(channel).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) DeleteChannel(ctx context.Context, channelID string) error {
	if err := r.db.Delete(&models.Channel{}, channelID).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) GetGuildChannels(ctx context.Context, guildID string) ([]models.Channel, error) {
	var channels []models.Channel
	if err := r.db.Where("guild_id = ?", guildID).Find(&channels).Error; err != nil {
		return nil, err
	}
	return channels, nil
}

func (r *ChannelsRepository) GetChannelMessages(ctx context.Context, channelID string, opts ...SearchRequest) ([]models.Message, error) {

	if len(opts) == 0 {
		opts = append(opts, SearchRequest{Limit: 50})
	}

	var messages []models.Message
	query := r.db.Where("channel_id = ?", channelID).Order("created_at desc")
	query = constructSearchQuery(query, opts[0])

	if err := query.Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *ChannelsRepository) GetMessageByID(ctx context.Context, messageID string) (*models.Message, error) {
	var message models.Message
	if err := r.db.First(&message, messageID).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *ChannelsRepository) CreateMessage(ctx context.Context, message *models.Message) error {
	if err := r.db.Create(message).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) UpdateMessage(ctx context.Context, message *models.Message) error {
	if err := r.db.Save(message).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) DeleteMessage(ctx context.Context, messageID string) error {
	if err := r.db.Delete(&models.Message{}, messageID).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) DeleteMessages(ctx context.Context, channelID string, messages []string) error {
	if err := r.db.Where("channel_id = ? AND id IN ?", channelID, messages).Delete(&models.Message{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) GetPinnedMessages(ctx context.Context, channelID string) ([]models.Message, error) {
	var messages []models.Message
	if err := r.db.Where("channel_id = ? AND pinned = ?", channelID, true).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *ChannelsRepository) PinMessage(ctx context.Context, messageID string) error {
	if err := r.db.Model(&models.Message{}).Where("id = ?", messageID).Update("pinned", true).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) UnpinMessage(ctx context.Context, messageID string) error {
	if err := r.db.Model(&models.Message{}).Where("id = ?", messageID).Update("pinned", false).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) AddGroupDMRecipient(ctx context.Context, channelID, userID string) error {
	if err := r.db.Model(&models.Channel{}).Where("id = ?", channelID).Update("recipients", gorm.Expr("array_append(recipients, ?)", userID)).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelsRepository) RemoveGroupDMRecipient(ctx context.Context, channelID, userID string) error {
	if err := r.db.Model(&models.Channel{}).Where("id = ?", channelID).Update("recipients", gorm.Expr("array_remove(recipients, ?)", userID)).Error; err != nil {
		return err
	}
	return nil
}
