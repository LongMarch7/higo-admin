package member

import (
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/util/token"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "github.com/LongMarch7/higo/util/define"
    "time"
)

type UserRole struct {
    models.MicroUser  `xorm:"extends"`
    RoleStatus int    `json:"role_status"`
    RoleName   string `json:"role_name"`
    RoleId     int    `json:"role_id"`
}

func GetSysMemberListWhereSql(where map[string]string) (string, []interface{}) {
    var sql = "user_type != ?"
    var value []interface{}
    value = append(value, define.AdminType)
    if v, ok := where["user_name"]; ok && v != "" {
        keywords := define.TrimString(v)
        sql += " AND user_login like ?"
        value = append(value,"%"+keywords+"%")
    }

    if v, ok := where["create_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND create_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["create_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND create_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["update_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND update_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["update_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND update_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["login_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND last_login_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["login_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND last_login_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["freeze_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND freeze_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["freeze_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND freeze_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["status_flag"]; ok && v != "" {
        status, _ :=strconv.Atoi(v)
        if status >0 && status <4{
            sql += " AND user_status = ?"
            value = append(value,status)
        }
    }
    return sql, value
}

func GetMemberList(where map[string]string, page int, rows int)  (user_list []models.MicroUser, count int){
    engine := db.NewDb(db.DefaultNAME).Engine()
    count = 0
    sql , value:= GetSysMemberListWhereSql(where)
    user_list = make([]models.MicroUser, 0)
    err := engine.Table(models.MicroUser{}.TableName()).Where(sql, value...).Limit(rows,(page-1)*rows).
        Find(&user_list)
    if err != nil {
        grpclog.Error("get member list err: ",err.Error())
        return
    }
    ret,err :=engine.Table(models.MicroUser{}.TableName()).Where(sql, value...).Count(&models.MicroUser{})
    if err != nil {
        grpclog.Error("get member count err: ",err.Error())
        return
    }
    count = int(ret)
    return
}

type AuthorizedRole struct{
    ID       int
    Name     string
    Status   int
    Remark   string
}

func GetMemberInfo(user_id string) models.MicroUser{
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
func Delete(users_id []string) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg ="执行出错"
    success = false
    for _,value := range users_id{
        info :=utils.GetUserInfo(value)
        if info.UserType == define.AdminType{
            msg = "权限不足"
            return
        }
        user :=models.MicroUser{}
        if _,err :=engine.Table(models.MicroUser{}.TableName()).ID(value).Delete(&user); err != nil{
            grpclog.Error("delete failed !",err.Error())
            msg = info.UserLogin + "删除用户失败"
            return
        }
        db.MemcacheClient.Delete(define.UserPrefix + info.UserLogin)
    }
    msg = "操作成功"
    success = true
    return
}

func StatusChange(users_id []string, users_status int) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    user :=models.MicroUser{}
    success = false
    msg = "操作出错"
    for _,value := range users_id{
        info :=utils.GetUserInfo(value)
        if info.UserType == define.AdminType{
            msg = "权限不足"
            return
        }
        user.UserStatus = users_status
        user.UpdateTime = int(time.Now().Unix())
        if _,err :=engine.Table(models.MicroUser{}.TableName()).ID(value).Cols("user_status").Update(&user); err != nil{
            grpclog.Error("update user failed ", err.Error())
            msg = info.UserLogin + "用户状态更新失败"
            return
        }
        db.MemcacheClient.Delete(define.UserPrefix + info.UserLogin)
    }
    msg = "操作成功"
    success = true
    return
}

func Edit(user_data utils.MemberPostData, is_super bool) (success bool,msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg = "执行错误"
    success = false

    nowTime := int(time.Now().Unix())
    user :=models.MicroUser{
        UserLogin:user_data.UserLogin,
        UpdateTime:int(time.Now().Unix()),
        UserEmail: user_data.UserEmail,
        Avatar: user_data.Avatar,
        BindId: user_data.BindId,
        Sex: user_data.Sex,
        Birthday: int(define.GetTimestamp(user_data.Birthday)),
        UserNickname: user_data.UserNickname,
        Signature: user_data.Signature,
    }
    if user_data.Id == 0 {
        user.UserType = user_data.UserType
    }
    if is_super && len(user_data.Balance) > 0 {
        user.Balance = user_data.Balance
    }
    if is_super && user_data.Coin > 0 {
        user.Coin = user_data.Coin
    }
    if is_super && user_data.Score > 0 {
        user.Score = user_data.Score
    }
    if len(user_data.UserPass) > 0 {
        user.UserPass = token.NewTokenWithSalt(user_data.UserPass)
    }
    if len(user_data.PayPass) > 0 {
        user.PayPass = token.NewTokenWithSalt(user_data.PayPass)
    }
    var user_id = user_data.Id
    if user_id >0{
        if  _,err := engine.Table(models.MicroUser{}.TableName()).ID(user_id).Update(&user); err != nil{
            grpclog.Error("User update:",err.Error())
            msg = "更新出错"
            return
        }
    }else{
        has,err :=engine.Table(models.MicroUser{}.TableName()).Where("user_login = ?", user_data.UserLogin).Exist(&models.MicroUser{})
        if err == nil && has{
            grpclog.Error("User name exits:")
            msg = "用户存在"
            return
        }
        user.CreateTime = nowTime
        user.UserStatus = 1
        user.Balance = "0.00"
        if  _,err = engine.Table(models.MicroUser{}.TableName()).Insert(&user); err != nil{
            grpclog.Error("User insert:",err.Error())
            return
        }
    }
    db.MemcacheClient.Delete(define.UserPrefix + user_data.UserLogin)
    success = true
    return
}