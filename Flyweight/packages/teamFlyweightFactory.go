package packages

// Team_A,Team_B Const
const (
	Team_A = iota
	Team_B
)

// TeamFlyweightFactory struct
type TeamFlyweightFactory struct {
	CreatedTeams map[int]*Team
}

// GetTeamFactory func
func GetTeamFactory(team int) Team {
	switch team {
	case Team_A:
		return Team{
			ID:   2,
			Name: "Team_A",
		}
	default:
		return Team{
			ID:   3,
			Name: "Team_B",
		}
	}
}

// NewTeamFactory func
func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		CreatedTeams: make(map[int]*Team, 0),
	}
}

// GetTeam method
func (t *TeamFlyweightFactory) GetTeam(teamName int) *Team {
	if t.CreatedTeams[teamName] != nil {
		return t.CreatedTeams[teamName]
	}

	team := GetTeamFactory(teamName)
	t.CreatedTeams[teamName] = &team

	return t.CreatedTeams[teamName]
}

// GetNumberOfObject method
func (t *TeamFlyweightFactory) GetNumberOfObject() int {
	return len(t.CreatedTeams)
}
