# Usage

#### JWT Authentication

1. ###### Create an instance of `AuthService` with your secret key:

   ```go
   authService := libauth.NewAuthService([]byte("your-secret-key"))
   ```

2. ###### Generate a JWT token:

   ```go
   token, err := authService.GenerateJWT("userID", time.Hour)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println("Generated Token:", token)
   ```

3. ###### Validate a JWT token:

   ```go
   userID, err := authService.ValidateJWT(token)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println("Validated User ID:", userID)
   ```

4. ###### Use the JWT middleware in your HTTP server:

   ```go
   jwtMiddleware := libauth.NewJWTMiddleware(authService)
   http.Handle("/api", jwtMiddleware.Middleware(http.HandlerFunc(apiHandler)))
   ```

#### Session-Based Authentication

1. ###### Create an instance of `InMemorySessionStore`:

   ```go
   sessionStore := libauth.NewInMemorySessionStore()
   ```

2. ###### Set a session:

   ```go
   err := sessionStore.Set("sessionID", "userID")
   if err != nil {
       log.Fatal(err)
   }
   ```

3. ###### Get a session:

   ```go
   userID, err := sessionStore.Get("sessionID")
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println("Session User ID:", userID)
   ```

4. ###### Use the session middleware in your HTTP server:

   ```go
   sessionMiddleware := libauth.NewSessionMiddleware(sessionStore)
   http.Handle("/dashboard", sessionMiddleware.Middleware(http.HandlerFunc(dashboardHandler)))
   ```
