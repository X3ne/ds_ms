package models

import (
	"database/sql"
	"time"
)

type GuildMember struct {
	UserID  string         `db:"id" gorm:"primaryKey;autoIncrement:false"`
	GuildID string         `db:"guild_id" gorm:"not null"`
	Nick    sql.NullString `db:"nick"`
	Avatar  sql.NullString `db:"avatar"`
	//Roles                      pq.StringArray `db:"roles"`
	Deaf                       bool      `db:"deaf"`
	Mute                       bool      `db:"mute"`
	JoinedAt                   time.Time `db:"joined_at"`
	CommunicationDisabledUntil time.Time `db:"communication_disabled_until"`
}
