package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

type AuthController struct {
	BaseController
}
func (this *AuthController) Login() {
	phoneNumber := this.GetString("phoneNumber")
    password := this.GetString("password")

    beego.Debug("Login:"+phoneNumber+":"+password)

	var token = MD5(phoneNumber,password)

	this.RetError(&ControllerError{200, 0, "成功", string(token), ""})
}

func MD5(userName string,password string) string  {
		h := md5.New()
		h.Write([]byte(userName+password))
		return hex.EncodeToString(h.Sum(nil))
}



