package setting

import (
    "bytes"
    "github.com/LongMarch7/higo-admin/models/admin/setting"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/controller/base"
    "context"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/view"
    "github.com/LongMarch7/higo-admin/models/admin"
)

func (s* adminSettingController)Site(ctx context.Context) (rs string , err error){
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            siteInfo := utils.SitePostData{}
            setting.GetOptionInfoFromCache(define.OptionNameSite, &siteInfo)
            data["Info"] = siteInfo
            view.NewView().Render(out, name + "/site", data)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (s* adminSettingController)SitePost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok, _ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            site := utils.SitePostData{}
            if utils.FormDataCheck(param.PostFormParams, &site) {
                if setting.SetInfo(define.OptionNameSite, site) {
                    return base.NewLayuiRet(0, "网站信息配置成功",0,nil), nil
                }
            }
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
