package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-medium-ecomerce/model"
	"github.com/xvbnm48/go-medium-ecomerce/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
	GetAllUser(*gin.Context)
	SignInUser(*gin.Context)
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
	hPas, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPas)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *userHandler) GetAllUser(c *gin.Context) {
	fmt.Println(c.Get("userID"))
	user, err := h.repo.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "Invalid user ID",
		})
		return
	}
	user, err := h.repo.GetUser(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *userHandler) SignInUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dbUser, err := h.repo.GetByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": err.Error(),
		})
		return
	}

	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue == true {
		fmt.Println("user before", dbUser.ID)
		token := GenerateToken(dbUser.ID)
		c.JSON(http.StatusOK, gin.H{
			"status": "success login",
			"token":  token,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": "Invalid email or password",
	})
	return
}

func (h *userHandler) AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashPassword(&user.Password)
	user, err := h.repo.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user.ID = uint(intID)
	user, err = h.repo.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var user model.User
	id := ctx.Param("user")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)
	user, err := h.repo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}

func (h *userHandler) GetProductOrdered(ctx *gin.Context) {

	userStr := ctx.Param("user")
	userID, _ := strconv.Atoi(userStr)
	if products, err := h.repo.GetProductOrdered(userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, products)
	}
}
