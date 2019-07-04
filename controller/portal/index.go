package portal

import (
    "bytes"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/view"
    "golang.org/x/net/context"
)

type portalController struct {
}
var controller = &portalController{}
var name = "portal"
func Init(){
    base.AddController(name, controller)
}

func (a* portalController)Index(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    view.NewView().Render(out, name + "/index",nil)
    return out.String(), nil
}

