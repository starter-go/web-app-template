package dao

import (
	"github.com/starter-go/web-app-template/app/data/dxo"
	"github.com/starter-go/web-app-template/app/data/entity"
	"gorm.io/gorm"
)

// ExampleDAO ...
type ExampleDAO interface {
	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)
}
