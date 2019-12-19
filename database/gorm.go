package database

import (
	"bao-bet365-api/model/env"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Eloquent *gorm.DB

func Open() {
	var err error

	// 取得config參數
	host := env.Enviroment.Database.Host
	port := env.Enviroment.Database.Port
	user := env.Enviroment.Database.User
	password := env.Enviroment.Database.Password
	name := env.Enviroment.Database.Name
	parameter := env.Enviroment.Database.Parameter

	// 連線
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, name, parameter)
	fmt.Println("conn: ", conn)
	Eloquent, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	if Eloquent.Error != nil {
		panic(Eloquent.Error)
	}

	Eloquent.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	Eloquent.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	Eloquent.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// debug mode開啟log
	Eloquent.LogMode(gin.DebugMode == gin.Mode())

	// table name用單數
	Eloquent.SingularTable(true)
}

func Close() {
	Eloquent.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedAt"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedAt", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedOnField := scope.FieldByName("DeletedAt")
		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
