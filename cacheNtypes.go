package main

type user struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type article struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Time    string    `json:"time"`
	Name    string    `json:"name"`
	Comment []comment `json:"comment"`
}

type comment struct {
	ComTime    string `json:"comtime"`
	ComContent string `json:"comcontent"`
	ComName    string `json:"comname"`
}

// Tutorials - cache of tutor articles
var Tutorials = []article{
	article{
		Title: "How to start?",
		Content: `At first you need to download the game (yep, download from Play Market f.e.). 
		If you did your best in this, you can play, but at first - sign up/in. You can make an account in this forum or in game. 
		Please, use normal username (guys such as "PussyDestroyer69" or "PososiMorgenstern228" will be permanently banned). 
		Enter your real email adress, because you will get validation email. 
		Ok, now you have game on phone and account. That's cool!`,
	},
	article{
		Title: "Plots",
		Content: `You can see many plots on territory. 
		There two types of plots: empty and active.`,
	},
	article{
		Title: "Empty plot",
		Content: `If you press on empty plot, you'll see, that you can place some objects on it. 
		User-placed objects can be modified and deleted.`,
	},
	article{
		Title: "Active plot",
		Content: `If you press on active plot, you'll see, that it already has object on it. 
		Sometimes you can delete or modify them.`,
	},
	article{
		Title: "Objects",
		Content: `Active plots use objects to get money. 
		Better objects on better plots - give you more money. 
		Some objects such as "Candle" - you can't delete or modify.`,
	},
}

// FAQs - cache of frequently asked questions
var FAQs = []article{
	article{
		Title: "What is NULESandbox?",
		Content: `NULESandbox - it's a game, where you can change territory of our beautiful university! 
		Don't like Student's square? 
		No problem! Just place McDonald's instead. 
		Third corpus looks ugly? 
		Ok, paint it in PINK! 
		Cool? Yeah!`,
	},
	article{
		Title: "Who is creator?",
		Content: `I'm Alexander Yakushev: human, person, student. 
		Created this forum for courses, like a final exam. 
		Made the game as my diploma in NULES. `,
	},
	article{
		Title: "How to start new game?",
		Content: `At main menu you can see many colorful buttons. 
		"New game" - starts new game.`,
	},
	article{
		Title:   "How can I see the creator without leaving the game?",
		Content: `At main menu press "Credits" button.`,
	},
	article{
		Title:   "What's better: Slipknot or Pharaon?",
		Content: `They create songs in different styles, but I like Slipknot much more.`,
	},
}
