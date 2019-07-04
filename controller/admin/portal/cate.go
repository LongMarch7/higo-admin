package portal

import (
    "bytes"
    "context"
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/LongMarch7/higo/view"
    "github.com/LongMarch7/higo-admin/models/admin"
    "google.golang.org/grpc/grpclog"
    "github.com/LongMarch7/higo-admin/models/admin/portal"
    "strconv"
)

func (s* adminPortalController)CateIndex(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            data["CateList"] = portal.GetCateList(false)
            grpclog.Info(data["CateList"])
            view.NewView().Render(out, name + "/cate", data)
            return out.String(), nil
        }
    }
    data["content"] = "网络繁忙"
    grpclog.Error("CateIndex parameter err")
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}
func (s* adminPortalController)CateRestore(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            data["CateList"] = portal.GetCateList(true)
            grpclog.Info(data["CateList"])
            view.NewView().Render(out, name + "/cate_restore", data)
            return out.String(), nil
        }
    }
    data["content"] = "网络繁忙"
    grpclog.Error("CateRestore parameter err")
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (s* adminPortalController)CateDelete(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("Cate id 不合法",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            action := param.PostFormParams.Get("action")
            if succsess,msg :=portal.CateDelete(id, action); succsess{
                return base.NewLayuiRet(0, "删除成功", 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg, 0, nil), nil
            }
        }
    }
    grpclog.Error("CateDelete parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminPortalController)CateStatusChange(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("CateStatusChange id validate err：", err.Error())
                return base.NewLayuiRet(-6, "参数不合法-6",0,nil), nil
            }
            status := param.PostFormParams.Get("status")
            sta,err :=strconv.Atoi(status)
            if err != nil{
                grpclog.Error("CateStatusChange status err：", err.Error())
                return base.NewLayuiRet(-5, "参数不合法-5",0,nil), nil
            }else if (sta <1) || (sta >3){
                grpclog.Error("CateStatusChange status2 err")
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if success,msg :=portal.CateStatusChange(id, sta); success{
                var message = "分类:"
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
    grpclog.Error("CateStatusChange parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminPortalController)CateEdit(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            cateId := param.GetParams.Get("id")
            var cateInfo models.MicroCategory
            if len(cateId) > 0{
                cateInfo = portal.GetCateInfo(cateId)
            }
            parentId := param.GetParams.Get("parent_id")
            p_id, err := strconv.ParseInt(parentId, 10, 64)
            if err == nil{
                cateInfo.ParentId = p_id
            }
            cateInfo.Level = 1
            levelStr := param.GetParams.Get("level")
            level,err :=strconv.Atoi(levelStr)
            if err == nil{
                cateInfo.Level = level
            }
            data["CateInfo"] = cateInfo
            view.NewView().Render(out, name+"/cate_edit", data)
            return out.String(), nil
        }
    }
    data["content"] = "操作异常"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}


func (s* adminPortalController)CateEditPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            grpclog.Info(param.PostFormParams)
            tag := utils.CatePostData{}
            if !utils.FormDataCheck(param.PostFormParams, &tag){
                grpclog.Error("CateEditPost FormDataCheck err")
                return base.NewLayuiRet(-4, "参数错误",0,nil), nil
            }
            if success, msg :=portal.CateEdit(tag); success{
                return base.NewLayuiRet(0, "提交成功",0,nil), nil
            }else{
                grpclog.Error("CateEditPost model set err")
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }
            return base.NewLayuiRet(0, "提交成功",0,nil), nil
        }
    }
    grpclog.Error("CateEditPost-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
