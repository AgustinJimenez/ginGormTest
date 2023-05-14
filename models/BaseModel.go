package models

import (
	"go_practice/initializers"

	"gorm.io/gorm"
)



type BaseModel struct {
	gorm.Model
}


func (b *BaseModel) First(self interface{}, where ...interface{}) *gorm.DB {
    return initializers.DB.First(self, where...)
}

func (b *BaseModel) Find(self interface{}, conds ...interface{}) *gorm.DB {
    return initializers.DB.Find(self, conds...)
}

func (b *BaseModel) Create() *gorm.DB {
    return initializers.DB.Create(b)
}
