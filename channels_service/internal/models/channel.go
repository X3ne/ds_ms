package models

import (
	"database/sql"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

type Channel struct {
	ID            string                 `db:"id" gorm:"primaryKey;autoIncrement:false"`
	Name          string                 `db:"name" gorm:"not null"`
	Type          channelsv1.ChannelType `db:"type" gorm:"not null"`
	GuildID       string                 `db:"guild_id" gorm:"not null"`
	Position      int32                  `db:"position" gorm:"not null"`
	Topic         sql.NullString         `db:"topic"`
	Icon          sql.NullString         `db:"icon"`
	UserLimit     int32                  `db:"user_limit" gorm:"default:-1"`
	Recipients    []string               `db:"recipients"`
	OwnerID       string                 `db:"owner_id" gorm:"not null"`
	ParentID      string                 `db:"parent_id"`
	Permissions   string                 `db:"permissions" gorm:"not null"`
	LastMessageID string                 `db:"last_message_id"`
	IsNSFW        bool                   `db:"nsfw" gorm:"not null"`
	IsVoice       bool                   `db:"voice" gorm:"not null"`
	CreatedAt     time.Time              `db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time              `db:"updated_at" gorm:"autoUpdateTime"`
}

func (channel *Channel) BeforeCreate(tx *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	channel.ID = node.Generate().String()

	return nil
}
