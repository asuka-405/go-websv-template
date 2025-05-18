package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitFirebaseSDK(secretJsonPath string) (*auth.Client, error) {
	opt := option.WithCredentialsFile(secretJsonPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app.Auth(context.Background())
}
