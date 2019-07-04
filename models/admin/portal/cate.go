package portal

import (
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/db"
    "google.golang.org/grpc/grpclog"
    "strings"
    "time"
    "github.com/LongMarch7/higo-admin/db/object/models"
)


func GetCateList(is_delete bool)  (cate_list []models.MicroCategory){
    engine := db.NewDb(db.DefaultNAME).Engine()
    cate_list = make([]models.MicroCategory,0)
    var err error
    if is_delete{
        err = engine.Table(models.MicroCategory{}.TableName()).Where("delete_time > ?", 0).Find(&cate_list)
    }else{
        err = engine.Table(models.MicroCategory{}.TableName()).Where("delete_time = ?", 0).Find(&cate_list)
    }
    if err != nil {
        grpclog.Error("get cate list err: ",err.Error())
        return
    }
    return
}
func GetCateInfo(cate_id string)  (cate_list models.MicroCategory){
    engine := db.NewDb(db.DefaultNAME).Engine()
    engine.Table(models.MicroCategory{}.TableName()).ID(cate_id).Get(&cate_list)
    test,_ :=engine.Query("SELECT f.id,s.id,t.id FROM micro_category AS f LEFT JOIN micro_category AS s ON f.id = s.parent_id LEFT JOIN micro_category t  ON s.id = t.parent_id where f.id = 7")
    grpclog.Info(test)
    return
}

func CateDelete(cates_id []string, action string) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg ="执行出错"
    success = false
    if strings.Compare(action,"delete") == 0 || strings.Compare(action,"restore") == 0{
        deleteTime := 0
        errMsg := "删除分类失败"
        if strings.Compare(action,"delete") == 0{
            deleteTime = int(time.Now().Unix())
            errMsg = "恢复分类失败"
        }
        for _,value := range cates_id{
            link :=models.MicroCategory{DeleteTime: deleteTime}
            if _,err :=engine.Table(models.MicroCategory{}.TableName()).ID(value).Cols("delete_time").Update(&link); err != nil{
                grpclog.Error(action + "cate failed !",err.Error())
                msg = value + errMsg
                return
            }
        }
    }else if strings.Compare(action,"clear") == 0{
        for _,value := range cates_id{
            if _,err :=engine.Table(models.MicroCategory{}.TableName()).ID(value).Delete(&models.MicroCategory{}); err != nil{
                grpclog.Error("clear cate failed !",err.Error())
                msg = value + "清除分类失败"
                return
            }
        }
    }else{
        return
    }
    msg = "操作成功"
    success = true
    return
}

func CateStatusChange(cates_id []string, cates_status int) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    cate :=models.MicroCategory{}
    success = false
    msg = "操作出错"
    for _,value := range cates_id{
        cate.Status = cates_status
        cate.UpdateTime = int(time.Now().Unix())
        if _,err :=engine.Table(models.MicroCategory{}.TableName()).ID(value).Cols("status","update_time").Update(&cate); err != nil{
            grpclog.Error("update tag failed ", err.Error())
            msg = value + "标签状态更新失败"
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}
func CateEdit(cate_data utils.CatePostData) (success bool,msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg = "执行错误"
    success = false

    nowTime := int(time.Now().Unix())
    tag :=models.MicroCategory{
        Name: cate_data.Name,
        UpdateTime: nowTime,
        SeoDescription: cate_data.SeoDescription,
        SeoKeywords: cate_data.SeoKeywords,
        SeoTitle: cate_data.SeoTitle,
        Description: cate_data.Description,
        ParentId: cate_data.ParentId,
        Level: cate_data.Level,
    }
    var cate_id = cate_data.Id
    if cate_id >0{
        if  _,err := engine.Table(models.MicroCategory{}.TableName()).ID(cate_id).Update(&tag); err != nil{
            grpclog.Error("Cate update:",err.Error())
            msg = "更新出错"
            return
        }
    }else{
        tag.CreateTime = nowTime
        tag.Status = 1
        has,err :=engine.Table(models.MicroCategory{}.TableName()).Where("name = ?", cate_data.Name).Exist(&models.MicroCategory{})
        if err == nil && has{
            grpclog.Error("Cate exist")
            msg = cate_data.Name + "分类名已存在"
            return
        }
        if  _,err := engine.Table(models.MicroCategory{}.TableName()).Insert(&tag); err != nil{
            grpclog.Error("Cate insert:",err.Error())
            return
        }
    }
    success = true
    return
}
