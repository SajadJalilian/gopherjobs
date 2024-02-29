package main


// In production, the secret key of the CookieStore
// and database name would be obtained from a .env file
const (
	SECRET_KEY string = "secret"
	DB_NAME    string = "app_data.db"
)
func main() {
		e := echo.New()

	e.Static("/", "assets")

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// Helpers Middleware
	// e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Session Middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

	store, err := db.NewStore(DB_NAME)
	if err != nil {
		e.Logger.Fatalf("failed to create store: %s", err)
	}

	// Setting Routes
	handlers.SetupRoutes(e)

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}