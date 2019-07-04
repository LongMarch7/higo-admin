package user

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

func GetSysUserListWhereSql(where map[string]string, user_name string) (string, []interface{}) {
    var sql = "u.user_login != ? AND u.user_type = ?"
    var value []interface{}
    value = append(value, user_name, define.AdminType)
    if v, ok := where["user_name"]; ok && v != "" {
        keywords := define.TrimString(v)
        sql += " AND u.user_login like ?"
        value = append(value,"%"+keywords+"%")
    }

    if v, ok := where["create_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND u.create_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["create_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND u.create_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["update_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND u.update_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["update_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND u.update_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["login_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND u.last_login_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["login_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND u.last_login_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["status_flag"]; ok && v != "" {
        status, _ :=strconv.Atoi(v)
        if status >0 && status <3{
            sql += " AND u.user_status = ?"
            value = append(value,status)
        }
    }
    return sql, value
}

func GetUserList(where map[string]string, page int, rows int, user_info []utils.LoginInfo, user_name string)  (user_role_list []models.MicroUser, count int){
    engine := db.NewDb(db.DefaultNAME).Engine()
    count = 0
    sql , value:= GetSysUserListWhereSql(where, user_name)
    isSuper  := false
    for _,user :=range user_info {
        if user.RoleId == define.SuperRoleId {
            isSuper = true
            break
        }
    }
    user_role_list = make([]models.MicroUser, 0)
    if isSuper {
        err := engine.Table(models.MicroUser{}.TableName()).Alias("u").Where(sql, value...).Where("user_type = ?", define.AdminType).Limit(rows,(page-1)*rows).Find(&user_role_list)
        if err != nil {
            grpclog.Error("get user list err: ",err.Error())
            return
        }
        ret,err := engine.Table(models.MicroUser{}.TableName()).Alias("u").Where(sql, value...).Where("user_type = ?", define.AdminType).Limit(rows,(page-1)*rows).Count(&models.MicroUser{})
        if err != nil {
            grpclog.Error("get user count err: ",err.Error())
            return
        }
        count = int(ret)
    }else{
        _, sql1,value1 :=GetRoleChilds(user_info,"ur.role_id", true)
        err := engine.Table(models.MicroUser{}.TableName()).Alias("u").Select("*, count(distinct u.id)").Where(sql, value...).
            Join("INNER",[]string{models.MicroRoleUser{}.TableName(),"ur"},"u.id = ur.user_id").Where(sql1, value1...).
            Join("INNER",[]string{models.MicroRole{}.TableName(),"r"},"ur.role_id = r.id").GroupBy("u.id").Limit(rows,(page-1)*rows).
            Find(&user_role_list)
        if err != nil {
            grpclog.Error("get user list err: ",err.Error())
            return
        }
        ret,err :=engine.Table(models.MicroUser{}.TableName()).Alias("u").Where(sql, value...).
            Join("INNER",[]string{models.MicroRoleUser{}.TableName(),"ur"},"u.id = ur.user_id").Where(sql1, value1...).
            Join("INNER",[]string{models.MicroRole{}.TableName(),"r"},"ur.role_id = r.id").Distinct("u.id").Count(&models.MicroUser{})
        if err != nil {
            grpclog.Error("get user count err: ",err.Error())
            return
        }
        count = int(ret)
    }

    return
}

type AuthorizedRole struct{
    ID       int
    Name     string
    Status   int
    Remark   string
}
func HasPower(user_id string, roles map[int]AuthorizedRole) (ret map[int]AuthorizedRole){
    engine := db.NewDb(db.DefaultNAME).Engine()
    roleUsers := make([]models.MicroRoleUser, 0)
    ret = roles
    err :=engine.Table(models.MicroRoleUser{}.TableName()).Where("user_id = ?", user_id).Find(&roleUsers)
    if err != nil{
        return
    }
    for _, roleUser := range roleUsers{
        if value,ok := ret[roleUser.RoleId]; ok{
            ret[roleUser.RoleId] = AuthorizedRole{ID:value.ID, Name:value.Name, Status:1, Remark:value.Remark}
        }
    }
    return
}

func DelUsers(users_id []string, users_name []string, user_info []utils.LoginInfo) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    session := engine.NewSession()
    defer session.Close()
    msg ="执行出错"
    success = false
    for index,value := range users_id{
        if ok,message :=UserPowerCheck(user_info, value); !ok{
            msg = message
            return
        }
        err := session.Begin()
        user :=models.MicroUser{}
        if _,err =session.Table(models.MicroUser{}.TableName()).ID(value).Delete(&user); err != nil{
            session.Rollback()
            grpclog.Error("delete failed !",err.Error())
            msg = users_name[index] + "删除用户失败"
            return
        }
        roleUser :=models.MicroRoleUser{}
        if _,err =session.Table(models.MicroRoleUser{}.TableName()).Where("user_id = ?", value).Delete(&roleUser); err != nil{
            session.Rollback()
            grpclog.Error("delete failed !",err.Error())
            msg = users_name[index] + "删除用户失败"
            return
        }
        if err = session.Commit();err != nil{
            session.Rollback()
            grpclog.Error("delete  failed by session !",err.Error())
            return
        }
        db.MemcacheClient.Delete(define.UserPrefix + users_name[index])
    }
    msg = "操作成功"
    success = true
    return
}

func UsersStatusChange(users_id []string, users_status int, user_info []utils.LoginInfo) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    user :=models.MicroUser{}
    success = false
    msg = "操作出错"
    for _,value := range users_id{
        if ok,message :=UserPowerCheck(user_info, value); !ok{
            msg = message
            return
        }
        user.UserStatus = users_status
        user.UpdateTime = int(time.Now().Unix())
        if _,err :=engine.Table(models.MicroUser{}.TableName()).ID(value).Cols("user_status").Update(&user); err != nil{
            grpclog.Error("update user failed ", err.Error())
            return
        }
        info :=utils.GetUserInfo(value)
        db.MemcacheClient.Delete(define.UserPrefix + info.UserLogin)
    }
    msg = "操作成功"
    success = true
    return
}


func UserPowerCheck(user_info []utils.LoginInfo, user_id string)  (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    success = false
    msg = "操作出错"
    var userId int64
    for _,user := range user_info{
        userId = user.UserId
        if user.RoleId == define.SuperRoleId {
            success = true
            msg = "super user"
            return
        }
    }
    id, err := strconv.ParseInt(user_id, 10, 64)
    defer func() {
        if  userId == id{
            success = false
            msg = "无法编辑自己"
            return
        }
    }()
    if err != nil {
        return
    }
    var user = make([]UserRole, 0)
    err = engine.Table(models.MicroUser{}.TableName()).Alias("u").Where("u.id = ?", user_id).
        Join("INNER",[]string{models.MicroRoleUser{}.TableName(),"ur"},"u.id = ur.user_id").
        Join("INNER",[]string{models.MicroRole{}.TableName(),"r"},"ur.role_id = r.id").
        Find(&user)
    if err != nil{
        msg = "查询失败"
        grpclog.Error("UserPowerCheck Exec sql failed",err.Error())
        return
    }
    roleMap, _,_ :=GetRoleChilds(user_info,"ur.role_id", true)
    for _,value := range user{
        if _,ok :=roleMap[value.RoleId]; ok{
            success = true
            msg = "操作成功"
            break
        }
    }
    return

}

func UserEdit(user_data utils.UserPostData, roles map[int]AuthorizedRole, role_no_check []string, user_info []utils.LoginInfo) (success bool,msg string){
    //casbin := auth.NewCasbin().Enforcer()
    engine := db.NewDb(db.DefaultNAME).Engine()
    session := engine.NewSession()
    defer session.Close()
    err := session.Begin()
    msg = "执行错误"
    success = false

    nowTime := int(time.Now().Unix())
    user :=models.MicroUser{
        UserLogin: user_data.UserLogin,
        UpdateTime: nowTime,
        UserEmail: user_data.UserEmail,
        Balance: "0.00",
    }
    if len(user_data.UserPass) > 0 {
        user.UserPass = token.NewTokenWithSalt(user_data.UserPass)
    }
    if len(user_data.PayPass) > 0 {
        user.PayPass = token.NewTokenWithSalt(user_data.PayPass)
    }
    var user_id = user_data.Id
    if user_id >0{
        if ok,message :=UserPowerCheck(user_info, strconv.FormatInt(user_id,10)); !ok{
            msg = message
            return
        }
        if  _,err = session.Table(models.MicroUser{}.TableName()).ID(user_id).Update(&user); err != nil{
            session.Rollback()
            grpclog.Error("User update:",err.Error())
            msg = "更新出错"
            return
        }
    }else{
        has,err :=session.Table(models.MicroUser{}.TableName()).Where("user_login = ?", user_data.UserLogin).Exist(&models.MicroUser{})
        if err == nil && has{
            session.Rollback()
            grpclog.Error("User name exits:")
            msg = "用户存在"
            return
        }
        user.CreateTime = nowTime
        user.UserStatus = 1
        user.UserType = 1
        if  _,err = session.Table(models.MicroUser{}.TableName()).Insert(&user); err != nil{
            session.Rollback()
            grpclog.Error("User insert:",err.Error())
            return
        }
        user_id = user.Id
    }
    for _,id := range user_data.RoleId{
        if _, ok := roles[id]; !ok{
            session.Rollback()
            grpclog.Error("User role_check role not found",err.Error())
            return
        }
        has,err :=session.Table(models.MicroRoleUser{}.TableName()).Where("role_id = ? AND user_id = ?", id, user_id).Exist(&models.MicroRoleUser{})
        if err == nil && has{
            continue
        }else{
            if  _,err = session.Table(models.MicroRoleUser{}.TableName()).Insert(&models.MicroRoleUser{RoleId:id, UserId: user_id}); err != nil{
                session.Rollback()
                grpclog.Error("RoleUser insert:",err.Error())
                return
            }
        }
    }
    for _, value := range role_no_check {
        id, err := strconv.Atoi(value)
        if err != nil {
            session.Rollback()
            grpclog.Error("User role_no_check atoi err:",err.Error())
            return
        }
        if _, ok := roles[id]; !ok{
            session.Rollback()
            grpclog.Error("User role_no_check role not found",err.Error())
            return
        }
        _, err =session.Table(models.MicroRoleUser{}.TableName()).Where("role_id = ? AND user_id = ?",id, user_id).Delete(&models.MicroRoleUser{})
        if err != nil {
            session.Rollback()
            grpclog.Error("User RoleUser delete error",err.Error())
            return
        }
    }

    if err = session.Commit();err != nil{
        session.Rollback()
        return
    }
    db.MemcacheClient.Delete(define.UserPrefix + user_data.UserLogin)
    success = true
    return
}