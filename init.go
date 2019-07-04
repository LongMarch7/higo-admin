package main

import (
    "fmt"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo/auth"
    "github.com/LongMarch7/higo/config"
    "github.com/LongMarch7/higo/db"
    _ "github.com/go-sql-driver/mysql"
    "os"
    "io"
)

func Initialization(config *config.Configer) {
    sql :=sqlResolving(config)
    args := sql.User + ":" + sql.Pwd + "@" + sql.Net + "(" + sql.Addr + ":" + sql.Port + ")/"
    engine := db.NewDb(db.DefaultNAME, db.Dialect(sql.Driver),
       db.Args(args),
       db.MaxOpenConns(10),
       db.MaxIdleConns(10),
       db.Show(sql.Show),
       db.Level(sql.Level)).Engine()

    if _, err := engine.Exec("CREATE DATABASE " + sql.Db + " DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"); err != nil {
       panic(err.Error())
    }
    db.DeleteDB(db.DefaultNAME)
    engine = db.NewDb(db.DefaultNAME, db.Dialect(sql.Driver),
      db.Args(args + sql.Db),
      db.MaxOpenConns(10),
      db.MaxIdleConns(10),
       db.Show(sql.Show),
       db.Level(sql.Level)).Engine()
    defer db.DeleteDB(db.DefaultNAME)
    err := engine.Ping()
    if err != nil {
        fmt.Println(err)
        return
    }

    var r io.Reader
    r, err = os.Open(sql.File)
    _, err = engine.Import(r)
    if err != nil{
       panic(err.Error())
    }
    var menuList = make([]models.MicroAdminMenu, 0)
    engine.Table(models.MicroAdminMenu{}.TableName()).Find(&menuList)
    casbin := auth.NewCasbin()
    for _,value := range menuList{
        if len(value.Url) == 0 ||  len(value.Func) == 0 || len(value.Method) == 0 {
            continue
        }
        casbin.Enforcer().AddPolicy([]string{"admin",value.Url + ":" + value.Func, value.Method})
    }
    casbin.Enforcer().AddPolicy([]string{"admin", "admin:LoginPost", "POST"})
    casbin.Enforcer().AddPolicy([]string{"admin", "admin:Index", "GET"})
    casbin.Enforcer().AddPolicy([]string{"admin", "admin:Welcome", "GET"})
    casbin.Enforcer().AddPolicy([]string{"admin", "admin/user:Update", "GET"})
    casbin.Enforcer().AddPolicy([]string{"admin", "admin:Upload", "POST"})
    //var roleList = make([]models.MicroRole, 0)
    //engine.Table(models.MicroRole{}.TableName()).Find(&roleList)
    //for _,value := range roleList{
    //    if value.RoleName == "super"{
    //        continue
    //    }
    //    casbin.Enforcer().AddRoleForUser("admin",value.RoleName)
    //}
}
