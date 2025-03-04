/*
you have been asked to manage a basketball league and are going to write a program to help you.

	Define two types.
	 The first one, called Team, has a field for the name of the team and a field for the player names.
	 The second type is called League and has a field called Teams for the teams in the league
	  and a field called Wins that maps a team’s name to its number of wins.
*/
package ch8

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	teamName    string
	playerNames []string
}

type League struct {
	Name  string
	teams map[string]Team
	wins  map[string]int
}

/*
Add two methods to League.
The first method is called MatchResult.
It takes four parameters: the name of the first team, its score in the game, the name of the second team, and its score in the game.
 This method should update the Wins field in League.
  Add a second method to League called Ranking that returns a slice of the team names in order of wins.
Build your data structures and call these methods from the main function in your program using some sample data.
*/

func (l League) MatchResult(t1 string, score1 int, t2 string, score2 int) {
	if _, ok := l.teams[t1]; !ok {
		return
	}
	if _, ok := l.teams[t2]; !ok {
		return
	}
	if score1 == score2 {
		return
	}
	if score1 > score2 {

		l.wins[t1]++
	} else {
		l.wins[t2]++
	}
}
func (l League) Ranking() []string {
	names := make([]string, 0, len(l.teams))
	for k := range l.teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.wins[names[i]] > l.wins[names[j]]
	})
	return names
}

func Master() {
	l := League{
		Name: "ipl",
		teams: map[string]Team{
			"rcb": {
				"rcb", []string{"kohli", "patidar"},
			},
			"mi": {
				"mi", []string{"rohith", "pandya"},
			},
			"csk": {
				"csk", []string{"dhoni", "jadeja"},
			},
			"srh": {
				"srh", []string{"abhisekh", "klassen"},
			},
		},
		wins: map[string]int{},
	}

	l.MatchResult("rcb", 180, "mi", 80)
	l.MatchResult("rcb", 182, "mi", 190)
	l.MatchResult("rcb", 49, "mi", 80)
	l.MatchResult("csk", 49, "srh", 80)
	l.MatchResult("rcb", 49, "csk", 80)
	l.MatchResult("rcb", 19, "srh", 80)
	l.MatchResult("srh", 149, "mi", 80)

	result := l.Ranking()
	fmt.Println(result)

	RankPrinter(l, os.Stdout)

}

// Define an interface called Ranker that has a single method called Ranking that returns a slice of strings.
// Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer.
// Use the io.WriteString function to write the values returned by Ranker to the io.Writer, with a newline separating each result.
//  Call this function from main.

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}
