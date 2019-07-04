package setting

import (
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/util/define"
    "google.golang.org/grpc/grpclog"
    "gopkg.in/gomail.v2"
    "strconv"
)

func SendMail(mailTo []string,subject string, body string) bool{
    m := gomail.NewMessage()
    mailInfo := utils.MailPostData{}
    GetOptionInfoFromCache(define.OptionNameMail, &mailInfo)
    p, _ := strconv.Atoi(mailInfo.Cache)
    dialer := gomail.NewDialer(mailInfo.SmtpServer, p, mailInfo.SendMail, mailInfo.SendPwd)

    if len(mailInfo.SendNickname) > 0{
        m.SetHeader("From",mailInfo.SendNickname + "<" + mailInfo.SendMail + ">")
    }else{
        m.SetHeader("From",mailInfo.SendMail)
    }
    m.SetHeader("To", mailTo...)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    err := dialer.DialAndSend(m)
    grpclog.Info(err)
    return err == nil
}