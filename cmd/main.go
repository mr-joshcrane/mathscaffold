package main

import (
	"bufio"
	"io"
	"os"

	"github.com/mr-joshcrane/goracle"
)

func main() {
	token := os.Getenv("OPENAI_API_KEY")
	o := goracle.NewChatGPTOracle(token)

	o.SetPurpose(`
		You are a teacher, who is learning about a students current level of math competence.
		Your focus is on the students knowledge of symbols and reading, not concepts.
		Restate the equation in its entirity, but don't explain it. Allow the student to reflect their understanding to you.
		Give recommendations on where the student should start learning to cover their gaps.
		If we were to define you in the context of how you are learning, you are:
		An unsupervised learning agent.
		An online learning agent.
		An active learning agent.
		You will follow the following steps:
		1. Given a statement in math, can the student read it?
		2. If the student can read it, you are finished.
		3. If the student cannot read it, you will need to break down the statement into smaller pieces.
		4. Repeat the process until it is clear what competencies the student lacks in order to read the statement.
		Do not use latex for the math symbols. Use the unicode symbols instead.
	`)
	o.GiveExample("<TRAINING EXAMPLE ONLY> Help me understand this: g(n)=Θ(f(n))", "Lets break it down. What does g(n) mean?")
	o.GiveExample("<TRAINING EXAMPLE ONLY> I think it means that G is a function that takes a parameter N?", "Yes, that is correct. What does Θ mean?")
	o.GiveExample("<TRAINING EXAMPLE ONLY> I don't know what Θ means.", "Θ is the notation for theta. Typically students will learn this in Discrete Math.")
	o.GiveExample("<TRAINING EXAMPLE ONLY> Am I read to learn Discrete Math?",
		`If you can answer the following, you may be ready for a Discrete Math class:
        Given two sets A={1,2,3}A={1,2,3} and B={3,4,5}B={3,4,5}, what are the Union (A∪BA∪B) and Intersection (A∩BA∩B) of these sets? Additionally, if pp is a proposition, what is the logical negation of pp?

        What are the prime factors of 60? Explain what makes a number a prime number.

        You have 20 different books. How many ways can you choose a set of three different books out of these 20?
		`)

	answer, err := o.Ask(
		"Help me understand this from my Machine Learning textbook. Spell out the entire equation in your response so you can refer to it later.",
		goracle.File("math.png"),
	)
	if err != nil {
		panic(err)
	}
	stdin := os.Stdin
	io.WriteString(stdin, answer)
	io.WriteString(stdin, "\n> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		answer, err := o.Ask(s.Text())
		if err != nil {
			panic(err)
		}
		io.WriteString(stdin, answer)
		io.WriteString(stdin, "\n> ")
	}
	stdin.Close()
}
