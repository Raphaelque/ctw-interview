package model

import (
	"errors"

	"gorm.io/gorm"
)

type Task struct {
	Id        int64          `json:"id"`
	CreatedAt int64          `gorm:"type:int;column:created_at;not null"`
	UpdatedAt int64          `gorm:"type:int;column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t *Task) Save() (*Task, error) {
	result := DB.Save(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func GetTaskById(id int64) (*Task, error) {
	if id == 0 {
		return nil, errors.New("id 为空！")
	}
	task := Task{Id: id}
	var err error = nil
	err = DB.First(&task, "id = ?", id).Error
	return &task, err
}
