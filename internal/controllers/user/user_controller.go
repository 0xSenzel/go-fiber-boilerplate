package user

import "github.com/0xsenzel/go-fiber-boilerplate/internal/services/user"

func HelloWorldHandler() string {
	return user.GetHellowWorld()
}