package member

import (
    "bytes"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo-admin/models/admin"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/LongMarch7/higo/view"
    "google.golang.org/grpc/grpclog"
    "github.com/LongMarch7/higo-admin/models/admin/member"
    "strconv"
    "context"
    "strings"
)

type adminMemberController struct {
}
var Controller = &adminMemberController{}
var name = "admin/member"
func Init(){
    base.AddController(name, Controller)
}

func (u* adminMemberController)Index(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            view.NewView().Render(out, name + "/index", nil)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (u* adminMemberController)List(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
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
            }else if strings.Compare(timeFlag,"freeze") == 0 {
                where["freeze_start_time"] = startTime
                where["freeze_end_time"]   = endTime
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
            userList, count:= member.GetMemberList(where, page, row)
            return base.NewLayuiRet(0, "获取成功",count, userList), nil
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
func (u* adminMemberController)Delete(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id   := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("member id 不合法",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if succsess,msg :=member.Delete(id); succsess{
                return base.NewLayuiRet(0, "删除成功", 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg, 0, nil), nil
            }
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (u* adminMemberController)StatusChange(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("memberStatusChange id validate err：", err.Error())
                return base.NewLayuiRet(-6, "参数不合法-6",0,nil), nil
            }
            status := param.PostFormParams.Get("status")
            sta,err :=strconv.Atoi(status)
            if err != nil{
                grpclog.Error("memberStatusChange status err：", err.Error())
                return base.NewLayuiRet(-5, "参数不合法-5",0,nil), nil
            }else if (sta <1) || (sta >3){
                grpclog.Error("memberStatusChange status2 err")
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if success,msg :=member.StatusChange(id, sta); success{
                var message = "会员:"
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
    grpclog.Error("memberStatusChange parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (u* adminMemberController)Edit(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            userId := param.GetParams.Get("id")
            var userInfo models.MicroUser
            if len(userId) > 0{
                userInfo = member.GetMemberInfo(userId)
            }
            data["UserInfo"] = userInfo
            data["IsSuper"] = false
            for _,user := range info{
                if user.RoleId == define.SuperRoleId {
                    data["IsSuper"] = true
                    break
                }
            }
            view.NewView().Render(out, name+"/edit", data)
            return out.String(), nil
        }
    }
    data["content"] = "操作异常"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (u* adminMemberController)EditPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            userInfo := utils.MemberPostData{}
            if !utils.FormDataCheck(param.PostFormParams, &userInfo){
                return base.NewLayuiRet(-4, "参数错误",0,nil), nil
            }
            isSuper := false
            for _,user := range info{
                if user.RoleId == define.SuperRoleId {
                    isSuper = false
                    break
                }
            }
            if success, msg :=member.Edit(userInfo,isSuper); success{
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