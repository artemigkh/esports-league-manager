package databaseAccess

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

// TeamCore
func (team *TeamCore) validate(leagueId, teamId int) (bool, string, error) {
	return validate(
		team.name(),
		team.uniqueness(leagueId, teamId),
		team.tag())
}

func (team *TeamCore) ValidateNew(leagueId int) (bool, string, error) {
	return team.validate(leagueId, 0)
}

func (team *TeamCore) ValidateEdit(leagueId, teamId int) (bool, string, error) {
	return team.validate(leagueId, teamId)
}

func (team *TeamCore) name() ValidateFunc {
	return func(problemDest *string, _ *error) bool {
		valid := false
		if len(team.Name) > MaxNameLength {
			*problemDest = NameTooLong
		} else if len(team.Name) < MinInformationLength {
			*problemDest = NameTooShort
		} else {
			valid = true
		}
		return valid
	}
}

func (team *TeamCore) uniqueness(leagueId, teamId int) ValidateFunc {
	return func(problemDest *string, errorDest *error) bool {
		valid := false
		inUse, problem, err := teamsDAO.IsInfoInUse(leagueId, teamId, team.Name, team.Tag)
		if err != nil {
			*errorDest = err
		} else if inUse {
			*problemDest = problem
		} else {
			valid = true
		}
		return valid
	}
}

func (team *TeamCore) tag() ValidateFunc {
	return func(problemDest *string, _ *error) bool {
		valid := false
		if len(team.Tag) > MaxTagLength {
			*problemDest = TagTooLong
		} else if len(team.Tag) < MinInformationLength {
			*problemDest = TagTooShort
		} else {
			valid = true
		}
		return valid
	}
}

// PlayerCore
func (player *PlayerCore) validate(leagueId, teamId, playerId int) (bool, string, error) {
	return validate(
		player.name(),
		player.uniqueness(leagueId, teamId, playerId))
}

func (player *PlayerCore) ValidateNew(leagueId, teamId int) (bool, string, error) {
	return player.validate(leagueId, teamId, 0)
}

func (player *PlayerCore) ValidateEdit(leagueId, teamId, playerId int) (bool, string, error) {
	return player.validate(leagueId, teamId, playerId)
}

func (player *PlayerCore) name() ValidateFunc {
	return func(problemDest *string, _ *error) bool {
		valid := false
		if len(player.Name) > MaxNameLength {
			*problemDest = NameTooLong
		} else if len(player.Name) < MinInformationLength {
			*problemDest = NameTooShort
		} else {
			valid = true
		}
		return valid
	}
}

func (player *PlayerCore) uniqueness(leagueId, teamId, playerId int) ValidateFunc {
	return func(problemDest *string, _ *error) bool {
		//TODO: implement this
		return true
	}
}

// TeamPermissionsCore
type TeamPermissionsCoreArray struct {
	rows []*TeamPermissionsCore
}

func GetScannedTeamPermissionsCore(rows squirrel.RowScanner) (*TeamPermissionsCore, error) {
	var teamPermissions TeamPermissionsCore
	if err := rows.Scan(
		&teamPermissions.Administrator,
		&teamPermissions.Information,
		&teamPermissions.Games,
	); err != nil {
		return nil, err
	} else {
		return &teamPermissions, nil
	}
}

func (r *TeamPermissionsCoreArray) Scan(rows *sql.Rows) error {
	row, err := GetScannedTeamPermissionsCore(rows)
	if err != nil {
		return err
	} else {
		r.rows = append(r.rows, row)
		return nil
	}
}

func (p *TeamPermissionsCore) Validate() (bool, string, error) {
	return validate(p.consistent())
}

func (p *TeamPermissionsCore) consistent() ValidateFunc {
	return func(problemDest *string, _ *error) bool {
		if (p.Information || p.Games) && p.Administrator {
			*problemDest = AdminLackingPermissions
			return false
		} else {
			return true
		}
	}
}

// TeamWithPlayers
type TeamWithPlayersArray struct {
	rows []*TeamWithPlayers
}

func getTeamWithPlayersSelector() squirrel.SelectBuilder {
	return psql.Select(
		"team.team_id",
		"team.name",
		"team.description",
		"team.tag",
		"team.icon_small",
		"team.icon_large",
		"team.wins",
		"team.losses",
		"player.player_id",
		"player.name",
		"player.game_identifier",
		"player.main_roster",
	).
		From("team").
		LeftJoin("player ON team.team_id = player.team_id")
}

func GetScannedTeamWithPlayers(rows *sql.Rows) (*TeamWithPlayers, error) {
	defer rows.Close()

	var team TeamWithPlayers

	for rows.Next() {
		var (
			playerId             sql.NullInt64
			playerName           sql.NullString
			playerGameIdentifier sql.NullString
			playerMainRoster     sql.NullBool
		)
		if err := rows.Scan(
			&team.TeamId,
			&team.Name,
			&team.Description,
			&team.Tag,
			&team.IconSmall,
			&team.IconLarge,
			&team.Wins,
			&team.Losses,
			&playerId,
			&playerName,
			&playerGameIdentifier,
			&playerMainRoster,
		); err != nil {
			return nil, err
		}
		if playerId.Valid {
			team.Players = append(team.Players, &Player{
				PlayerId:       int(playerId.Int64),
				Name:           playerName.String,
				GameIdentifier: playerGameIdentifier.String,
				MainRoster:     playerMainRoster.Bool,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &team, nil
}

func GetScannedAllTeamWithPlayers(rows *sql.Rows) ([]*TeamWithPlayers, error) {
	var teams []*TeamWithPlayers
	getUniqueTeam := func(newTeam *TeamWithPlayers) *TeamWithPlayers {
		for _, team := range teams {
			if newTeam.TeamId == team.TeamId {
				return team
			}
		}
		teams = append(teams, newTeam)
		return newTeam
	}

	defer rows.Close()
	for rows.Next() {
		var team TeamWithPlayers
		var (
			playerId             sql.NullInt64
			playerName           sql.NullString
			playerGameIdentifier sql.NullString
			playerMainRoster     sql.NullBool
		)
		if err := rows.Scan(
			&team.TeamId,
			&team.Name,
			&team.Description,
			&team.Tag,
			&team.IconSmall,
			&team.IconLarge,
			&team.Wins,
			&team.Losses,
			&playerId,
			&playerName,
			&playerGameIdentifier,
			&playerMainRoster,
		); err != nil {
			return nil, err
		}

		uniqueTeam := getUniqueTeam(&team)
		if playerId.Valid {
			uniqueTeam.Players = append(uniqueTeam.Players, &Player{
				PlayerId:       int(playerId.Int64),
				Name:           playerName.String,
				GameIdentifier: playerGameIdentifier.String,
				MainRoster:     playerMainRoster.Bool,
			}) // TODO: refactor this to not have dupe code with above
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *TeamWithPlayersArray) Scan(rows *sql.Rows) error {
	row, err := GetScannedTeamWithPlayers(rows)
	if err != nil {
		return err
	} else {
		r.rows = append(r.rows, row)
		return nil
	}
}

// TeamDisplay
type TeamDisplayArray struct {
	rows []*TeamDisplay
}

func getTeamDisplaySelector() squirrel.SelectBuilder {
	return psql.Select(
		"team_id",
		"name",
		"tag",
		"icon_small",
	).From("team")
}

func GetScannedTeamDisplay(rows squirrel.RowScanner) (*TeamDisplay, error) {
	var team TeamDisplay
	if err := rows.Scan(
		&team.TeamId,
		&team.Name,
		&team.Tag,
		&team.IconSmall,
	); err != nil {
		return nil, err
	} else {
		return &team, nil
	}
}

func (r *TeamDisplayArray) Scan(rows *sql.Rows) error {
	row, err := GetScannedTeamDisplay(rows)
	if err != nil {
		return err
	} else {
		r.rows = append(r.rows, row)
		return nil
	}
}

// TeamPermissions
func getTeamPermissionsSelector() squirrel.SelectBuilder {
	return psql.Select(
		"team.team_id",
		"team.name",
		"team.tag",
		"team.icon_small",
		"team_permissions.administrator",
		"team_permissions.information",
		"team_permissions.games",
	).
		From("team").
		Join("team_permissions ON team.team_id = team_permissions.team_id")
}

type TeamPermissionsArray struct {
	rows []*TeamPermissions
}

func GetScannedTeamPermissions(rows squirrel.RowScanner) (*TeamPermissions, error) {
	var teamPermissions TeamPermissions
	if err := rows.Scan(
		&teamPermissions.TeamId,
		&teamPermissions.Name,
		&teamPermissions.Tag,
		&teamPermissions.IconSmall,
		&teamPermissions.Administrator,
		&teamPermissions.Information,
		&teamPermissions.Games,
	); err != nil {
		return nil, err
	} else {
		return &teamPermissions, nil
	}
}

func (r *TeamPermissionsArray) Scan(rows *sql.Rows) error {
	row, err := GetScannedTeamPermissions(rows)
	if err != nil {
		return err
	} else {
		r.rows = append(r.rows, row)
		return nil
	}
}