package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	model "praktikum/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockUserModel struct{}

func (um *MockUserModel) FindUserAccount(user model.User) *model.User {
	return &user
}

func (um *MockUserModel) FindUser(paramId string) map[string]interface{} {
	if paramId == "1" {
		return map[string]interface{} {
			"status": 200,
			"message": "Bad Request!",
			"id": 1,
		}	
	}

	if paramId == "satu" {
		return map[string]interface{} {
			"status": 400,
			"message": "Bad Request!",
		}	
	}

	return map[string]interface{} {
		"status": 404,
		"message": "Not Found!",
	}
}

func (bm *MockUserModel) InsertUser(newUser model.User) *model.User {
	return nil
}

func (bm *MockUserModel) SelectAllUser() []model.User {
	return nil
}

func (bm *MockUserModel) UpdateUser(UserId int, newUser model.User) bool {
	return false
}

func (bm *MockUserModel) DeleteUser(UserId int) bool {
	return false
}

type SuccessMockUserModel struct{}

func (um *SuccessMockUserModel) FindUserAccount(user model.User) *model.User {
	return &user
}

func (um *SuccessMockUserModel) FindUser(paramId string) map[string]interface{} {
	return map[string]interface{} {
		"status": 200,
		"user":   map[string]interface{}{
			"id": 1,
			"username": "sarbinus",
			"email": "sarbin@example.com",
			"password": "siSarbin123",
		},
		"id": 1,
	}	
}

func (bm *SuccessMockUserModel) InsertUser(newUser model.User) *model.User {
	return &newUser
}

func (bm *SuccessMockUserModel) SelectAllUser() []model.User {
	var users = []model.User{{
		Id: 1,
		Username: "sarbinus",
		Email: "sarbin@example.com",
		Password: "siSarbin123",
	}}

	return users
}

func (bm *SuccessMockUserModel) UpdateUser(UserId int, newUser model.User) bool {
	return true
}

func (bm *SuccessMockUserModel) DeleteUser(UserId int) bool {
	return true
}

func TestGetUsers(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Message string         `json:"message"`
		Users   map[string]any `json:"users"`
	}

	var tmp = Response{}


	t.Run("Fetch All Failed", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.GET("/users", controller.GetUsers())

		var req = httptest.NewRequest(http.MethodGet, "/users", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)

		assert.NotNil(t, tmp)
		
		assert.Nil(t, err)
		assert.Nil(t, tmp.Users)
	})

	t.Run("Fetch All Success", func(t *testing.T) {
		var model = SuccessMockUserModel{}
		var controller = NewUserController(&model)

		e.GET("/users", controller.GetUsers())

		var req = httptest.NewRequest(http.MethodGet, "/users", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var resData = json.NewDecoder(res.Result().Body)
		result := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)

		assert.NotNil(t, tmp)
		assert.NotNil(t, result)

	})
}

func TestCreateUser(t *testing.T) {
	
	var e = echo.New()
	
	t.Run("No User Created", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)
		
		e.POST("/users", controller.CreateUser())

		var req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"test": "test"}`)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type Response struct {
			Status int `json:"status"`
			Message string `json:"message"`
		}

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Something went Wrong!", tmp.Message)

		assert.NotNil(t, tmp)
		
		assert.Nil(t, err)
	})

	t.Run("Success User Created", func(t *testing.T) {
		var reqData string = `{"username": "sarbinus", "email": "sarbin@example.com", "password": "siSarbin123"}`

		var model = SuccessMockUserModel{}
		var controller = NewUserController(&model)

		e.POST("/users", controller.CreateUser())

		var req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(reqData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type Response struct {
			Status int `json:"status"`
			Message string `json:"message"`
			User map[string]any `json:"user"`
		}

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "User Created!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.NotNil(t, tmp.User)

		assert.Nil(t, err)
	})
}

func TestGetUser(t *testing.T) {
	var e = echo.New()

	type FailedResponse struct {
		Status int `json:"status"`
		Message string `json:"message"`
	}

	t.Run("Bad Request", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.GET("/users/:id", controller.GetUser())

		var req = httptest.NewRequest(http.MethodGet, "/users/satu", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = FailedResponse{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, http.StatusBadRequest, tmp.Status)
		assert.Equal(t, "Bad Request!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("User Not Found", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.GET("/users/:id", controller.GetUser())

		var req = httptest.NewRequest(http.MethodGet, "/users/999", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = FailedResponse{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, http.StatusNotFound, tmp.Status)
		assert.Equal(t, "Not Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Success Get User", func(t *testing.T) {
		var model = SuccessMockUserModel{}
		var controller = NewUserController(&model)

		e.GET("/users/:id", controller.GetUser())

		var req = httptest.NewRequest(http.MethodGet, "/users/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type SuccessResponse struct {
			Status int `json:"status"`
			Message string `json:"message"`
			User map[string]any `json:"user"`
		}

		var tmp = SuccessResponse{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "User Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.NotNil(t, tmp.User)
		assert.Nil(t, err)
	})
}

func TestEditUser(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Status int `json:"status"`
		Message string `json:"message"`
	}

	var reqData = `{"title": "User 1", "author": "Sarbin", "publisher": "PT. Gremadia"}`

	t.Run("Bad Request", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.PUT("/users/:id", controller.EditUser())

		var req = httptest.NewRequest(http.MethodPut, "/users/satu", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, http.StatusBadRequest, tmp.Status)
		assert.Equal(t, "Bad Request!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("User Not Found", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.PUT("/users/:id", controller.EditUser())

		var req = httptest.NewRequest(http.MethodPut, "/users/999", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, http.StatusNotFound, tmp.Status)
		assert.Equal(t, "Not Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Update Failed", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.PUT("/users/:id", controller.EditUser())

		var req = httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, http.StatusInternalServerError, tmp.Status)
		assert.Equal(t, "Something Went Wrong!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Update Success", func(t *testing.T) {
		var model = SuccessMockUserModel{}
		var controller = NewUserController(&model)

		e.PUT("/users/:id", controller.EditUser())

		var req = httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewReader([]byte(reqData)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "User Updated!", tmp.Message)

		assert.Nil(t, err)
	})
}

func TestRemoveUser(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Status int `json:"status"`
		Message string `json:"message"`
	}

	t.Run("Bad Request", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.DELETE("/users/:id", controller.RemoveUser())

		var req = httptest.NewRequest(http.MethodDelete, "/users/satu",nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, http.StatusBadRequest, tmp.Status)
		assert.Equal(t, "Bad Request!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("User Not Found", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.DELETE("/users/:id", controller.RemoveUser())

		var req = httptest.NewRequest(http.MethodDelete, "/users/999",nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, http.StatusNotFound, tmp.Status)
		assert.Equal(t, "Not Found!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Delete Failed", func(t *testing.T) {
		var model = MockUserModel{}
		var controller = NewUserController(&model)

		e.DELETE("/users/:id", controller.RemoveUser())

		var req = httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, http.StatusInternalServerError, tmp.Status)
		assert.Equal(t, "Something Went Wrong!", tmp.Message)

		assert.NotNil(t, tmp)
		assert.Nil(t, err)
	})

	t.Run("Delete Success", func(t *testing.T) {
		var model = SuccessMockUserModel{}
		var controller = NewUserController(&model)

		e.DELETE("/users/:id", controller.RemoveUser())

		var req = httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = Response{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, http.StatusOK, tmp.Status)
		assert.Equal(t, "User Deleted!", tmp.Message)

		assert.Nil(t, err)
	})
}