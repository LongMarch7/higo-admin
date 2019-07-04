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

func (s* adminSettingController)Mail(ctx context.Context) (rs string , err error){
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            mailInfo := utils.MailPostData{}
            setting.GetOptionInfoFromCache(define.OptionNameMail, &mailInfo)
            data["Info"] = mailInfo
            view.NewView().Render(out, name + "/mail", data)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (s* adminSettingController)MailConf(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok, _ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            mailSetting := utils.MailPostData{}
            if utils.FormDataCheck(param.PostFormParams,&mailSetting) {
                if setting.SetInfo(define.OptionNameMail, mailSetting) {
                    return base.NewLayuiRet(0, "邮箱配置成功",0,nil), nil
                }
            }
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminSettingController)MailTest(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok, _ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            mailSend :=  utils.MailSendData{}
            if utils.FormDataCheck(param.PostFormParams, &mailSend) {
                if setting.SendMail([]string{mailSend.Receive}, mailSend.SendSubject, mailSend.SendSubject){
                    return base.NewLayuiRet(0, "邮箱发送成功",0,nil), nil
                }
            }
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}