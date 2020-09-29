package main

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Name    string `json:"name"`
}

// Tutorials - cache of tutor articles
var Tutorials = []article{
	article{
		ID:    1,
		Title: "How to start?",
		Content: `At first you need to download the game (yep, download from Play Market f.e.). 
		If you did your best in this, you can play, but at first - sign up/in. You can make an account in this forum or in game. 
		Please, use normal username (guys such as "PussyDestroyer69" or "PososiMorgenstern228" will be permanently banned). 
		Enter your real email adress, because you will get validation email. 
		Ok, now you have game on phone and account. That's cool!`,
	},
	article{
		ID:    2,
		Title: "What makes this button?",
		Content: `At main menu you can see many colorful buttons. 
		"New game" - starts new game. 
		"Continue" - launchs last game. 
		"Credits" - shows my name, name of creator.`,
	},
}

// FAQs - cache of frequently asked questions
var FAQs = []article{
	article{
		ID:    1,
		Title: "What is NULESandbox?",
		Content: `NULESandbox - it's a game, where you can change territory of our beautiful university! 
		Don't like Student's square? 
		No problem! Just place McDonald's instead. 
		Third corpus looks ugly? 
		Ok, paint it in PINK! 
		Cool? Yeah!`,
	},
	article{
		ID:    2,
		Title: "Who is creator?",
		Content: `I'm Alexander Yakushev: human, person, student. 
		Created this forum for courses, like a final exam. 
		Made the game as my diploma in NULES. `,
	},
	article{
		ID:    3,
		Title: "Can I...?",
		Content: `- Can I print something bad in articles? 
		- No. 
		-----
		- Can I have badass nickname? 
		- No. 
		-----
		- Can I say something bad about someone? 
		- No. 
		-----
		- Can I...? 
		- No. `,
	},
}
