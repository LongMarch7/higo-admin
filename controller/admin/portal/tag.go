package portal

import (
    "bytes"
    "context"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/LongMarch7/higo/view"
    "github.com/LongMarch7/higo-admin/models/admin"
    "google.golang.org/grpc/grpclog"
    "github.com/LongMarch7/higo-admin/models/admin/portal"
    "strconv"
    "strings"
)

func (s* adminPortalController)TagIndex(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            view.NewView().Render(out, name + "/tag", nil)
            return out.String(), nil
        }
    }
    grpclog.Error("TagIndex parameter err")
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}


func (s* adminPortalController)TagList(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            where := make(map[string]string)
            where["tag_name"] 	= param.GetParams.Get("tag_name")
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
            linkList, count:= portal.GetTagList(where, page, row)
            return base.NewLayuiRet(0, "获取成功",count, linkList), nil
        }
    }
    grpclog.Error("TagList parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
func (s* adminPortalController)TagDelete(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id   := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("Tag id 不合法",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if succsess,msg :=portal.TagDelete(id); succsess{
                return base.NewLayuiRet(0, "删除成功", 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg, 0, nil), nil
            }
        }
    }
    grpclog.Error("TagDelete parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminPortalController)TagStatusChange(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("TagStatusChange id validate err：", err.Error())
                return base.NewLayuiRet(-6, "参数不合法-6",0,nil), nil
            }
            status := param.PostFormParams.Get("status")
            sta,err :=strconv.Atoi(status)
            if err != nil{
                grpclog.Error("TagStatusChange status err：", err.Error())
                return base.NewLayuiRet(-5, "参数不合法-5",0,nil), nil
            }else if (sta <1) || (sta >3){
                grpclog.Error("TagStatusChange status2 err")
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if success,msg :=portal.TagStatusChange(id, sta); success{
                var message = "标签:"
                for _,value := range id{
                    message += "[" + value + "]"
                }
                if sta == 1 {
                    message += "发布"
                }else{
                    message += "不发布"
                }
                return base.NewLayuiRet(0, message, 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }

        }
    }
    grpclog.Error("TagStatusChange parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminPortalController)TagEditPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            grpclog.Info(param.PostFormParams)
            tag := utils.TagPostData{}
            if !utils.FormDataCheck(param.PostFormParams, &tag){
                grpclog.Error("TagEditPost FormDataCheck err")
                return base.NewLayuiRet(-4, "参数不合法",0,nil), nil
            }
            if success, msg :=portal.TagEdit(tag); success{
                return base.NewLayuiRet(0, "提交成功",0,nil), nil
            }else{
                grpclog.Error("TagEditPost model set err")
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }
            return base.NewLayuiRet(0, "提交成功",0,nil), nil
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
