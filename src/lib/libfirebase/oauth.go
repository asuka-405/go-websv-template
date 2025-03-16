package libfirebase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func Initialize() {

	workdir, _ := os.Getwd()
	credentialsFile := filepath.Join(workdir, "firebase-adminsdk.json")

	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialsFile)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Err intializing firebase app: %v", err)
	}
	FirebaseAuth, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Err initializing firebase auth: %v", err)
	}
	log.Println("Firebase initialization successful!!")

}

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	authHeader := r.Header.Get("Authorization")
	const prefix = "Bearer "
	if len(authHeader) < len(prefix) || !strings.HasPrefix(authHeader, prefix) {
		http.Error(w, fmt.Sprintf("Invalid ID token"), http.StatusUnauthorized)
		return
	}

	// Safely extract the token
	token, err := FirebaseAuth.VerifyIDToken(ctx, authHeader[len(prefix):])
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID token: %v", err), http.StatusUnauthorized)
		return
	}

	userData := map[string]interface{}{
		"uid":   token.UID,
		"email": token.Claims["email"],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}
