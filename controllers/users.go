package controllers

import (
	"fazztrack/backend/lib"
	"fazztrack/backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllUser(ctx *gin.Context) {
	result := models.FindAllUsers()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Users",
		Results: result,
	})
}

func DetailUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindOneUser(id)

	if dataUser.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "User Found",
			Results: dataUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User Not Found",
		})
	}
}

func CreateUser(ctx *gin.Context) {
	newUser := models.User{}
	result := models.FindAllUsers()

	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	ids := 0
	for _, v := range result {
		ids = v.Id
	}
	newUser.Id = ids + 1

	err := models.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User created successfully",
		Results: newUser,
	})
}

func UpdateUser(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	data := models.FindAllUsers()

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := models.User{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id == 0 {
		c.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "user with id " + param + " not found",
		})
		return
	}

	models.EditUser(user.Email, user.Username, user.Password, param)

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "user with id " + param + " Edit Success",
		Results: user,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindOneUser(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	err = models.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User deleted successfully",
		Results: dataUser,
	})
}
