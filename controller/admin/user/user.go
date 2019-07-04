package user

import (
    "bytes"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo-admin/models/admin"
    "github.com/LongMarch7/higo-admin/models/admin/user"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/LongMarch7/higo/view"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "context"
    "strings"
)

type adminUserController struct {
}
var Controller = &adminUserController{}
var name = "admin/user"
func Init(){
    base.AddController(name, Controller)
}

func (u* adminUserController)UserIndex(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            view.NewView().Render(out, name + "/user_index", nil)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (u* adminUserController)UserList(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            where := make(map[string]string)
            where["user_name"] 	= param.GetParams.Get("user_name")
            timeFlag := param.GetParams.Get("time_flag")
            startTime := param.GetParams.Get("start_time")
            endTime := param.GetParams.Get("end_time")
            if strings.Compare(timeFlag,"create") == 0 {
                where["create_start_time"] = startTime
                where["create_end_time"]   = endTime
            }else if strings.Compare(timeFlag,"login") == 0 {
                where["login_start_time"] = startTime
                where["login_end_time"]   = endTime
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
            userList, count:= user.GetUserList(where, page, row, info, param.Cookie.U)
            return base.NewLayuiRet(0, "获取成功",count, userList), nil
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
func (u* adminUserController)UserDelete(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id   := param.PostFormParams["id"]
            name := param.PostFormParams["name"]
            if err :=validator.Validate.Var(&name, validator.ArrayAlphanum); err != nil {
                grpclog.Error("user name 不合法",err.Error())
                return base.NewLayuiRet(-5, "参数错误-5",0,nil), nil
            }
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("user id 不合法",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if succsess,msg :=user.DelUsers(id, name, info); succsess{
                return base.NewLayuiRet(0, "删除成功", 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg, 0, nil), nil
            }
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (u* adminUserController)UserStatusChange(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("UserStatusChange id validate err：", err.Error())
                return base.NewLayuiRet(-6, "参数不合法-6",0,nil), nil
            }
            status := param.PostFormParams.Get("status")
            sta,err :=strconv.Atoi(status)
            if err != nil{
                grpclog.Error("UserStatusChange status err：", err.Error())
                return base.NewLayuiRet(-5, "参数不合法-5",0,nil), nil
            }else if (sta <1) || (sta >3){
                grpclog.Error("UserStatusChange status2 err")
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if success,msg :=user.UsersStatusChange(id, sta, info); success{
                var message = "用户:"
                for _,value := range id{
                    message += "[" + value + "]"
                }
                if sta == 1 {
                    message += "已启用"
                }else{
                    message += "已禁用"
                }
                return base.NewLayuiRet(0, message, 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }

        }
    }
    grpclog.Error("UserStatusChange parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (u* adminUserController)UserEdit(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            userId := param.GetParams.Get("user_id")
            var userInfo models.MicroUser
            rolesMap := user.GetAuthorizedRoleList(info)
            if len(userId) > 0{
                if success,msg :=user.UserPowerCheck(info, userId); !success{
                    data["content"] = msg
                    view.NewView().Render(out,"error", data)
                    return out.String(), nil
                }
                userInfo = utils.GetUserInfo(userId)
                rolesMap = user.HasPower(userId, rolesMap)
            }
            //data["CurrUserMenu"] = user.MenuPowerCheck(admin.GetMenuList(setteruserName), setteruserName, userName)
            data["UserInfo"] = userInfo
            data["RolesMap"] = rolesMap
            view.NewView().Render(out, name+"/user_edit", data)
            return out.String(), nil
        }
    }
    data["content"] = "操作异常"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (u* adminUserController)UserEditPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            userInfo := utils.UserPostData{}
            if !utils.FormDataCheck(param.PostFormParams, &userInfo){
                return base.NewLayuiRet(-4, "参数错误",0,nil), nil
            }
            noCheck :=param.PostFormParams.Get("NoCheck")
            var noCheckSplit []string
            if len(noCheck) > 0 {
                noCheckSplit = strings.Split(noCheck,",")
            }
            rolesMap := user.GetAuthorizedRoleList(info)

            if success, msg :=user.UserEdit(userInfo, rolesMap, noCheckSplit, info); success{
                return base.NewLayuiRet(0, "提交成功",0,nil), nil
            }else{
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }
            return base.NewLayuiRet(0, "提交成功",0,nil), nil
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}