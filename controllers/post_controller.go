package controllers

type PostController struct {
	BaseController
}

func (this *PostController) GetPosts() {

	var str=[]byte(`[{"title":"Sophia","author":"","imgUrl":"","githubUrl":"female","pluginUrl":""},{"title":"Sophia","author":"","imgUrl":"","githubUrl":"female","pluginUrl":""}]`)

	this.RetError(&ControllerError{200, 0, "成功", string(str), ""})
}



