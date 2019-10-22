package template

import (
	orm "template/main/databases"

	"github.com/jinzhu/gorm"
)

type Template struct {
	gorm.Model
	Flex1 string `json:"flex1" form:"flex1" gorm:"size:20"`
	Flex2 string `json:"flex2" form:"flex2" gorm:"size:20"`
	Flex3 string `json:"flex3" form:"flex3" gorm:"size:20"`
	Flex4 string `json:"flex4" form:"flex4" gorm:"size:20"`
	Flex5 string `json:"flex5" form:"flex5" gorm:"size:20"`
}

func (entity *Template) GetAll() (results []Template, err error) {

	if err = orm.Eloquent.Find(&results).Error; err != nil {
		return
	}
	return
}

func (entity *Template) GetById() (result Template, err error) {

	if err = orm.Eloquent.First(&result, entity.ID).Error; err != nil {
		return
	}
	return
}

func (entity *Template) GetByCriteria() (results []Template, err error) {

	if err = orm.Eloquent.Where(&entity).Find(&results).Error; err != nil {
		return
	}

	return
}

func (entity *Template) Create() error {

	err := orm.Eloquent.Create(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (entity *Template) Update() error {

	err := orm.Eloquent.Model(&entity).Update(&entity).Error
	if err != nil {
		return err
	}

	return nil
}

func (entity *Template) Delete() error {

	err := orm.Eloquent.Delete(&entity).Error
	if err != nil {
		return err
	}

	return nil
}
