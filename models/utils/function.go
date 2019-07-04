package utils

import (
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/bradfitz/gomemcache/memcache"
    "google.golang.org/grpc/grpclog"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "net/url"
)

func UpdateRoleStatus(){
    var roleStatus = make(map[int]bool)
    engine :=db.NewDb(db.DefaultNAME).Engine()
    list := make([]models.MicroRole, 0)
    err :=engine.Table(models.MicroRole{}.TableName()).Find(&list)
    if err != nil {
        grpclog.Error("Load role list err: ",err.Error())
        return
    }
    for _,value := range list{
        if value.RoleStatus == 1{
            roleStatus[value.Id] = true
        }else{
            roleStatus[value.Id] = false
        }
    }

    for _,value := range list{
        if value.ParentId >0 {
            if stat,ok :=roleStatus[value.ParentId]; ok && stat == false{
                roleStatus[value.Id] = false
            }
        }
        buf :=define.ToByte(roleStatus[value.Id])
        db.MemcacheClient.Set(&memcache.Item{Key: define.RolePrefix + value.RoleName, Value: buf})
    }
}

func GetUserInfo(user_id string) models.MicroUser{
    engine := db.NewDb(db.DefaultNAME).Engine()
    userInfo := models.MicroUser{}

    has, err :=engine.Table(models.MicroUser{}.TableName()).ID(user_id).Get(&userInfo)
    if err != nil{
        grpclog.Error("Found user failed")
    }else if !has {
        grpclog.Error("Not found user")
    }
    return userInfo
}


func FormDataCheck(values url.Values, data interface{}) (success bool){
    success = false
    err := validator.Decoder.Decode(data, values)
    if err != nil{
        grpclog.Error("form decode error: ",err.Error())
        return
    }
    if err := validator.Validate.Struct(data); err != nil {
        grpclog.Error("validator Struct error: ",err.Error())
        return
    }
    success = true
    return
}

