package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	config "praktikum/config"
	model "praktikum/models"
	util "praktikum/utils"
)

func FindBlog(paramId string) map[string]interface{} {
	var blog model.Blog

	blogId, err := strconv.Atoi(paramId)

	if err != nil {
		return util.Response(400, "Bad Request!")
	}

	result := config.DB.First(&blog, blogId)

	if result.RowsAffected < 1 {
		return util.Response(404, "Not Found!")
	}

	return map[string]interface{}{
		"status": 200,
		"blog":   blog,
		"id":     blogId,
	}
}

func GetBlogs(ctx echo.Context) error {
	var blogs []model.Blog

	err := config.DB.Find(&blogs).Error

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, err.Error()))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "blogs Listed!",
		"blogs": blogs,
	})
}

func CreateBlog(ctx echo.Context) error {
	var blog model.Blog

	ctx.Bind(&blog)

	result := config.DB.Create(&blog)

	if result.Error != nil || result.RowsAffected < 1 {
		ctx.JSON(http.StatusInternalServerError, util.Response(500, result.Error.Error()))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "blog Created!",
		"blog": blog,
	})
}

func GetBlog(ctx echo.Context) error {
	blog := FindBlog(ctx.Param("id"))

	if blog["status"] != http.StatusOK {
		return ctx.JSON(blog["status"].(int), blog)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"status": 200,
		"message": "blog Found!",
		"blog": blog["blog"],
	})
}

func UpdateBlog(ctx echo.Context) error {
	blog := FindBlog(ctx.Param("id"))

	if blog["status"] != http.StatusOK {
		return ctx.JSON(blog["status"].(int), blog)
	}

	var newblogData model.Blog

	ctx.Bind(&newblogData)

	result := config.DB.Table("blogs").Where("id", blog["id"]).Updates(newblogData)

	if result.RowsAffected < 1 {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, util.Response(200, "blog Updated!"))
}

func DeleteBlog(ctx echo.Context) error {
	blog := FindBlog(ctx.Param("id"))

	if blog["status"] != http.StatusOK {
		return ctx.JSON(blog["status"].(int), blog)
	}

	result := config.DB.Delete(&model.Blog{}, blog["id"])

	if result.RowsAffected < 1 {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, util.Response(200, "blog Deleted!"))
}