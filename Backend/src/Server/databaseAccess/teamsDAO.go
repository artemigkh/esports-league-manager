package databaseAccess

import (
	"database/sql"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

type PgTeamsDAO struct{}

func tryGetUniqueIcon(leagueId int) (string, string, error) {
	// get list of icons used
	rows, err := psql.Select("icon_small").
		From("team").
		Where("league_id = ?", leagueId).
		RunWith(db).Query()

	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	// generate bool who's indices indicate if that number is available
	var availableIcons []bool
	for i := 0; i < 9; i++ {
		availableIcons = append(availableIcons, true)
	}

	// mark numbers as taken if the filename associated with it is present
	var icon string
	for rows.Next() {
		err := rows.Scan(&icon)
		if err != nil {
			return "", "", err
		}
		for i := 0; i < 9; i++ {
			if icon == fmt.Sprintf("generic-%v-small.png", i+1) {
				availableIcons[i] = false
			}
		}
	}
	if rows.Err() != nil {
		return "", "", err
	}

	// create list of available generic icons
	var availableNumbers []int
	for i := 0; i < 9; i++ {
		if availableIcons[i] {
			availableNumbers = append(availableNumbers, i+1)
		}
	}

	// select one either from available or if all taken a random one
	var newIconNumber int
	println(fmt.Sprintf("Available numbers: %v", len(availableNumbers)))
	if len(availableNumbers) == 0 {
		newIconNumber = randomdata.Number(1, 9)
	} else if len(availableNumbers) == 1 {
		newIconNumber = availableNumbers[0]
	} else {
		newIconNumber = availableNumbers[randomdata.Number(0, len(availableNumbers)-1)]
	}

	return fmt.Sprintf("generic-%v-small.png", newIconNumber),
		fmt.Sprintf("generic-%v-large.png", newIconNumber), nil
}

// Teams

func (d *PgTeamsDAO) CreateTeam(leagueId, userId int, teamInfo TeamDTO) (int, error) {
	iconSmall, iconLarge, err := tryGetUniqueIcon(leagueId)
	if err != nil {
		return -1, err
	}

	teamInfo.IconSmall = iconSmall
	teamInfo.IconLarge = iconLarge

	return d.CreateTeamWithIcon(leagueId, userId, teamInfo)
}

func (d *PgTeamsDAO) CreateTeamWithIcon(leagueId, userId int, teamInfo TeamDTO) (int, error) {
	var teamId = -1
	err := db.QueryRow("SELECT create_team($1,$2,$3,$4,$5,$6,$7)",
		leagueId,
		teamInfo.Name,
		teamInfo.Tag,
		teamInfo.Description,
		teamInfo.IconSmall,
		teamInfo.IconLarge,
		userId,
	).Scan(&teamId)

	return teamId, err
}

func (d *PgTeamsDAO) DeleteTeam(leagueId, teamId int) error {
	//remove players from team
	_, err := psql.Delete("player").
		Where("team_id = ?", teamId).
		RunWith(db).Exec()
	if err != nil {
		return err
	}

	//remove team
	_, err = psql.Delete("team").
		Where("id = ? AND league_id = ?", teamId, leagueId).
		RunWith(db).Exec()
	return err
}

func (d *PgTeamsDAO) UpdateTeam(leagueId, teamInformation TeamDTO) error {
	_, err := psql.Update("team").
		Set("name", teamInformation.Name).
		Set("tag", teamInformation.Tag).
		Set("description", teamInformation.Description).
		Where("league_id = ? AND id = ?", leagueId, teamInformation.Id).
		RunWith(db).Exec()

	return err
}

func (d *PgTeamsDAO) UpdateTeamIcon(leagueId, teamId int, small, large string) error {
	_, err := psql.Update("team").
		Set("icon_small", small).
		Set("icon_large", large).
		Where("id = ? AND league_id = ?", teamId, leagueId).
		RunWith(db).Exec()

	return err
}

func (d *PgTeamsDAO) GetTeamInformation(leagueId, teamId int) (*TeamDTO, error) {
	teamInformation, err := GetScannedTeamDTO(psql.Select(
		"id",
		"name",
		"tag",
		"description",
		"wins",
		"losses",
		"icon_small",
		"icon_large").
		From("team").
		Where("id = ? AND league_id = ?", teamId, leagueId).
		RunWith(db).QueryRow())

	if err != nil {
		return nil, err
	}

	//get players of team
	var players PlayerDTOArray
	if err := ScanRows(psql.Select(
		"id",
		"teamId",
		"name",
		"game_identifier",
		"external_id",
		"position",
		"main_roster").
		From("player").
		Where("team_id = ?", teamId), &players); err != nil {
		return nil, err
	}

	teamInformation.Players = players.rows
	return teamInformation, nil
}

// Players

func (d *PgTeamsDAO) AddNewPlayer(playerInfo PlayerDTO) (int, error) {
	var playerId int
	if err := psql.Insert("player").
		Columns(
			"team_id",
			"game_identifier",
			"name",
			"external_id",
			"position",
			"main_roster",
		).
		Values(
			playerInfo.TeamId,
			playerInfo.GameIdentifier,
			playerInfo.Name,
			playerInfo.ExternalId,
			playerInfo.Position,
			playerInfo.MainRoster,
		).
		Suffix("RETURNING \"id\"").
		RunWith(db).QueryRow().Scan(&playerId); err != nil {
		return -1, err
	}

	return playerId, nil
}

func (d *PgTeamsDAO) RemovePlayer(teamId, playerId int) error {
	_, err := psql.Delete("player").
		Where("id = ? AND team_id = ?", playerId, teamId).
		RunWith(db).Exec()
	return err
}

func (d *PgTeamsDAO) UpdatePlayer(playerInfo PlayerDTO) error {
	_, err := psql.Update("player").
		Set("game_identifier", playerInfo.GameIdentifier).
		Set("name", playerInfo.Name).
		Set("external_id", playerInfo.ExternalId).
		Set("position", playerInfo.Position).
		Set("main_roster", playerInfo.MainRoster).
		Where("id = ? AND team_id = ?", playerInfo.Id, playerInfo.TeamId).
		RunWith(db).Exec()

	return err
}

// Get Information For Team and Player Management

func (d *PgTeamsDAO) GetTeamPermissions(teamId, userId int) (*TeamPermissionsDTO, error) {
	teamPermissions, err := GetScannedTeamPermissionsDTO(psql.Select(
		"administrator",
		"information",
		"players",
		"report_results").
		From("team_permissions").
		Where("user_id = ? AND team_id = ?", userId, teamId).
		RunWith(db).QueryRow())

	if err == sql.ErrNoRows {
		return &TeamPermissionsDTO{
			Administrator: false,
			Information:   false,
			Players:       false,
			ReportResults: false,
		}, nil
	} else if err != nil {
		return nil, err
	}

	return teamPermissions, nil
}

func (d *PgTeamsDAO) IsInfoInUse(leagueId, teamId int, name, tag string) (bool, string, error) {
	var nameCount int
	var tagCount int

	if err := psql.Select("count(name)", "count(tag)").
		From("team").
		Where("name = ? AND teamId != ? AND leagueId = ?", name, teamId, leagueId).
		RunWith(db).QueryRow().Scan(&nameCount, &tagCount); err != nil {
		return false, "", err
	} else if nameCount > 0 {
		return true, "nameInUse", nil
	} else if tagCount > 0 {
		return true, "tagInUse", nil
	} else {
		return false, "", nil
	}
}

func (d *PgTeamsDAO) DoesTeamExist(leagueId, teamId int) (bool, error) {
	var count int
	if err := psql.Select("count(*)").
		From("team").
		Where("id = ? AND league_id = ?", teamId, leagueId).
		RunWith(db).QueryRow().Scan(&count); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (d *PgTeamsDAO) DoesPlayerExist(teamId, playerId int) (bool, error) {
	var count int
	if err := psql.Select("count(name)").
		From("player").
		Where("id = ? AND team_id = ?", playerId, teamId).
		RunWith(db).QueryRow().Scan(&count); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (d *PgTeamsDAO) IsTeamActive(leagueId, teamId int) (bool, error) {
	var count int
	if err := psql.Select("count(id)").
		From("game").
		Where("league_id = ? AND ( team1_id = ? OR team2_id = ?)", leagueId, teamId, teamId).
		RunWith(db).QueryRow().Scan(&count); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

// Managers

func (d *PgTeamsDAO) ChangeManagerPermissions(teamId, userId int, teamPermissionInformation TeamPermissionsDTO) error {
	_, err := psql.Update("team_permissions").
		Set("administrator", teamPermissionInformation.Administrator).
		Set("information", teamPermissionInformation.Information).
		Set("players", teamPermissionInformation.Players).
		Set("report_results", teamPermissionInformation.ReportResults).
		Where("team_id = ? AND user_id = ?", teamId, userId).
		RunWith(db).Exec()

	return err
}
