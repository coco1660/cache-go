package entity

import (
	"time"
)

type Tables struct {
	Id   int64  `xorm:"pk autoincr BIGINT 'id'" json:"id"`
	Name string `xorm:"not null unique VARCHAR(64) 'name'" json:"name"`
}

type CacheItems struct {
	Id       int64     `xorm:"pk autoincr BIGINT 'id'" json:"id"`
	TableID  int64     `xorm:"not null index BIGINT 'table_id'" json:"table_id"`
	Key      string    `xorm:"not null VARCHAR(128) 'key'" json:"key"`
	Value    string    `xorm:"not null TEXT 'value'" json:"value"`
	ExpireAt time.Time `xorm:"DATETIME 'expire_at'" json:"expire_at"`

	CreateTime time.Time `xorm:"created 'created_at'" json:"created_at"`
	UpdateTime time.Time `xorm:"updated 'updated_at'" json:"updated_at"`

	AccessCount int64 `xorm:"BIGINT default 0 'access_count'" json:"access_count"`
}
