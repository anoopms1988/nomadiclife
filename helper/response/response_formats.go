package response

import (
	"encoding/base64"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type Helper struct {
}

type Res struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Msg    string      `json:"message"`
	Data   interface{} `json:"data"`
}

type FailureRes struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"message"`
}

func Response(status bool, code int, msg string, data interface{}, u beego.Controller) {
	u.Data["json"] = Res{status, code, msg, data}
	u.Ctx.ResponseWriter.WriteHeader(code)
	u.ServeJSON()
}

func SuccessResponse(msg string, data interface{}, u beego.Controller) {
	u.Data["json"] = Res{true, 200, msg, data}
	u.Ctx.ResponseWriter.WriteHeader(200)
	u.ServeJSON()
}

func FailureResponse(msg string, u beego.Controller) {
	u.Data["json"] = FailureRes{false, 400, msg}
	u.Ctx.ResponseWriter.WriteHeader(400)
	u.ServeJSON()
}

func ImageResponse(data []byte, c beego.Controller) {
	c.Data["json"] = &data
	c.ServeJSON()
}

func EncodeBase64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return string(base64Text)
}

func DecodeBase64(message string) (retour []byte) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	i, _ := base64.StdEncoding.DecodeString(message)
	fmt.Printf("base64: %s\n", i)
	return base64Text
}

func TimeNow() (t string) {
	return time.Now().Format("02/01/2006 15:04:05")
}
