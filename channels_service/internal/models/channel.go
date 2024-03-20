package models

import (
	"database/sql"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

type ChannelType int32

const (
	GuildText ChannelType = iota
	DM
	GuildVoice
	GroupDM
	GuildCategory
)

type Channel struct {
	ID            int64          `db:"id" gorm:"primaryKey;autoIncrement:false"`
	Name          string         `db:"name" gorm:"unique;not null"`
	Type          ChannelType    `db:"type" gorm:"not null"`
	GuildID       int64          `db:"guild_id" gorm:"not null"`
	Position      int32          `db:"position" gorm:"not null"`
	Topic         sql.NullString `db:"topic"`
	Icon          sql.NullString `db:"icon"`
	UserLimit     int32          `db:"user_limit" gorm:"default:-1"`
	Recipients    []int64        `db:"recipients" gorm:"type:bigint[]"`
	OwnerID       int64          `db:"owner_id" gorm:"not null"`
	ParentID      int64          `db:"parent_id"`
	Permissions   string         `db:"permissions" gorm:"not null"`
	LastMessageID int64          `db:"last_message_id"`
	IsNSFW        bool           `db:"nsfw" gorm:"not null"`
	IsVoice       bool           `db:"voice" gorm:"not null"`
	CreatedAt     time.Time      `db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `db:"updated_at" gorm:"autoUpdateTime"`
}

func (channel *Channel) BeforeCreate(tx *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	channel.ID = node.Generate().Int64()

	return nil
}
