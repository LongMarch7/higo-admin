package user

import (
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/auth"
    "github.com/LongMarch7/higo/db"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "github.com/LongMarch7/higo/util/define"
    "time"
)

func GetSysRoleListWhereSql(where map[string]string) (string, []interface{}) {
    var sql = "(1=1"
    var value []interface{}
    if v, ok := where["role_name"]; ok && v != "" {
        keywords := define.TrimString(v)
        sql += " AND role_name like ?"
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

    if v, ok := where["status_flag"]; ok && v != "" {
        status, _ :=strconv.Atoi(v)
        if status >0 && status <3{
            sql += " AND role_status = ?"
            value = append(value,status)
        }
    }
    sql += ")"
    return sql, value
}

func GetRoleList(where map[string]string, page int, rows int, user_info []utils.LoginInfo)  (list []models.MicroRole, count int){
    engine := db.NewDb(db.DefaultNAME).Engine()
    list = make([]models.MicroRole, 0)
    count = 0
    sql , value:= GetSysRoleListWhereSql(where)
    _, sql1 , value1:= GetRoleChilds(user_info,"id", false)
    notSuper := true
    for _, user := range user_info{
        if user.RoleId == define.SuperRoleId {
            notSuper = false
            break
        }
    }
    if notSuper {
        sql += "AND" + sql1
        value = append(value,value1...)
    } else{
        sql += "AND ( id != ? )"
        value = append(value, define.SuperRoleId)
    }

    err :=engine.Table(models.MicroRole{}.TableName()).Limit(rows,(page-1)*rows).Where(sql, value...).Limit(rows,(page-1)*rows).Find(&list)
    if err != nil {
        grpclog.Error("get role list err: ",err.Error())
        return
    }
    ret,err :=engine.Table(models.MicroRole{}.TableName()).Where(sql, value...).Count(&models.MicroRole{})
    if err != nil {
        grpclog.Error("get role count err: ",err.Error())
        return
    }
    count = int(ret)
    return
}

func GetRoleInfo(role_id string) models.MicroRole{
    engine := db.NewDb(db.DefaultNAME).Engine()
    roleInfo := models.MicroRole{}

    has, err :=engine.Table(models.MicroRole{}.TableName()).ID(role_id).Get(&roleInfo)
    if err != nil{
        grpclog.Error("Found role failed")
    }else if !has {
        grpclog.Error("Not found user")
    }
    return roleInfo
}
func DelRoles(roles_id []string, roles_name []string, user_info []utils.LoginInfo) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    session := engine.NewSession()
    defer session.Close()
    role :=models.MicroRole{}
    msg ="执行出错"
    success = false
    for index,value := range roles_id{
        if ok,message :=RolePowerCheck(user_info, value); !ok{
            success = ok
            msg = message
            return
        }
        err := session.Begin()
        has,err :=session.Table(models.MicroRole{}.TableName()).Where("parent_id = ?",value).Exist(&models.MicroRole{})
        if err == nil && has{
            grpclog.Error("delete failed ,because role has child !")
            msg = roles_name[index] + "角色删除失败，存在子角色依赖"
            return
        }
        has,err = session.Table(models.MicroRole{}.TableName()).Alias("r").Where("r.id = ?", value).
            Join("INNER",[]string{models.MicroRoleUser{}.TableName(),"ur"},"r.id = ur.role_id").
            Join("INNER",[]string{models.MicroUser{}.TableName(),"u"},"ur.user_id = u.id").Exist(&utils.LoginInfo{})
        if err == nil && has{
            grpclog.Error("delete failed ,because role has user !")
            msg = roles_name[index] + "角色删除失败，该角色存在用户"
            return
        }
        if _,err =session.Table(models.MicroRole{}.TableName()).ID(value).Delete(&role); err != nil{
            session.Rollback()
            grpclog.Error("delete failed !",err.Error())
            msg = roles_name[index] + "删除角色失败"
            return
        }
        if _,err = session.Table(models.MicroCasbinRule{}.TableName()).Where("v0 = ?",roles_name[index]).Delete(&models.MicroCasbinRule{}); err != nil{
            session.Rollback()
            grpclog.Error("Role delete policy err:",err.Error())
            msg = roles_name[index] + "角色规则删除失败"
            return
        }
        if err = session.Commit();err != nil{
            session.Rollback()
            grpclog.Error("delete  failed by session !",err.Error())
            return
        }
        db.MemcacheClient.Delete(define.RolePrefix + roles_name[index])
    }
    msg = "操作成功"
    success = true
    return
}

func RolesStatusChange(roles_id []string, roles_status int, user_info []utils.LoginInfo) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    role :=models.MicroRole{}
    success = false
    msg = "操作出错"
    for _,value := range roles_id{
        if ok,message :=RolePowerCheck(user_info, value); !ok{
            success = ok
            msg = message
            return
        }
        role.RoleStatus = roles_status
        role.UpdateTime = int(time.Now().Unix())
        if _,err :=engine.Table(models.MicroRole{}.TableName()).ID(value).Cols("role_status").Update(&role); err != nil{
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}

func RolePowerCheck(user_info []utils.LoginInfo, role_id string)  (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    success = false
    msg = "操作出错"
    parentRoles := make(map[int]bool)
    id,err := strconv.Atoi(role_id)
    if err != nil {
        return
    }
    for _,user := range user_info{
        if user.RoleId == define.SuperRoleId {
            if user.RoleId != id {
                success = true
            }
            msg = "super user"
            return
        }
        parentRoles[user.RoleId] = true
    }
    list := make([]models.MicroRole, 0)
    rolesMap := make(map[int]models.MicroRole)
    err =engine.Table(models.MicroRole{}.TableName()).Find(&list)
    if err != nil {
        grpclog.Error("get role count err: ",err.Error())
        return
    }
    for _,value := range list{
        if _,ok :=parentRoles[value.ParentId]; ok{
            rolesMap[value.Id] = value
        }
        if _,ok := rolesMap[value.ParentId]; ok{
            rolesMap[value.Id] = value
        }
    }
    if _,ok := rolesMap[id]; ok{
        success = true
        msg ="拥有权限"
    }
    return
}

func MenuPowerCheck(list []models.MicroAdminMenu, user_info []utils.LoginInfo, role_name string) []models.MicroAdminMenu{
    casbin := auth.NewCasbin().Enforcer()
    isSuper := false
    for _,user := range user_info{
        if user.RoleId == define.SuperRoleId {
            isSuper = true
            break
        }
    }
    j := 0
    for _,value :=range list{
        for _,user := range user_info {
            if casbin.Enforce(user.RoleName, value.Url+":"+value.Func, value.Method) || isSuper {
                value.Status = 0
                if len(role_name) > 0 && casbin.Enforce(role_name, value.Url+":"+value.Func, value.Method) {
                    value.Status = 1
                }
                list[j] = value
                j++
                break
            }
        }
    }
    return list[:j]
}

func makePolicy(list ...string) models.MicroCasbinRule{
    line := models.MicroCasbinRule{}

    if len(list) > 0 {
        line.PType = list[0]
    }
    if len(list) > 1 {
        line.V0 = list[1]
    }
    if len(list) > 2 {
        line.V1 = list[2]
    }
    if len(list) > 3 {
        line.V2 = list[3]
    }
    if len(list) > 4 {
        line.V3 = list[4]
    }
    if len(list) > 5 {
        line.V4 = list[5]
    }
    if len(list) > 6 {
        line.V5 = list[6]
    }
    return line
}
func policyWhere(line models.MicroCasbinRule) (string, []interface{}) {
    queryArgs := []interface{}{line.PType}
    queryStr := "p_type = ?"
    if line.V0 != "" {
        queryStr += " and v0 = ?"
        queryArgs = append(queryArgs, line.V0)
    }
    if line.V1 != "" {
        queryStr += " and v1 = ?"
        queryArgs = append(queryArgs, line.V1)
    }
    if line.V2 != "" {
        queryStr += " and v2 = ?"
        queryArgs = append(queryArgs, line.V2)
    }
    if line.V3 != "" {
        queryStr += " and v3 = ?"
        queryArgs = append(queryArgs, line.V3)
    }
    if line.V4 != "" {
        queryStr += " and v4 = ?"
        queryArgs = append(queryArgs, line.V4)
    }
    if line.V5 != "" {
        queryStr += " and v5 = ?"
        queryArgs = append(queryArgs, line.V5)
    }
    return queryStr,queryArgs
}
func RoleEdit(list []utils.MenuPower, role_remark string, role_id string, role_parent_id int, role_name string, user_info []utils.LoginInfo) (success bool,msg string){
    casbin := auth.NewCasbin().Enforcer()
    engine := db.NewDb(db.DefaultNAME).Engine()
    session := engine.NewSession()
    defer session.Close()
    err := session.Begin()
    msg = "执行错误"
    success = false

    nowTime := int(time.Now().Unix())
    role :=models.MicroRole{UpdateTime: nowTime, Remark: role_remark, RoleName: role_name}
    if len(role_id) >0{
        if ok,message :=RolePowerCheck(user_info, role_id); !ok{
            success = ok
            msg = message
            return
        }
        if  _,err = session.Table(models.MicroRole{}.TableName()).ID(role_id).Cols("update_time","remark","role_name").Update(&role); err != nil{
            session.Rollback()
            grpclog.Error("Role update:",err.Error())
            msg = "更新出错"
            return
        }
    }else{
        has,err :=session.Table(models.MicroRole{}.TableName()).Where("role_name = ?", role_name).Exist(&models.MicroRole{})
        if err == nil && has{
            grpclog.Error("Role name exits:")
            msg = "角色名存在"
            return
        }
        role.CreateTime = nowTime
        role.ParentId = role_parent_id
        role.RoleStatus = 1
        if  _,err = session.Table(models.MicroRole{}.TableName()).Insert(&role); err != nil{
            session.Rollback()
            grpclog.Error("Role insert:",err.Error())
            return
        }
    }
    isSuper := false
    for _,user := range user_info{
        if user.RoleId == define.SuperRoleId {
            isSuper = true
            break
        }
    }
    for _,value := range list{
        policy :=makePolicy("p",role_name,value.Pattern, value.Method)
        sql,where :=policyWhere(policy)
        for _,user := range user_info {
            if casbin.Enforce(user.RoleName, value.Pattern, value.Method) || isSuper{
                if value.Status {
                    has, err := session.Table(models.MicroCasbinRule{}.TableName()).Where(sql, where...).Exist(&models.MicroCasbinRule{})
                    if err == nil && has {
                        break
                    }
                    if _, err = session.Table(models.MicroCasbinRule{}.TableName()).Insert(&policy); err != nil {
                        session.Rollback()
                        grpclog.Error("Role insert policy err:", err.Error())
                        return
                    }
                } else {
                    if _, err = session.Table(models.MicroCasbinRule{}.TableName()).Where(sql, where...).Delete(&policy); err != nil {
                        session.Rollback()
                        grpclog.Error("Role delete policy err:", err.Error())
                        return
                    }
                }
                break
            }
        }
    }

    if err = session.Commit();err != nil{
        session.Rollback()
        return
    }
    success = true
    return
}

func GetRoleChilds(user_info []utils.LoginInfo, field string, is_contain_self bool) (roles map[int]models.MicroRole, sql string, where []interface{}){
    sql = "(1=2"
    defer func() {
        sql +=")"
    }()
    field = " OR " + field + " = ?"
    engine := db.NewDb(db.DefaultNAME).Engine()
    list := make([]models.MicroRole, 0)
    roles = make(map[int]models.MicroRole)
    err :=engine.Table(models.MicroRole{}.TableName()).OrderBy("id").Find(&list)
    if err != nil {
        grpclog.Error("get role count err: ",err.Error())
        return
    }
    for _,user := range user_info{
        roles[user.RoleId] = models.MicroRole{}
    }
    for _,role := range list{
        if _,ok := roles[role.ParentId]; ok{
            roles[role.Id] = role
        }
    }
    for key,value := range roles{
        if is_contain_self || (!is_contain_self && value.Id == key) {
            sql += field
            where = append(where,key)
        }
    }
    grpclog.Info(list)
    grpclog.Info(roles)
    return
}


func GetAuthorizedRoleList(user_info []utils.LoginInfo)(roles map[int]AuthorizedRole){
    engine := db.NewDb(db.DefaultNAME).Engine()
    listRole := make([]models.MicroRole, 0)
    err :=engine.Table(models.MicroRole{}.TableName()).Find(&listRole)
    if err != nil {
        grpclog.Error("Load user list err: ",err.Error())
        return
    }
    roles = make(map[int]AuthorizedRole)
    for _,user := range user_info{
        roles[user.RoleId] = AuthorizedRole{ID:user.RoleId,Name:user.RoleName}
    }
    for _,value :=range listRole{
        if _,ok := roles[value.Id]; ok{
            roles[value.Id] = AuthorizedRole{ID:value.Id,Name:value.RoleName,Remark:value.Remark}
        }
        if _,ok := roles[value.ParentId]; ok{
            roles[value.Id] = AuthorizedRole{ID:value.Id,Name:value.RoleName,Remark:value.Remark}
        }
    }
    //if _,ok := roles[define.SuperRoleId]; ok{
    //    delete(roles,define.SuperRoleId)
    //}
    return
}