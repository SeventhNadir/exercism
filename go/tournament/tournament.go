package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name                     string
	played, won, drawn, lost int
}

type allTeams map[string]team

func (t *team) points() int {
	return 3*t.won + 1*t.drawn
}

//Tally takes a list of records and converts them into a ranked table.
func Tally(results io.Reader, output io.Writer) error {
	teams := allTeams{}
	scanner := bufio.NewScanner(results)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")

		// Ignore blank lines and comment lines starting with "#"
		if line[0] == "" || strings.HasPrefix(line[0], "#") {
			continue
		}
		if len(line) != 3 {
			return fmt.Errorf("Illegal input %s. Need two teams, one result, semicolon seperators", line)
		}

		team1, team2 := teams[line[0]], teams[line[1]]
		team1.name, team2.name = line[0], line[1]

		switch line[2] {
		case "win":
			team1.won++
			team2.lost++
		case "loss":
			team1.lost++
			team2.won++
		case "draw":
			team1.drawn++
			team2.drawn++
		default:
			return fmt.Errorf("Invalid result: %s", line[2])
		}
		team1.played++
		team2.played++
		teams[team1.name], teams[team2.name] = team1, team2
	}
	sortedTeams := []team{}
	for _, value := range teams {
		sortedTeams = append(sortedTeams, value)
	}

	sort.Slice(sortedTeams, func(i, j int) bool {
		if sortedTeams[i].points() == sortedTeams[j].points() {
			//Teams with equal points() are sorted alphabetically
			return sortedTeams[i].name < sortedTeams[j].name
		}
		return sortedTeams[i].points() > sortedTeams[j].points()
	})

	fmt.Fprintf(output, "%-30s | %2s | %2s | %2s | %2s | %2s\n",
		"Team", "MP", "W", "D", "L", "P")
	for _, team := range sortedTeams {
		fmt.Fprintf(output, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team.name, team.played, team.won, team.drawn, team.lost, team.points())
	}

	return nil
}
