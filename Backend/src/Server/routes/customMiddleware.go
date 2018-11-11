package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := ElmSessions.AuthenticateAndGetUserId(ctx)
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if userId == -1 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "notLoggedIn"})
		} else {
			ctx.Set("userId", userId)
			ctx.Next()
		}
	}
}

func getActiveLeague() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		leagueId, err := ElmSessions.GetActiveLeague(ctx)
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if leagueId == -1 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "noActiveLeague"})
		} else {
			ctx.Set("leagueId", leagueId)
			ctx.Next()
		}
	}
}

func getUrlId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "IdMustBeInteger"})
		} else {
			ctx.Set("urlId", urlId)
			ctx.Next()
		}
	}
}

//TODO: make general case on failing of lack of league permissions
//TODO: change inserting logic so administrator always has to have true on all perm fields

func failIfNoTeamCreatePermissions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lp, err := LeaguesDAO.GetLeaguePermissions(ctx.GetInt("leagueId"), ctx.GetInt("userId"))
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if !(lp.Administrator || lp.CreateTeams) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "noEditTeamPermissions"})
		} else {
			ctx.Next()
		}
	}
}

func failIfNoEditSchedulePermissions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lp, err := LeaguesDAO.GetLeaguePermissions(ctx.GetInt("leagueId"), ctx.GetInt("userId"))
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if !(lp.Administrator || lp.EditGames) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "noEditSchedulePermissions"})
		} else {
			ctx.Next()
		}
	}
}

func failIfLeagueDoesNotExist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := LeaguesDAO.GetLeagueInformation(ctx.GetInt("urlId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "leagueDoesNotExist"})
		} else {
			ctx.Next()
		}
	}
}

func failIfNoReportResultPermissions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		canReportResult, err := GamesDAO.HasReportResultPermissions(
			ctx.GetInt("leagueId"),
			ctx.GetInt("urlId"),
			ctx.GetInt("userId"),
		)
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if !canReportResult {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "noReportResultPermissions"})
		} else {
			ctx.Next()
		}
	}
}

func failIfCannotJoinLeague() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		canJoin, err := LeaguesDAO.CanJoinLeague(ctx.GetInt("leagueId"), ctx.GetInt("userId"))
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if !canJoin {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "canNotJoin"})
		} else {
			ctx.Next()
		}
	}
}

func failIfNotLeagueAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lp, err := LeaguesDAO.GetLeaguePermissions(ctx.GetInt("leagueId"), ctx.GetInt("userId"))
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if !lp.Administrator {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "notAdmin"})
		} else {
			ctx.Next()
		}
	}
}

func failIfNotTeamAdministrator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lp, tp, err := getLeagueAndTeamPermissions(ctx.GetInt("leagueId"), ctx.GetInt("urlId"), ctx.GetInt("userId"))
		if checkErr(ctx, err) {
			ctx.Abort()
		}
		isAdmin := lp.Administrator || tp.Administrator

		if !isAdmin {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "notTeamAdmin"})
		} else {
			ctx.Next()
		}
	}
}

func failIfCanNotEditTeamInformation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lp, tp, err := getLeagueAndTeamPermissions(ctx.GetInt("leagueId"), ctx.GetInt("urlId"), ctx.GetInt("userId"))
		if checkErr(ctx, err) {
			ctx.Abort()
		}
		canEdit := lp.Administrator || tp.Administrator || lp.EditTeams || tp.Information

		if !canEdit {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "noEditTeamInformationPermissions"})
		} else {
			ctx.Next()
		}
	}
}

func failIfTeamActive() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isActive, err := TeamsDAO.IsTeamActive(ctx.GetInt("leagueId"), ctx.GetInt("urlId"))
		if checkErr(ctx, err) {
			ctx.Abort()
		} else if isActive {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "teamIsActive"})
		} else {
			ctx.Next()
		}
	}
}