package routes

import (
	"Server/config"
	lolApi "github.com/artemigkh/GoLang-LeagueOfLegendsAPIV4Framework"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type LeagueOfLegendsPlayerInformation struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	GameIdentifier string `json:"gameIdentifier"` // Jersey Number, IGN, etc.
	ExternalId     string `json:"externalId"`
	MainRoster     bool   `json:"mainRoster"`
	Rank           string `json:"rank"`
	Tier           string `json:"tier"`
}

type LeagueOfLegendsTeamInformation struct {
	Name        string                             `json:"name"`
	Tag         string                             `json:"tag"`
	Description string                             `json:"description"`
	Wins        int                                `json:"wins"`
	Losses      int                                `json:"losses"`
	IconSmall   string                             `json:"iconSmall"`
	IconLarge   string                             `json:"iconLarge"`
	Players     []LeagueOfLegendsPlayerInformation `json:"players"`
}

/**
 * @api{POST} /api/league-of-legends/teams/addPlayer Add Player To Team
 * @apiGroup Teams
 * @apiDescription Create a new player and add him to the teams roster
 *
 * @apiParam {int} teamId The unique numerical identifier of the team the player is to be added to
 * @apiParam {string} name The name of the player (can be left blank)
 * @apiParam {string} gameIdentifier The in-game name identifier of the player (jersey number, ign, etc.)
 * @apiParam {bool} mainRoster If true, this player is on the main roster, otherwise is a substitute
 *
 * @apiSuccess {int} id the unique numerical identifier of the created player
 *
 * @apiError notLoggedIn No user is logged in
 * @apiError noActiveLeague There is no active league selected
 * @apiError teamDoesNotExist The specified team does not exist
 * @apiError canNotEditPlayers The currently logged in player does not have permission to edit the players on this team
 * @apiError gameIdentifierTooLong The game identifier exceeds 50 characters
 * @apiError nameTooLong The name exceeds 50 characters
 * @apiError gameIdentifierTooShort The game identifier is smaller than 2 characters
 * @apiError gameIdentifierInUse This game identifier is already in use in this league
 * @apiError lolApiError There was an error retrieving information from the league of legends api
 */
func leagueOfLegendsGetSummonerId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var playerInfo PlayerInformation
		err := ctx.ShouldBindBodyWith(&playerInfo, binding.JSON)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "malformedInput"})
		}
		summonerId, err := lolApi.FromName(playerInfo.GameIdentifier).SummonerId()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "lolApiError"})
		}
		ctx.Set("externalId", summonerId)
		ctx.Next()
	}
}

/**
 * @api{GET} /api/league-of-legends/teams/:id Get Team Information
 * @apiGroup Teams
 * @apiDescription Get information about the team with specified id
 *
 * @apiParam {int} id The unique numerical identifier of the team
 *
 * @apiSuccess {string} name The name of the team
 * @apiSuccess {string} tag The tag of the team
 * @apiSuccess {string} description The team description
 * @apiSuccess {int} wins The number of wins this team has
 * @apiSuccess {int} losses The number of losses this team has
 * @apiSuccess {string} iconSmall The small icon filename
 * @apiSuccess {string} iconLarge The large icon filename
 * @apiSuccess {[]Object} players An array of json objects representing the players on the team
 * @apiSuccess {int} players.id The unique numerical identifier of the player
 * @apiSuccess {string} players.name The name of the player
 * @apiSuccess {string} players.rank The rank of the player
 * @apiSuccess {string} players.tier The tier of the player
 * @apiSuccess {string} players.gameIdentifier The in-game name identifier of the player (jersey number, ign, etc.)
 * @apiSuccess {bool} players.mainRoster If true, this player is on the main roster, otherwise is a substitute
 *
 * @apiError IdMustBeInteger The id in the url must be an integer value
 * @apiError noActiveLeague There is no active league selected
 * @apiError teamDoesNotExist The specified team does not exist
 */
func leagueOfLegendsGetTeamInformation(ctx *gin.Context) {
	if failIfTeamDoesNotExist(ctx, ctx.GetInt("leagueId"), ctx.GetInt("urlId")) {
		return
	}

	teamInfo, err := TeamsDAO.GetTeamInformation(ctx.GetInt("leagueId"), ctx.GetInt("urlId"))
	if checkErr(ctx, err) {
		return
	}

	leagueTeamInfo := LeagueOfLegendsTeamInformation{
		Name:        teamInfo.Name,
		Tag:         teamInfo.Tag,
		Description: teamInfo.Description,
		Wins:        teamInfo.Wins,
		Losses:      teamInfo.Losses,
		IconSmall:   teamInfo.IconSmall,
		IconLarge:   teamInfo.IconLarge,
		Players:     nil,
	}

	for _, player := range teamInfo.Players {
		lolPlayer := LeagueOfLegendsPlayerInformation{
			Id:             player.Id,
			Name:           player.Name,
			GameIdentifier: player.GameIdentifier,
			ExternalId:     player.ExternalId,
			MainRoster:     player.MainRoster,
			Rank:           "",
			Tier:           "",
		}

		if lolPlayer.ExternalId != "" {
			p := lolApi.FromSummonerId(player.ExternalId)
			soloq, err := p.League(lolApi.SummonersRiftSoloQueue)
			flexq, err := p.League(lolApi.SummonersRiftFlexQueue)
			if soloq != nil && err == nil {
				lolPlayer.Rank = soloq.Rank
				lolPlayer.Tier = soloq.Tier
			} else if flexq != nil && err == nil {
				lolPlayer.Rank = flexq.Rank
				lolPlayer.Tier = flexq.Tier
			}
		}

		leagueTeamInfo.Players = append(leagueTeamInfo.Players, lolPlayer)
	}

	ctx.JSON(http.StatusOK, leagueTeamInfo)
}

func RegisterLeagueOfLegendsHandlers(g *gin.RouterGroup, conf config.Config) {
	lolApi.Init("NA1", conf.GetLeagueOfLegendsApiKey())
	g.Use(getActiveLeague())

	g.POST("/teams/addPlayer", authenticate(), leagueOfLegendsGetSummonerId(), addPlayerToTeam)
	g.PUT("/teams/updatePlayer", authenticate(), leagueOfLegendsGetSummonerId(), updatePlayer)
	g.GET("/teams/:id", getUrlId(), leagueOfLegendsGetTeamInformation)
}