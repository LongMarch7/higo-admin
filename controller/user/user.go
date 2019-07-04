package user

import (
    "bytes"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/view"
    "golang.org/x/net/context"
)

type userController struct {
}
var controller = &userController{}
var name = "user"
func Init(){
    base.AddController(name, controller)
}

func (a* userController)Index(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    view.NewView().Render(out, name + "/index",nil)
    return out.String(), nil
}

