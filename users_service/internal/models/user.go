package models

import (
	"database/sql"
	"strings"
	"time"

	"github.com/X3ne/ds_ms/users_service/services"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `db:"id" gorm:"primaryKey;autoIncrement:false"`
	Username  string         `db:"username" gorm:"unique;not null"`
	Email     string         `db:"email" gorm:"unique;not null"`
	Password  sql.NullString `db:"password"`
	CreatedAt time.Time      `db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"autoUpdateTime"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	user.ID = node.Generate().String()
	user.Email = strings.ToLower(user.Email)

	if user.Password.Valid {
		hash, err := services.HashPassword(user.Password.String, &services.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			KeyLength:   32,
			SaltLength:  16,
		})
		if err != nil {
			return err
		}

		user.Password = sql.NullString{
			String: hash,
			Valid:  true,
		}
	}
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.Email = strings.ToLower(user.Email)

	if user.Password.Valid {
		hash, err := services.HashPassword(user.Password.String, &services.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			KeyLength:   32,
			SaltLength:  16,
		})
		if err != nil {
			return err
		}

		user.Password = sql.NullString{
			String: hash,
			Valid:  true,
		}
	}
	return
}
