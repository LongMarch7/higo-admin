package admin

import (
    "encoding/json"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/auth"
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/util/token"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/bradfitz/gomemcache/memcache"
    "google.golang.org/grpc/grpclog"
    "strings"
    "time"
)


//func GetUserInfo(user_token string, user_name string) UserRole{
//   key := user_token + user_name
//   var userInfo = UserRole{}
//   if it, err := db.MemcacheClient.Get(user_token); err == nil && it.Key == user_token{
//       name := string(it.Value)
//       if strings.Compare(user_name,name) == 0 {
//           if it, err := db.MemcacheClient.Get(name); err == nil && it.Key == name{
//               json.Unmarshal(it.Value,&userInfo)
//           }
//           if len(userInfo.RoleName) >0 && global.RolePowerCheck(userInfo.RoleName){
//               return true
//           }
//       }
//   }
//   return userInfo
//}
//func GetUserInfo(user_name string) UserRole{
//    engine := db.NewDb(db.DefaultNAME).Engine()
//    var user = UserRole{}
//    ok, err := engine.Table(models.MicroUser{}.TableName()).Alias("u").Where("user_login = ?",user_name).
//        Join("INNER",[]string{models.MicroRoleUser{}.TableName(),"ur"},"u.id = ur.user_id").
//        Join("INNER",[]string{models.MicroRole{}.TableName(),"r"},"ur.role_id = r.id").
//        Get(&user)
//    if err != nil{
//        grpclog.Error("Exec sql failed")
//    }else if !ok {
//        grpclog.Error("Not found user")
//    }
//    return user
//}
func CheckRoleStatus(role_name string) bool{
    var stat = false
    key := define.RolePrefix + role_name
    if it, err :=  db.MemcacheClient.Get(key); err == nil && it.Key == key{
        define.FromByte(it.Value,&stat)
    }else {
        utils.UpdateRoleStatus()
        if it, err :=  db.MemcacheClient.Get(key); err == nil && it.Key == key{
            define.FromByte(it.Value,&stat)
        }
    }
    //grpclog.Info("role ",role_name ,"stat", stat)
    return stat
}
func IsLogin(user_token string, user_name string, pattern string, method string) (is_login bool,user_info []utils.LoginInfo){
    user_info = make([]utils.LoginInfo, 0)
    is_login = false
    defer func() {
       if is_login && strings.Compare(pattern,define.LoginPattern) != 0 {
           is_login = false
           casbin := auth.NewCasbin().Enforcer()
           for _,user := range user_info{
               if user.RoleId == define.SuperRoleId || casbin.Enforce(user.RoleName,pattern, method) {
                   is_login = true
                   break
               }
           }
       }
    }()
    if it, err := db.MemcacheClient.Get(user_token); err == nil && it.Key == user_token{
        name := string(it.Value)
        
        if strings.Compare(user_name,name) == 0 {
            key := define.UserPrefix + user_name
            if it, err := db.MemcacheClient.Get(key); err == nil && it.Key == key{
                json.Unmarshal(it.Value,&user_info)
                is_login = true
                return
            }else if success, _, info := AdminLogin(user_token, user_name, "", false); success{
                user_info = info
                is_login = true
                return
            }
        }
    }
    return
}

func AdminLogin(user_token string, user_name string, pw string, need_check_pw bool) (succecss bool, t string, user_info []utils.LoginInfo){
    engine := db.NewDb(db.DefaultNAME).Engine()
    succecss = false
    t = user_token
    sql := "user_login = ?"
    where := make([]interface{},0)
    where = append(where, user_name)
    if need_check_pw {
        sql += " AND user_pass = ?"
        where = append(where, token.NewTokenWithSalt2(pw))
        t = token.NewTokenWithTime(user_name)
    }
    user_info = make([]utils.LoginInfo, 0)
    err := engine.Table(models.MicroUser{}.TableName()).Alias("u").Where(sql,where...).
        Join("INNER",[]string{models.MicroRoleUser{}.TableName(),"ur"},"u.id = ur.user_id").
        Join("INNER",[]string{models.MicroRole{}.TableName(),"r"},"ur.role_id = r.id").
        Find(&user_info)
    if err != nil{
        grpclog.Error("AdminLogin Exec sql failed")
        return
    }
    if len(user_info) == 0 {
        grpclog.Error("Not found user")
        return
    }
    j :=0
    casbin := auth.NewCasbin().Enforcer()
    permission := false
    for _,value := range user_info {
        if err := validator.Validate.Struct(value); err != nil {
            grpclog.Error("User Info check err")
            return
        }
        if value.UserStatus != 1 {
            grpclog.Error("User is banned")
            return
        }else if len(value.RoleName)!=0 && CheckRoleStatus(value.RoleName){
            if casbin.Enforce(value.RoleName,"admin:LoginPost", "POST") || value.RoleId == define.SuperRoleId{
                permission = true
            }
            user_info[j] = value
            j++
        }
    }

    user_info = user_info[0:j]
    if !permission{
        grpclog.Error("Role is banned")
        return
    }
    db.MemcacheClient.Set(&memcache.Item{Key: t, Value: []byte(user_name)})
    if value,err :=json.Marshal(user_info); err ==nil{
        db.MemcacheClient.Set(&memcache.Item{Key: define.UserPrefix + user_name, Value: value})
    }else{
        grpclog.Info( err.Error() )
    }
    engine.Table(models.MicroUser{}.TableName()).Cols("last_login_time").Update(&models.MicroUser{UpdateTime:int(time.Now().Unix())})
    succecss = true
    return
}