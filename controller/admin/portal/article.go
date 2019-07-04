package portal

import (
    "bytes"
    "encoding/json"
    "github.com/LongMarch7/higo-admin/models/admin/portal"
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

func (s* adminPortalController)ArticleIndex(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            out := &bytes.Buffer{}
            view.NewView().Render(out, name + "/article", nil)
            return out.String(), nil
        }
    }
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    data["content"] = "网络繁忙"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}

func (s* adminPortalController)ArticleList(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            where := make(map[string]string)
            where["article_title"] 	= param.GetParams.Get("article_title")
            where["tag_name"] 	= param.GetParams.Get("tag_name")
            where["cate_name"] 	= param.GetParams.Get("cate_name")
            timeFlag := param.GetParams.Get("time_flag")
            startTime := param.GetParams.Get("start_time")
            endTime := param.GetParams.Get("end_time")
            if strings.Compare(timeFlag,"create") == 0 {
                where["create_start_time"] = startTime
                where["create_end_time"]   = endTime
            }else if strings.Compare(timeFlag,"publish") == 0{
                where["publish_start_time"] = startTime
                where["publish_end_time"]   = endTime
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
            articleList, count:= portal.GetArticleList(where, page, row)
            return base.NewLayuiRet(0, "获取成功",count, articleList), nil
        }
    }
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
func (s* adminPortalController)ArticleDelete(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ :=admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id   := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("article id 不合法",err.Error())
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if succsess,msg :=portal.ArticleDelete(id); succsess{
                return base.NewLayuiRet(0, "删除成功", 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg, 0, nil), nil
            }
        }
    }
    grpclog.Error("ArticleDelete参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminPortalController)ArticleStatusChange(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            id := param.PostFormParams["id"]
            if err :=validator.Validate.Var(&id, validator.ArrayNumeric); err != nil {
                grpclog.Error("ArticleStatusChange id validate err：", err.Error())
                return base.NewLayuiRet(-6, "参数不合法-6",0,nil), nil
            }
            status := param.PostFormParams.Get("status")
            sta,err :=strconv.Atoi(status)
            if err != nil{
                grpclog.Error("ArticleStatusChange status err：", err.Error())
                return base.NewLayuiRet(-5, "参数不合法-5",0,nil), nil
            }else if (sta <1) || (sta >3){
                grpclog.Error("ArticleStatusChange status2 err")
                return base.NewLayuiRet(-4, "参数不合法-4",0,nil), nil
            }
            if success,msg :=portal.ArticleStatusChange(id, sta); success{
                var message = "文章:"
                for _,value := range id{
                    message += "[" + value + "]"
                }
                if sta == 1 {
                    message += "已发布"
                }else{
                    message += "未发布"
                }
                return base.NewLayuiRet(0, message, 0, nil), nil
            }else{
                return base.NewLayuiRet(-2, msg,0,nil), nil
            }

        }
    }
    grpclog.Error("ArticleStatusChange parameter err")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}

func (s* adminPortalController)ArticleEdit(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            articleId := param.GetParams.Get("id")
            var articleInfo models.MicroArticle
            var nowCate portal.Category
            var tagList []models.MicroTag
            var cateList []models.MicroCategory
            nowCate, articleInfo , tagList , cateList = portal.GetArticleInfo(articleId)
            data["ArticleInfo"] = articleInfo
            data["NowCate"] = nowCate
            data["TagList"] = tagList
            data["CateList"] = formateCateList(cateList)
            grpclog.Info( data["CateList"])
            view.NewView().Render(out, name+"/article_edit", data)
            return out.String(), nil
        }
    }
    data["content"] = "操作异常"
    view.NewView().Render(out,"error", data)
    return out.String(), nil
}
type CateObj struct{
    Name     string    `json:"name"`
    Id       int64     `json:"id"`
    Href     string    `json:"href"`
    Alias    string    `json:"alias"`
    Spread   bool      `json:"spread"`
    Children []CateObj `json:"children"`
}
func formateCateList(cate []models.MicroCategory)(list string){
    cateList := make([]CateObj,0)
    for _,value1 := range cate {
        if value1.Level == 1 {
            level1 := CateObj{Name:value1.Name,Id:value1.Id,Alias:value1.Name,Spread:true,Children:make([]CateObj,0)}
            for _,value2 := range cate{
                if value2.ParentId == value1.Id{
                    level2 := CateObj{Name:value2.Name,Id:value2.Id,Alias:value2.Name,Spread:true,Children:make([]CateObj,0)}
                    for _,value3 := range cate{
                        if value3.ParentId == value2.Id{
                            level2.Children = append(level2.Children, CateObj{Name:value3.Name,Id:value3.Id,Alias:value3.Name,Spread:true})
                        }
                    }
                    level1.Children = append(level1.Children, level2)
                }
            }
            cateList = append(cateList, level1)
        }
    }
    if listStr,err := json.Marshal(cateList); err == nil{
        return string(listStr)
    }else{
        return "{}"
    }
}

func (s* adminPortalController)ArticleEditPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok,info := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            article := utils.ArticlePostData{}
            if !utils.FormDataCheck(param.PostFormParams, &article){
               return base.NewLayuiRet(-4, "参数错误",0,nil), nil
            }
            var userId int64
            if len(info) > 0{
                userId = info[0].UserId
            }
            if success, msg :=portal.ArticleEdit(article, userId); success{
               return base.NewLayuiRet(0, "提交成功",0,nil), nil
            }else{
               return base.NewLayuiRet(-2, msg,0,nil), nil
            }
        }
    }
    grpclog.Error("ArticleEditPost参数错误-2")
    return base.NewLayuiRet(-2, "参数错误",0,nil), nil
}
