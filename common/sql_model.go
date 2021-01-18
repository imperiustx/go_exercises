package common

import "time"

type SQLModel struct {
	ID        int        `json:"-" gorm:"column:id;"`
	FakeID    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at;"`
}

func (m *SQLModel) GenUID(objectType int, shardId uint32) {
	uid := NewUID(uint32(m.ID), objectType, shardId)
	m.FakeID = &uid
}

type SQLModelCreate struct {
	ID     int  `json:"-" gorm:"column:id;"`
	FakeID *UID `json:"id" gorm:"-"`
	Status *int `json:"status" gorm:"column:status;default:1;"`
}

func (m *SQLModelCreate) GenUID(objectType int, shardId uint32) {
	uid := NewUID(uint32(m.ID), objectType, shardId)
	m.FakeID = &uid
}
