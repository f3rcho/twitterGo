package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func Register(ctx context.Context) models.ResposeAPI {
	var u models.User
	var r models.ResposeAPI

	fmt.Println("Register...")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(u.Email) == 0 {
		r.Message = "Email is required"
		fmt.Println(r.Message)
		return r
	}

	if len(u.Password) < 6 {
		r.Message = "Password must be at least 6 characters"
		fmt.Println(r.Message)
		return r
	}

	_, userFound, _ := db.IsUserExists(u.Email)
	if userFound {
		r.Message = "User already exists"
		fmt.Println(r.Message)
		return r
	}
	_, status, err := db.CreateUser(u)
	if err != nil {
		r.Message = "An error occurred creating user" + err.Error()
	}

	if !status {
		r.Message = "An error occurred creating user"
		fmt.Println(r.Message)
		return r
	}
	r.Status = 200
	r.Message = "User created successfully"
	fmt.Println(r.Message)
	return r
}
