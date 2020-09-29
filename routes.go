// routes.go

package main

func initializeRoutes() {

	// Use the setUserStatus and setAdminStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())
	router.Use(setAdminStatus())

	// Handling Index-page
	router.GET("/", showIndexPage)
	router.POST("/", ensureNotLoggedIn(), performLogin)

	// Handling Tutorial-page
	router.GET("/tutorial", showTutorialPage)
	router.GET("/tutorial/:article_id", getTutorialArticle)

	// Handling FAQ-page
	router.GET("/faq", showFAQPage)
	router.GET("/faq/:article_id", getFAQArticle)

	// Handling Conversation-page
	router.GET("/article", showConersationPage)

	// Group of user routs
	userRoutes := router.Group("/u")
	{
		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	// Group of article routs
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", getArticle)

		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)

		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}

	adminRoutes := router.Group("/admin")
	{
		adminRoutes.GET("/panel", ensureAdminned(), showAdminPanelPage)

		adminRoutes.POST("/panel", ensureAdminned(), delThisShit)
	}

}
