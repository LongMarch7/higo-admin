package setting

import (
    "bytes"
    "github.com/LongMarch7/higo-admin/models/admin/setting"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/LongMarch7/higo/view"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "context"
    "github.com/LongMarch7/higo-admin/models/admin"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "strings"
)

func (s* adminSettingController)Link(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            view.NewView().Render(out, name + "/link", nil)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (s* adminSettingController)LinkList(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            where := make(map[string]string)
            where["link_name"] 	= param.GetParams.Get("link_name")
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
            linkList, count:= setting.GetLinkList(where, page, row)
            return base.NewLayuiRet(0, "获取成功",count, linkList), nil
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
func (s* adminSettingController)LinkDelete(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id   := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("link id 不合法",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if succsess,msg :=setting.LinkDelete(id); succsess{
                return base.NewLayuiRet(0, "删除成功", 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg, 0, nil), nil
            }
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminSettingController)LinkStatusChange(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("LinkStatusChange id validate err：", err.Error())
                return base.NewLayuiRet(-6, "参数不合法-6",0,nil), nil
            }
            status := param.PostFormParams.Get("status")
            sta,err :=strconv.Atoi(status)
            if err != nil{
                grpclog.Error("LinkStatusChange status err：", err.Error())
                return base.NewLayuiRet(-5, "参数不合法-5",0,nil), nil
            }else if (sta <1) || (sta >3){
                grpclog.Error("LinkStatusChange status2 err")
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if success,msg :=setting.LinkStatusChange(id, sta); success{
                var message = "友情链接:"
                for _,value := range id{
                    message += "[" + value + "]"
                }
                if sta == 1 {
                    message += "已显示"
                }else{
                    message += "已隐藏"
                }
                return base.NewLayuiRet(0, message, 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }

        }
    }
    grpclog.Error("LinkStatusChange parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminSettingController)LinkEdit(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            userId := param.GetParams.Get("id")
            var linkInfo models.MicroLink
            if len(userId) > 0{
                linkInfo = setting.GetLinkInfo(userId)
            }
            data["LinkInfo"] = linkInfo
            view.NewView().Render(out, name+"/link_edit", data)
            return out.String(), nil
        }
    }
    data["content"] = "操作异常"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (s* adminSettingController)LinkEditPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            grpclog.Info(param.PostFormParams)
            link := utils.LinkPostData{}
            if !utils.FormDataCheck(param.PostFormParams, &link){
               return base.NewLayuiRet(-4, "参数错误",0,nil), nil
            }
            if success, msg :=setting.LinkEdit(link); success{
               return base.NewLayuiRet(0, "提交成功",0,nil), nil
            }else{
               return base.NewLayuiRet(-2, msg,0,nil), nil
            }
        }
    }
    grpclog.Error("参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
