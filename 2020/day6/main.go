package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	groups := parse(string(content))

	answeredQuestsionsCount := 0
	agreedQuestionsCount := 0
	for _, g := range groups {
		answeredQuestsionsCount += g.AnswersCount()
		agreedQuestionsCount += g.AgreedCount()
	}

	fmt.Println("Sum of answered questions per group member:", answeredQuestsionsCount)
	fmt.Println("Agreed questions:", agreedQuestionsCount)
}

func parse(content string) (groups []*Group) {
	groupBlocks := strings.Split(string(content), "\n\n")

	for _, g := range groupBlocks {
		lines := strings.Split(g, "\n")
		var members []*Member
		for _, l := range lines {
			runes := make(map[rune]bool)
			for _, r := range l {
				runes[r] = true
			}
			members = append(members, &Member{runes})
		}
		groups = append(groups, &Group{members})
	}
	return
}

type Group struct {
	members []*Member
}

func (g *Group) Questions() (questions []rune) {
	for _, m := range g.members {
		for q := range m.answers {
			questions = append(questions, q)
		}
	}
	return questions
}

func (g *Group) Answers() map[rune]bool {
	answers := make(map[rune]bool)
	for _, m := range g.members {
		for q, a := range m.GetAnswers(g.Questions()) {
			agreement := a
			if v, ok := answers[q]; ok {
				agreement = v && a
			}
			answers[q] = agreement
		}
	}
	return answers
}

func (g *Group) AnswersCount() int {
	return len(g.Answers())
}

func (g *Group) AgreedCount() int {
	agreed := 0
	for _, a := range g.Answers() {
		if a {
			agreed++
		}
	}
	return agreed
}

type Member struct {
	answers map[rune]bool
}

func (m *Member) GetAnswers(questions []rune) (answers map[rune]bool) {
	answers = m.answers
	for _, r := range questions {
		if _, ok := m.answers[r]; !ok {
			answers[r] = false
		}
	}
	return answers
}
