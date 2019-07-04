package upload

import (
    "encoding/json"
    "github.com/LongMarch7/higo/base"
    "github.com/LongMarch7/higo/util/global"
    "google.golang.org/grpc/grpclog"
    "io"
    "net/http"
    "os"
    "path"
    "strconv"
    "time"
)

func Upload(data string, err error, res http.ResponseWriter, req *http.Request){
    if err != nil {
        base.JsonRender(res, []byte("{\"code\": -1, \"msg\": \"解析错误\",\"data\":\"/error\"}"))
        return
    }
    //保存上传的图片
    _, h, err := req.FormFile("file")
    if err != nil {
        grpclog.Error(err)
        base.JsonRender(res, []byte("{\"code\": -1, \"msg\": \"解析错误\",\"data\":\"/error\"}"))
        return
    }
    fileName := h.Filename
    fileSuffix := path.Ext(fileName)
    newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
    date_url := time.Now().Format("2006-01-02") + "/"
    err = os.MkdirAll(global.UploadPath + date_url, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
    if err != nil {
        grpclog.Error(err)
    }
    path1 := global.UploadPath + date_url + newname //h.Filename
    Url := "/upload/" + date_url + newname
    err = SaveToFile( req,"file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
    if err != nil {
        grpclog.Error(err)
    }
    content, err := json.Marshal(map[string]interface{}{"state": "SUCCESS", "link": Url, "title": fileName, "original": fileName})
    if err == nil {
        base.JsonRender(res, content)
        return
    }
    base.JsonRender(res, []byte("{\"code\": -1, \"msg\": \"解析错误\",\"data\":\"/error\"}"))
}

func SaveToFile(req *http.Request,fromfile, tofile string) error {
    file, _, err := req.FormFile(fromfile)
    if err != nil {
        return err
    }
    defer file.Close()
    f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer f.Close()
    io.Copy(f, file)
    return nil
}
