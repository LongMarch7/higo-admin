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

func (s* adminSettingController)Upload(ctx context.Context) (rs string , err error){
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            uploadInfo := utils.UploadPostData{}
            setting.GetOptionInfoFromCache(define.OptionNameUpload, &uploadInfo)
            data["Info"] = uploadInfo
            view.NewView().Render(out, name + "/upload", data)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (s* adminSettingController)UploadPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok, _ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            upload := utils.UploadPostData{}
            if utils.FormDataCheck(param.PostFormParams, &upload) {
                if setting.SetInfo(define.OptionNameUpload, upload) {
                    return base.NewLayuiRet(0, "上传设置成功",0,nil), nil
                }
            }
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
