package main

import (
	"github.com/gin-gonic/gin"
)

// Shows start page
func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Home Page"}, "index.html")
}

// Shows Tutorial page
func showTutorialPage(c *gin.Context) {
	render(c, gin.H{
		"title":   "Tutorial",
		"payload": Tutorials}, "cached-tut.faq.html")
}

// Shows FAQ page
func showFAQPage(c *gin.Context) {
	render(c, gin.H{
		"title":   "FAQ",
		"payload": FAQs}, "cached-tut.faq.html")
}

// Shows Conversation page (user articles)
func showConversationPage(c *gin.Context) {
	articles := getArticleFromDB()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Conversation",
		"payload": articles}, "conversation.html")
}

// Shows page for creation articles
func showArticleCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

// Shows full Registration page
func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register"}, "register.html")
}
