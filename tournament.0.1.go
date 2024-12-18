/*
Significantly cleaned up draft. Next up, ask for help ensuring goal,
a readable clearly written bit of code
*/
package tournament
import (
    "errors"
    "fmt"
    "io"
    "sort"
    "strings"
)
type TeamStats struct {
    MP int // Matches Played
    W  int // Wins
    D  int // Draws
    L  int // Losses
    P  int // Points
}
func Tally(reader io.Reader, writer io.Writer) error {
    fmt.Printf("reader: %v", reader)
    // Read all data from the reader
    sliceOfBytes, err := io.ReadAll(reader)
    if err == nil {
        // Convert byte slice to string //
        data := string(sliceOfBytes)
        // Split data into lines
        lines := strings.Split(data, "\n")
        // Map to store team stats
        teamStats := mapTeamStats(lines)
        if _, exists := teamStats["notWinLoseDraw"]; exists {
            err = errors.New("notWinLoseDraw")
        } 
        if _, exists := teamStats["notEnoughData"]; exists {
            err = errors.New("notEnoughData")
        } 
        sortedTeams := sortTeamStats(teamStats)
        // Write header to output
        outputToWriter(sortedTeams, writer, teamStats)
    }
    return err
}
func sortTeamStats(teamStats map[string]*TeamStats) (teams []string) { 
    // Create a slice of teams for sorting
    teams = make([]string, 0, len(teamStats))
    for team := range teamStats {
        teams = append(teams, team)
    }
    // Sort teams by points (descending), then by name (ascending)
    sort.Slice(teams, func(i, j int) bool {
        if teamStats[teams[i]].P != teamStats[teams[j]].P {
            return teamStats[teams[i]].P > teamStats[teams[j]].P
        }
        return teams[i] < teams[j]
    })
    return teams
}

func mapTeamStats(lines []string) map[string]*TeamStats {
    teamStats := map[string]*TeamStats{}
    for _, line := range lines {
        // Skip empty lines
        if strings.TrimSpace(line) == "" {
            continue
        }
        // Skip comment lines
        if line[0] == '#' {
            continue
        }
        // Split each line by ';'
        parts := strings.Split(line, ";")
        if len(parts) != 3 {
             teamStats["notEnoughData"] = &TeamStats{}
             teamStats["notEnoughData"].MP = 1
             return teamStats
        }
        team1, team2, result := parts[0], parts[1], parts[2]
        // Initialize team stats if not already present
        if _, exists := teamStats[team1]; !exists {
            teamStats[team1] = &TeamStats{}
        }
        if _, exists := teamStats[team2]; !exists {
            teamStats[team2] = &TeamStats{}
        }
        // Update stats based on match result
        teamStats[team1].MP++
        teamStats[team2].MP++
        switch result {
        case "win":
            teamStats[team1].W++
            teamStats[team2].L++
            teamStats[team1].P += 3
        case "loss":
            teamStats[team1].L++
            teamStats[team2].W++
            teamStats[team2].P += 3
        case "draw":
            teamStats[team1].D++
            teamStats[team2].D++
            teamStats[team1].P++
            teamStats[team2].P++
        default:
             teamStats["notWinLoseDraw"] = &TeamStats{}
             teamStats["notWinLoseDraw"].MP ++
             return teamStats
        }
    }
    return teamStats
}
func outputToWriter(teams []string, writer io.Writer, teamStats map[string]*TeamStats) {
        // Write header to output
    _, _ = fmt.Fprintln(writer, "Team                           | MP |  W |  D |  L |  P")
    // Write each team's stats
    for _, team := range teams {
        stats := teamStats[team]
        _, _ = fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
            team, stats.MP, stats.W, stats.D, stats.L, stats.P)
    }
}
