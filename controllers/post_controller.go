package controllers

// JobController operations for Job
/*职位相关*/
type PostController struct {
	BaseController
}

/*获取职位详情页*/
func (this *PostController) GetPosts() {

	var str=[]byte(`[{"title":"Sophia","author":"","imgUrl":"","githubUrl":"female","pluginUrl":""},{"title":"Sophia","author":"","imgUrl":"","githubUrl":"female","pluginUrl":""}]`)

	this.RetError(&ControllerError{200, 0, "成功", string(str), ""})
}



