package controllers

import (
	"encoding/json"
	"nomadiclife/helper"
	"nomadiclife/models"

	beego "github.com/beego/beego/v2/server/web"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

type Result struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
}

type Email struct {
	Email string `json:"email"`
}

func (c *UserController) Prepare() {

}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create User
// @Param	first_name	json string	true		"body for User content"
// @Param	last_name	json string	true		"body for User content"
// @Param	email  json string	true		"body for User content"
// @Param	username json 	string	true		"body for User content"
// @Param	password json 	string	true		"body for User content"
// @Param	phone_number json 	string	true		"body for User content"
// @Success 201 {object} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var user models.User
	form := c.Ctx.Input.RequestBody
	err := json.Unmarshal(form, &user)
	if err != nil {
		helper.FailureResponse(err.Error(), c.Controller)
	}
	validation_error := user.Validate()
	if validation_error != nil {
		validation_err_msg, _ := json.Marshal(validation_error)
		helper.FailureResponse(string(validation_err_msg), c.Controller)
	} else {
		user.Password, _ = models.HashPassword(user.Password)
		id, err := models.AddUser(&user)
		if err == nil {
			helper.SuccessResponse("User successfully registered", id, c.Controller)
		} else {
			helper.FailureResponse(err.Error(), c.Controller)
		}
	}
}

// GetOne ...
// @Title GetOne
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {

}

// Forget password ...
// @Title Forgetpassword
// @Description Forget password feature
// @Param	email	body string	true
func (c *UserController) ForgetPassword() {
	email := Email{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &email)
	if err := models.ForgetPassword(email.Email); err != nil {
		helper.FailureResponse(err.Error(), c.Controller)
	}
	helper.SuccessResponse("New password successfully sent to given address", email, c.Controller)
}
