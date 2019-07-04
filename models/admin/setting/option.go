package setting

import (
    "encoding/json"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo/db"
    "github.com/bradfitz/gomemcache/memcache"
    "google.golang.org/grpc/grpclog"
)

func GetOptionInfoFromCache(option_key string, info interface{}) (){
    if it, err :=  db.MemcacheClient.Get(option_key); err == nil && it.Key == option_key{
        err :=json.Unmarshal(it.Value,info)
        if err == nil {
            return
        }
    }
    GetInfo(option_key, info)
    return
}

func GetInfo(option_key string,info interface{}){
    engine := db.NewDb(db.DefaultNAME).Engine()
    var optionInfo = models.MicroOption{}
    has, err :=engine.Table(models.MicroOption{}.TableName()).Where("option_name = ?", option_key).Get(&optionInfo)
    if err != nil{
        grpclog.Error("Found info failed")
    }else if !has {
        grpclog.Error("Not found info")
    }
    db.MemcacheClient.Set(&memcache.Item{Key: option_key, Value: []byte(optionInfo.OptionValue)})
    json.Unmarshal([]byte(optionInfo.OptionValue),info)
    return
}

func SetInfo(option_key string, info interface{}) (success bool){
    success = false
    engine := db.NewDb(db.DefaultNAME).Engine()
    infoByte, err:=json.Marshal(info)
    defer func() {
        if success {
            db.MemcacheClient.Delete(option_key)
        }
    }()
    if err != nil {
        grpclog.Error("SetMailInfo err ",err.Error())
        return
    }
    mailOption := models.MicroOption{OptionName:option_key, OptionValue:string(infoByte)}
    has, err := engine.Table(models.MicroOption{}.TableName()).Where("option_name = ?", option_key).Exist(&models.MicroOption{})
    if err == nil && has{
        _,err := engine.Table(models.MicroOption{}.TableName()).Where("option_name = ?", option_key).Update(&mailOption)
        if err == nil{
            success = true
        }else{
            grpclog.Error("Update Info err ", err.Error())
        }
    }else{
        _,err := engine.Table(models.MicroOption{}.TableName()).Insert(mailOption)
        if err == nil{
            success = true
        }else{
            grpclog.Error("Insert Info err ", err.Error())
        }
    }
    return
}

