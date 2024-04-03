package models

import "database/sql"

type GuildBan struct {
	UserID  string         `db:"id" gorm:"primaryKey;autoIncrement:false"`
	GuildID string         `db:"guild_id" gorm:"not null"`
	Reason  sql.NullString `db:"reason;default:null"`
}
