package main

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
)

func CreateNewKeyCloakUser() error {
	client := gocloak.NewClient("http://host.docker.internal:8888")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "admin", "admin", "master")
	if err != nil {
		panic(err)
	}

	user := gocloak.User{
		FirstName: gocloak.StringP("Bob"),
		LastName:  gocloak.StringP("Uncle"),
		Email:     gocloak.StringP("something@really.wrong"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP("CoolGuy"),
	}

	_, err = client.CreateUser(ctx, token.AccessToken, "master", user)
	if err != nil {
		panic("Oh no!, failed to create user :(")
	}
	return nil
}
