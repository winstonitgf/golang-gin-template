package services

import (
	"template/main/models/template"
)

func GetTemplate(entity *template.Template) (result []template.Template, err error) {
	result, err = entity.GetByCriteria()
	return
}

func PostTemplate(entity *template.Template) (err error) {
	err = entity.Create()
	return
}

func PutTemplate(entity *template.Template) (err error) {
	err = entity.Update()
	return
}

func DeleteTemplate(entity *template.Template) (err error) {
	err = entity.Delete()
	return
}
