# golang-gin-template

安裝Hot Reload套件(Optional)
> go get github.com/codegangsta/gin

安裝Swagger產生器(Optional)
> brew tap go-swagger/go-swagger

> brew install go-swagger

> 輸入命令：make generate-swagger產生json文檔

Migration
> 在migrattion目錄中定義好要migrate的model

> 輸入命令：make migration

1. 初始化專案
> go mod init {模組名稱}
2. 將專案中的import區塊，*template/main*，取代成你的模組名稱
3. 輸入命令*make start*啟動
