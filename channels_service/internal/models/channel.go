package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"strings"
	"time"
)

type StringArray []string

type Channel struct {
	ID            string                 `db:"id" gorm:"primaryKey;autoIncrement:false"`
	Name          string                 `db:"name" gorm:"not null"`
	Type          channelsv1.ChannelType `db:"type" gorm:"not null;default:0"`
	GuildID       sql.NullString         `db:"guild_id"`
	Position      sql.NullInt32          `db:"position"`
	Topic         sql.NullString         `db:"topic"`
	Icon          sql.NullString         `db:"icon"`
	UserLimit     int32                  `db:"user_limit" gorm:"default:-1"`
	Recipients    StringArray            `db:"recipients" gorm:"default:[];type:VARCHAR(255)"`
	OwnerID       string                 `db:"owner_id" gorm:"not null"`
	ParentID      sql.NullString         `db:"parent_id"`
	Permissions   sql.NullString         `db:"permissions" gorm:"not null"`
	LastMessageID sql.NullString         `db:"last_message_id"`
	IsNSFW        bool                   `db:"nsfw" gorm:"not null;default:false"`
	IsVoice       bool                   `db:"voice" gorm:"not null;default:false"`
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

func (a *StringArray) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*a = strings.Split(string(bytes), ",")
	return nil
}

func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}
	return strings.Join(a, ","), nil
}
