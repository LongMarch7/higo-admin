package portal

import (
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/util/define"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "time"
    "github.com/LongMarch7/higo-admin/db/object/models"
)

func GetSysArticleListWhereSql(where map[string]string) (string, []interface{}) {
    var sql = "1 = 1"
    var value []interface{}
    if v, ok := where["article_title"]; ok && v != "" {
        keywords := define.TrimString(v)
        sql += " AND article_title like ?"
        value = append(value,"%"+keywords+"%")
    }

    if v, ok := where["create_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND create_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["create_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND create_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["update_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND update_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["update_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND update_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["publish_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND published_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["publish_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND published_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["status_flag"]; ok && v != "" {
        status, _ :=strconv.Atoi(v)
        if status >0 && status <3{
            sql += " AND article_status = ?"
            value = append(value,status)
        }
    }
    return sql, value
}

func GetArticleList(where map[string]string, page int, rows int)  (article_list []models.MicroArticle, count int){
    engine := db.NewDb(db.DefaultNAME).Engine()
    count = 0
    sql , value:= GetSysArticleListWhereSql(where)
    article_list = make([]models.MicroArticle, 0)
    err := engine.Table(models.MicroArticle{}.TableName()).Where(sql, value...).Limit(rows,(page-1)*rows).
        Find(&article_list)
    if err != nil {
        grpclog.Error("get article list err: ",err.Error())
        return
    }
    ret,err :=engine.Table(models.MicroArticle{}.TableName()).Where(sql, value...).Count(&models.MicroArticle{})
    if err != nil {
        grpclog.Error("get article count err: ",err.Error())
        return
    }
    count = int(ret)
    return
}

func ArticleDelete(articles_id []string) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    session := engine.NewSession()
    defer session.Close()
    msg ="执行出错"
    success = false
    for _,value := range articles_id{
        err := session.Begin()
        if _,err :=session.Table(models.MicroArticle{}.TableName()).ID(value).Delete(&models.MicroArticle{}); err != nil{
            session.Rollback()
            grpclog.Error("delete article failed !",err.Error())
            msg = value + "删除文章失败"
            return
        }
        if _,err :=session.Table(models.MicroCategoryArticle{}.TableName()).Where("article_id = ?", value).Delete(&models.MicroCategoryArticle{}); err != nil{
            session.Rollback()
            grpclog.Error("delete article cate failed !",err.Error())
            msg = value + "删除文章分类映射失败"
            return
        }
        if _,err :=session.Table(models.MicroTagArticle{}.TableName()).Where("article_id = ?", value).Delete(&models.MicroTagArticle{}); err != nil{
            session.Rollback()
            grpclog.Error("delete article tag failed !",err.Error())
            msg = value + "删除文章标签映射失败"
            return
        }
        if err = session.Commit();err != nil{
            session.Rollback()
            grpclog.Error("delete article failed by session !",err.Error())
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}

func ArticleStatusChange(articles_id []string, articles_status int) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    article :=models.MicroArticle{}
    success = false
    msg = "操作出错"
    for _,value := range articles_id{
        article.ArticleStatus = articles_status
        article.UpdateTime = int(time.Now().Unix())
        if articles_status == 2 {
            article.PublishedTime = int(time.Now().Unix())
        }
        if _,err :=engine.Table(models.MicroArticle{}.TableName()).ID(value).Cols("article_status").Update(&article); err != nil{
            grpclog.Error("update article failed ", err.Error())
            msg = value + "文章状态更新失败"
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}
type Category struct{
    CategoryId int64
    Name       string
}
type SelectedTag struct {
    TagId      int64
}
func GetArticleInfo(article_id string)(now_cate Category, article_info models.MicroArticle, tag_list []models.MicroTag, cate_list []models.MicroCategory){
    engine := db.NewDb(db.DefaultNAME).Engine()
    now_cate = Category{}
    article_info = models.MicroArticle{}
    tag_list     = make([]models.MicroTag, 0)
    cate_list     = make([]models.MicroCategory, 0)
    tag_select := make([]SelectedTag, 0)
    engine.Table(models.MicroTag{}.TableName()).Find(&tag_list)
    engine.Table(models.MicroCategory{}.TableName()).Find(&cate_list)
    if len(article_id)>0 {
        engine.Table(models.MicroArticle{}.TableName()).Alias("a").Where("a.id =? ", article_id).
            Join("INNER",[]string{models.MicroCategoryArticle{}.TableName(),"ca"},"a.id = ca.article_id").
            Join("INNER",[]string{models.MicroCategory{}.TableName(),"c"},"c.id = ca.category_id").Get(&now_cate)
        engine.Table(models.MicroArticle{}.TableName()).ID(article_id).Get(&article_info)
        engine.Table(models.MicroArticle{}.TableName()).Alias("a").Where("a.id =? ", article_id).
            Join("INNER",[]string{models.MicroTagArticle{}.TableName(),"ta"},"a.id = ta.article_id").Find(&tag_select)
    }
    //复用现存status变量作为选中变量
    for  index,tagValue := range tag_list{
        tag_list[index].Status = 0
        for _,selectValue := range tag_select{
            if selectValue.TagId == tagValue.Id {
                tag_list[index].Status = 1
                break
            }
        }
    }
    //content, _ := base64.StdEncoding.DecodeString(article_info.ArticleContent)
    //article_info.ArticleContent = string(content)
    //excerpt, _ := base64.StdEncoding.DecodeString(article_info.ArticleExcerpt)
    //article_info.ArticleExcerpt = string(excerpt)
    return
}

func ArticleEdit(article_data utils.ArticlePostData, user_id int64) (success bool,msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    session := engine.NewSession()
    defer session.Close()
    msg = "执行错误"
    success = false

    nowTime := int(time.Now().Unix())
    article :=models.MicroArticle{
        ArticleTitle: article_data.ArticleTitle,
        UpdateTime: nowTime,
        ArticleKeywords: article_data.ArticleKeywords,
        ArticleExcerpt: article_data.ArticleExcerpt,
        Thumbnail: article_data.Thumbnail,
        ArticleContent: article_data.ArticleContent,
        ArticleType: 1,
        ArticleFormat: 1,
        UserId: user_id,
        ArticleHits: article_data.ArticleHits,
        ArticleFavorites: article_data.ArticleFavorites,
        ArticleLike: article_data.ArticleLike,
        CommentCount: article_data.CommentCount,
    }
    session.Begin()
    var article_id = article_data.Id
    if article_id >0{
        if  _,err := session.Table(models.MicroArticle{}.TableName()).ID(article_id).Update(&article); err != nil{
            session.Rollback()
            grpclog.Error("article update failed:",err.Error())
            msg = "更新出错"
            return
        }
    }else{
        article.CreateTime = nowTime
        article.PublishedTime = nowTime
        article.ArticleStatus = 1
        if  _,err := session.Table(models.MicroArticle{}.TableName()).Insert(&article); err != nil{
            session.Rollback()
            grpclog.Error("article insert failed:",err.Error())
            msg = "添加出错"
            return
        }
        article_id = article.Id
    }
    cateArticle := models.MicroCategoryArticle{ArticleId:article_id,CategoryId:article_data.Cate}
    has,err :=session.Table(models.MicroCategoryArticle{}.TableName()).Where("article_id = ?", article_id).Exist(&models.MicroCategoryArticle{})
    if err == nil && has{
        if  _,err := session.Table(models.MicroCategoryArticle{}.TableName()).Where("article_id = ?", article_id).Update(&cateArticle); err != nil{
            session.Rollback()
            grpclog.Error("category article update failed:",err.Error())
            msg = "更新文章分类映射出错"
            return
        }
    }else{
        if  _,err := session.Table(models.MicroCategoryArticle{}.TableName()).Insert(&cateArticle); err != nil{
            session.Rollback()
            grpclog.Error("category article insert failed:",err.Error())
            msg = "添加文章分类映射出错"
            return
        }
    }
    if _,err :=session.Table(models.MicroTagArticle{}.TableName()).Where("article_id = ?", article_id).Delete(&models.MicroTagArticle{}); err != nil{
        session.Rollback()
        grpclog.Error("delete article tag failed :",err.Error())
        msg = "删除文章标签映射失败"
        return
    }
    for _,value := range article_data.Tag{
        if _,err :=session.Table(models.MicroTagArticle{}.TableName()).Where("article_id = ?", article_id).Insert(&models.MicroTagArticle{ArticleId:article_id,TagId:value}); err != nil{
            session.Rollback()
            grpclog.Error("insert article tag failed !",err.Error())
            msg = "添加文章标签映射失败"
            return
        }
    }
    if err := session.Commit();err != nil{
        session.Rollback()
        grpclog.Error("delete  failed by session !",err.Error())
        return
    }
    success = true
    return
}
