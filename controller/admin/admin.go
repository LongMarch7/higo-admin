package admin

import (
    "bytes"
    "github.com/LongMarch7/higo-admin/models/admin"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/validator"
    "github.com/LongMarch7/higo/view"
    "golang.org/x/net/context"
    "google.golang.org/grpc/grpclog"
    "runtime"
    "time"
)

type adminController struct {
}
var Controller = &adminController{}
var name = "admin"
func Init(){
    base.AddController(name, Controller)
}

func (a* adminController)Index(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            for _, user := range info {
                if err := validator.Validate.Struct(user); err != nil {
                    return base.JumpToUrl("未登录", "/admin/login")
                }
            }
            out := &bytes.Buffer{}
            data := make(map[string]interface{})
            data["CurrUserMenu"] = admin.GetMenuViewList(info)
            grpclog.Info(data["CurrUserMenu"])
            data["AppName"] = "higo-admin"
            data["Version"] = "1.0.1"
            data["UserName"] = param.Cookie.U
            view.NewView().Render(out, name+"/index", data)
            return out.String(), nil
        }
    }
    return base.JumpToUrl("未登录", "/admin/login")
}

func (a* adminController)Welcome(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            data["AppName"]  = "higo-admin"
            data["Version"]  = "1.0.1"
            data["UserName"] = param.Cookie.U
            data["CurrTime"] =  time.Now().Format("2006/1/2 15:04:05")
            data["OS"] = runtime.GOOS
            data["GOVersion"] = runtime.Version()
            data["Author"] = "Huangyp"
        }
    }
    view.NewView().Render(out, name + "/welcome",data)
    return out.String(), nil
}

