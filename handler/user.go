package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-medium-ecomerce/model"
	repository "github.com/xvbnm48/go-medium-ecomerce/reponsitory"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
	GetByEmail(*gin.Context)
	GetAllUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	GetProductOrdered(*gin.Context)
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() UserHandler {
	return &userHandler{
		repo: repository.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *userHandler) SignInUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbUser, err := h.repo.GetByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := GenerateToken(dbUser.ID)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Success sign in ", "token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"msg": "Invalid email or password",
	})
	return
}
