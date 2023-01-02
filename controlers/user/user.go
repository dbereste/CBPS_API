package user

import (
	user "main/models/user"
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := user.GetAll()

		c.JSON(http.StatusOK, gin.H{
			"Status":  "OK",
			"content": data,
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string = c.Param("id")

		data := user.GetByID(id)
		if data != (user.User{}) {
			c.JSON(http.StatusOK, gin.H{
				"Status":  "OK",
				"content": data,
			})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userModel user.User

		c.BindJSON(&userModel)

		//TODO: Add proper validation
		if userModel.Login == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Login Provided",
			})
			return
		}

		if userModel.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Name Provided",
			})
			return
		}

		if userModel.Surname == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Surname Provided",
			})
			return
		}

		if userModel.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Password Provided",
			})
			return
		}

		_, err := mail.ParseAddress(userModel.Email)
		if userModel.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Email Provided",
			})
			return
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Email Provided",
			})
			return
		}

		if userModel.Business_name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Email Provided",
			})
			return
		}

		if userModel.Address == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Address Provided",
			})
			return
		}

		if userModel.Phone_number == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Phone_number Provided",
			})
			return
		}

		response := user.New(userModel)

		if response != 0 {
			c.JSON(http.StatusCreated, gin.H{
				"Status":     "OK",
				"InsertedID": response,
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"Status": "NOK",
				"Error":  "Unknown Error. Not Modified",
			})
		}
	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userModel user.User

		c.BindJSON(&userModel)
		var id string = c.Param("id")
		idint64, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  err,
			})
			return
		}

		idint := int(idint64)

		//TODO: Add validator separately
		if idint <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Invalid ID provided",
			})
			return
		}

		usertest := user.GetByID(id)

		if usertest == (user.User{}) {
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}

		if userModel.Login == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Login Provided",
			})
			return
		}

		if userModel.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Name Provided",
			})
			return
		}

		if userModel.Surname == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Surname Provided",
			})
			return
		}

		if userModel.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "No Password Provided",
			})
			return
		}

		req, err := mail.ParseAddress(userModel.Email)
		_ = req

		if userModel.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Email Provided",
			})
			return
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Email Provided",
			})
			return
		}

		if userModel.Business_name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Email Provided",
			})
			return
		}

		if userModel.Address == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Address Provided",
			})
			return
		}

		timeStamp, err := time.Parse("2006-01-02 15:04:05", userModel.Lastlogindate)
		_ = timeStamp
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Login DateTime Provided",
			})
			return
		}

		if userModel.Phone_number == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "NOK",
				"Error":  "Bad Phone_number Provided",
			})
			return
		}

		if userModel.Email != "" {
			usertest.Email = userModel.Email
		}

		if userModel.Password != "" {
			usertest.Password = userModel.Password
		}

		if userModel.Name != "" {
			usertest.Name = userModel.Name
		}

		if userModel.Surname != "" {
			usertest.Surname = userModel.Surname
		}

		if userModel.Login != "" {
			usertest.Login = userModel.Login
		}

		if userModel.Lastlogindate != "" {
			usertest.Lastlogindate = userModel.Lastlogindate
		}

		if userModel.Business_name != "" {
			usertest.Business_name = userModel.Business_name
		}

		if userModel.Address != "" {
			usertest.Address = userModel.Address
		}

		if userModel.Phone_number <= 0 {
			usertest.Phone_number = userModel.Phone_number
		}

		errr := user.Update(usertest)

		if errr == nil {
			c.JSON(http.StatusCreated, gin.H{
				"Status": "OK",
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"Status": "NOK",
				"Error":  errr,
			})
		}
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string = c.Param("id")

		usertest := user.GetByID(id)

		if usertest == (user.User{}) {
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
		data := user.Delete(id)
		if data == 0 {
			c.JSON(http.StatusOK, gin.H{
				"Result": "OK",
			})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
