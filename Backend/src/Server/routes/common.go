package routes

import (
	"Server/databaseAccess"
	"Server/icons"
	"Server/lolApi"
	"Server/markdown"
	"Server/sessionManager"
)

// Objects
var UsersDAO databaseAccess.UsersDAO
var LeaguesDAO databaseAccess.LeaguesDAO
var TeamsDAO databaseAccess.TeamsDAO
var GamesDAO databaseAccess.GamesDAO
var InviteCodesDAO databaseAccess.InviteCodesDAO
var LeagueOfLegendsDAO databaseAccess.LeagueOfLegendsDAO

var ElmSessions sessionManager.SessionManager

var IconManager icons.IconManager
var MarkdownManager markdown.MdManager

var LoLApi lolApi.LoLApi

// Structs
type userInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Functions
func getLeagueAndTeamPermissions(leagueId, teamId, userId int) (
	*databaseAccess.LeaguePermissions, *databaseAccess.TeamPermissions, error) {
	leaguePermissions, err := LeaguesDAO.GetLeaguePermissions(leagueId, userId)
	if err != nil {
		return nil, nil, err
	}

	teamPermissions, err := TeamsDAO.GetTeamPermissions(teamId, userId)
	if err != nil {
		return nil, nil, err
	}

	return leaguePermissions, teamPermissions, nil
}
