package models

import (
	"database/sql"
	"time"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type Guild struct {
	ID          string         `db:"id" gorm:"primaryKey;autoIncrement:false"`
	Name        string         `db:"name" gorm:"unique;not null"`
	Icon        sql.NullString `db:"icon"`
	Splash      sql.NullString `db:"splash"`
	Banner      sql.NullString `db:"banner"`
	Description sql.NullString `db:"description"`
	OwnerID     string         `db:"owner_id" gorm:"not null"`
	//Roles       []        `db:"roles" gorm:"many2many:guild_roles;"`
	//Members   []GuildMember `gorm:"foreignKey:GuildID;references:ID"`
	CreatedAt time.Time `db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `db:"updated_at" gorm:"autoUpdateTime"`
}

func (guild *Guild) BeforeCreate(tx *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	guild.ID = node.Generate().String()

	return nil
}
