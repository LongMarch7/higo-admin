[build]
go get -v github.com/LongMarch7/higo-admin
[cross compile]
GOOS=linux GOARCH=amd64 go get -v github.com/LongMarch7/higo-admin


[run]
higo-admin -mode="init" -conf="E:/go_project/higo/src/github.com/LongMarch7/higo-admin/config/config.json"

higo-admin -mode="cli" -name="gateway" -conf="E:/go_project/higo/src/github.com/LongMarch7/higo-admin/config/config.json"
higo-admin -mode="svr" -name="AdminServer" -conf="E://go_project/higo/src/github.com/LongMarch7/higo-admin/config/config.json" -port="10085" -ad_port="10086"