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
	UserLimit     sql.NullInt32          `db:"user_limit"`
	Recipients    StringArray            `db:"recipients" gorm:"default:[];type:VARCHAR(255)"`
	OwnerID       sql.NullString         `db:"owner_id"`
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

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = []string{}
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan StringArray")
	}
	*a = strings.Split(str, ",")
	return nil
}

func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}
	return strings.Join(a, ","), nil
}
