package user

import (
    "bytes"
    "context"
    "github.com/LongMarch7/higo-admin/models/admin"
    "github.com/LongMarch7/higo-admin/models/admin/user"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/auth"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/LongMarch7/higo/view"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "strings"
)

func (u* adminUserController)RoleIndex(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            view.NewView().Render(out, name + "/role_index", nil)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (u* adminUserController)RoleList(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            where := make(map[string]string)
            where["role_name"] 	= param.GetParams.Get("role_name")
            timeFlag := param.GetParams.Get("time_flag")
            startTime := param.GetParams.Get("start_time")
            endTime := param.GetParams.Get("end_time")
            if strings.Compare(timeFlag,"create") == 0 {
                where["create_start_time"] = startTime
                where["create_end_time"]   = endTime
            }else{
                where["update_start_time"] = startTime
                where["update_end_time"]   = endTime
            }
            where["status_flag"] = param.GetParams.Get("status_flag")
            page_num := param.GetParams.Get("page")
            page,err :=strconv.Atoi(page_num)
            if err != nil || page <=0{
                page = 1
            }
            row_num := param.GetParams.Get("limit")
            row,err :=strconv.Atoi(row_num)
            if err != nil || row <=0{
                row = 10
            }
            roleList, count:= user.GetRoleList(where, page, row, info)
            return base.NewLayuiRet(0, "获取成功",count, roleList), nil
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
func (u* adminUserController)RoleDelete(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id   := param.PostFormParams["id"]
            name := param.PostFormParams["name"]
            if err :=validator.Validate.Var(&name, validator.ArrayAlphanum); err != nil {
                grpclog.Error("role name 不合法",err.Error())
                return base.NewLayuiRet(-5, "参数错误-5",0,nil), nil
            }
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("role id 不合法",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-3",0,nil), nil
            }
            if succsess,msg :=user.DelRoles(id, name, info); succsess{
                auth.NewCasbin().ReloadPolicy()
                utils.UpdateRoleStatus()
                return base.NewLayuiRet(0, "删除成功", 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg, 0, nil), nil
            }
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (u* adminUserController)RoleStatusChange(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("RoleStatusChange id validate err：", err.Error())
                return base.NewLayuiRet(-6, "参数不合法-6",0,nil), nil
            }
            status := param.PostFormParams.Get("status")
            sta,err :=strconv.Atoi(status)
            if err != nil{
                grpclog.Error("RoleStatusChange status err：", err.Error())
                return base.NewLayuiRet(-5, "参数不合法-5",0,nil), nil
            }else if (sta <1) || (sta >2){
                grpclog.Error("RoleStatusChange status2 err")
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if success,msg :=user.RolesStatusChange(id, sta, info); success{
                var message = "角色:"
                for _,value := range id{
                    message += "[" + value + "]"
                }
                if sta == 1 {
                    message += "已启用"
                }else{
                    message += "已禁用"
                }
                utils.UpdateRoleStatus()
                return base.NewLayuiRet(0, message, 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }

        }
    }
    grpclog.Error("RoleStatusChange parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (u* adminUserController)RoleEdit(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            roleRemark := ""
            roleName := ""
            roleParentName := ""
            roleParentId := param.GetParams.Get("role_parent")
            roleId := param.GetParams.Get("role_id")
            rolesMap := user.GetAuthorizedRoleList(info)
            if len(roleParentId) > 0 {
                pId, err := strconv.Atoi(roleParentId)
                if err != nil {
                    data["content"] = "参数不合法"
                    view.NewView().Render(out,"error", data)
                    return out.String(), nil
                }
                role, ok := rolesMap[pId]
                if !ok {
                    data["content"] = "父权限不足"
                    view.NewView().Render(out,"error", data)
                    return out.String(), nil
                }
                role.Status = 1
                rolesMap[pId] = role
                roleParentInfo := user.GetRoleInfo(roleParentId)
                roleParentName = roleParentInfo.RoleName
                info = []utils.LoginInfo{{RoleId:roleParentInfo.Id,RoleName:roleParentInfo.RoleName}}
                if len(roleId) > 0 {
                    if success,msg :=user.RolePowerCheck(info, roleId); !success{
                        data["content"] = msg
                        view.NewView().Render(out,"error", data)
                        return out.String(), nil
                    }
                    roleInfo := user.GetRoleInfo(roleId)
                    if roleInfo.ParentId != pId{
                        data["content"] = "权限不足"
                        view.NewView().Render(out,"error", data)
                        return out.String(), nil
                    }
                    roleRemark = roleInfo.Remark
                    roleName = roleInfo.RoleName
                }
            }

            data["CurrUserMenu"] = user.MenuPowerCheck(admin.GetMenuList(info), info, roleName)
            data["RoleName"] = roleName
            data["RoleParentName"] = roleParentName
            data["RolesMap"] = rolesMap
            data["ParentId"] = roleParentId
            data["Intro"] = roleRemark
            data["Id"] = roleId
            view.NewView().Render(out, name+"/role_edit", data)
            return out.String(), nil
        }
    }
    data["content"] = "操作异常"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (u* adminUserController)RoleEditPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            noCheck :=param.PostFormParams.Get("NoCheck")
            var noCheckSplit []string
            if len(noCheck) > 0 {
                noCheckSplit = strings.Split(noCheck,",")
            }
            list,success := admin.GetMenuPower(param.PostFormParams["Power"], noCheckSplit)
            if !success {
                grpclog.Error("参数错误-5")
                return base.NewLayuiRet(-5, "参数错误-5",0,nil), nil
            }
            roleName := param.PostFormParams.Get("RoleName")
            if err :=validator.Validate.Var(roleName, validator.Alphanum); err != nil {
                grpclog.Error("RoleName validator check error:",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            remark :=param.PostFormParams.Get("Intro")
            id  :=param.PostFormParams.Get("Id")
            parentId  :=param.PostFormParams.Get("ParentId")
            rolesMap := user.GetAuthorizedRoleList(info)
            pId,err :=strconv.Atoi(parentId)
            if err !=nil {
                grpclog.Error("Role parentId invalid",err.Error())
                return base.NewLayuiRet(-3, "参数不合法-3",0,nil), nil
            }
            if _, ok := rolesMap[pId]; !ok {
                grpclog.Error("Role parentId invalid1",err.Error())
                return base.NewLayuiRet(-3, "参数不合法-3",0,nil), nil
            }
            if success, msg :=user.RoleEdit(list, remark, id, pId, roleName, info); success{
                auth.NewCasbin().ReloadPolicy()
                utils.UpdateRoleStatus()
                return base.NewLayuiRet(0, "提交成功",0,nil), nil
            }else{
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}