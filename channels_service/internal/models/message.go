package models

import (
	"database/sql"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	embedsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/embeds/v1"
	"github.com/bwmarrin/snowflake"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type EmbedArray []embedsv1.Embed

type Message struct {
	ID              string                 `db:"id" gorm:"primaryKey;autoIncrement:false"`
	ChannelID       string                 `db:"channel_id" gorm:"not null"`
	AuthorID        string                 `db:"author_id" gorm:"not null"`
	Content         sql.NullString         `db:"content"`
	MentionEveryone bool                   `db:"mention_everyone" gorm:"default:false"`
	Mentions        StringArray            `db:"mentions" gorm:"default:[];type:VARCHAR(255)"`
	MentionRoles    StringArray            `db:"mention_roles" gorm:"default:[];type:VARCHAR(255)"`
	MentionChannels StringArray            `db:"mention_channels" gorm:"default:[];type:VARCHAR(255)"`
	Attachments     StringArray            `db:"attachments" gorm:"default:[];type:VARCHAR(255)"`
	Embeds          datatypes.JSON         `db:"embeds" gorm:"default:[];type:VARCHAR(255)"`
	Reactions       StringArray            `db:"reactions" gorm:"default:[];type:VARCHAR(255)"`
	Nonce           int64                  `db:"nonce"`
	Pinned          bool                   `db:"pinned" gorm:"default:false"`
	Type            channelsv1.MessageType `db:"type" gorm:"not null"`
	CreatedAt       time.Time              `db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       sql.NullTime           `db:"updated_at" gorm:"autoUpdateTime"`
}

func (message *Message) BeforeCreate(tx *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	message.ID = node.Generate().String()

	return nil
}
