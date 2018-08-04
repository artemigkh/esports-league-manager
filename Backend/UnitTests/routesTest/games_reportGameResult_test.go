package routesTest

import (
	"testing"
	"github.com/gin-gonic/gin"
	"esports-league-manager/Backend/Server/routes"
	"bytes"
	"encoding/json"
	"esports-league-manager/mocks"
	"github.com/stretchr/testify/mock"
	"errors"
	"esports-league-manager/Backend/Server/databaseAccess"
)

func createGamesReportBody(winnerId, scoreTeam1, scoreTeam2 int) *bytes.Buffer {
	reqBody := routes.GameReportInformation{
		WinnerID:   winnerId,
		ScoreTeam1: scoreTeam1,
		ScoreTeam2: scoreTeam2,
	}
	reqBodyB, _ := json.Marshal(&reqBody)
	return bytes.NewBuffer(reqBodyB)
}

func testReportGameResultNoActiveLeague(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(-1, nil)

	routes.ElmSessions = mockSession

	httpTest(t, nil, "POST", "/report/1", 403, testParams{Error: "noActiveLeague"})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testReportGameResultSessionError(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(1, errors.New("fake session error"))

	routes.ElmSessions = mockSession

	httpTest(t, nil, "POST", "/report/1", 500, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testReportGameResultNotLoggedIn(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(14, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(-1, nil)

	routes.ElmSessions = mockSession

	httpTest(t, nil, "POST", "/report/1", 403, testParams{Error: "notLoggedIn"})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testReportGameResultNoId(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(1, nil)

	routes.ElmSessions = mockSession

	httpTest(t, nil, "POST", "/report", 404, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testReportGameResultIdNotInt(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(1, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(2, nil)

	routes.ElmSessions = mockSession

	httpTest(t, nil, "POST", "/report/a", 400, testParams{Error: "IdMustBeInteger"})

	mock.AssertExpectationsForObjects(t, mockSession)
}

func testReportGameResultNoReportPermissions(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(14, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(15, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("HasReportResultPermissions", 14, 16, 15).
		Return(false, nil)

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, nil, "POST", "/report/16", 403, testParams{Error: "noReportResultPermissions"})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func testReportGameResultDatabaseError(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(14, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(15, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("HasReportResultPermissions", 14, 16, 15).
		Return(false, errors.New("fake db error"))

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, nil, "POST", "/report/16", 500, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func testReportGameResultMalformedBody(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(14, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(15, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("HasReportResultPermissions", 14, 16, 15).
		Return(true, nil)

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, nil, "POST", "/report/16", 400, testParams{Error: "malformedInput"})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func testReportGameResultGameDoesNotExist(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(14, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(15, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("HasReportResultPermissions", 14, 16, 15).
		Return(true, nil)
	mockGamesDao.On("GetGameInformation", 16, 14).
		Return(nil, nil)

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, createGamesReportBody(5, 2, 1),
		"POST", "/report/16", 400, testParams{Error: "gameDoesNotExist"})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func testReportGameResultCorrectReportDatabaseError(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(14, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(15, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("HasReportResultPermissions", 14, 16, 15).
		Return(true, nil)
	mockGamesDao.On("GetGameInformation", 16, 14).
		Return(&databaseAccess.GameInformation{}, nil)
	mockGamesDao.On("ReportGame", 16, 14, 5, 2, 1).
		Return(errors.New("fake db error"))

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, createGamesReportBody(5, 2, 1),
		"POST", "/report/16", 500, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func testReportGameResultCorrectReport(t *testing.T) {
	mockSession := new(mocks.SessionManager)
	mockSession.On("GetActiveLeague", mock.Anything).
		Return(14, nil)
	mockSession.On("AuthenticateAndGetUserID", mock.Anything).
		Return(15, nil)

	mockGamesDao := new(mocks.GamesDAO)
	mockGamesDao.On("HasReportResultPermissions", 14, 16, 15).
		Return(true, nil)
	mockGamesDao.On("GetGameInformation", 16, 14).
		Return(&databaseAccess.GameInformation{}, nil)
	mockGamesDao.On("ReportGame", 16, 14, 5, 2, 1).
		Return(nil)

	routes.ElmSessions = mockSession
	routes.GamesDAO = mockGamesDao

	httpTest(t, createGamesReportBody(5, 2, 1),
		"POST", "/report/16", 200, testParams{})

	mock.AssertExpectationsForObjects(t, mockSession, mockGamesDao)
}

func Test_ReportGameResult(t *testing.T) {
	//set up router and path to test
	gin.SetMode(gin.ReleaseMode) //opposite of gin.DebugMode to make tests faster by removing logging
	router = gin.New()

	router.Use(routes.Testing_Export_getActiveLeague())
	router.POST("/report/:id",
		routes.Testing_Export_authenticate(),
		routes.Testing_Export_getUrlId(),
		routes.Testing_Export_getReportResultPermissions(),
		routes.Testing_Export_reportGameResult)

	t.Run("NoActiveLeague", testReportGameResultNoActiveLeague)
	t.Run("SessionError", testReportGameResultSessionError)
	t.Run("NotLoggedIn", testReportGameResultNotLoggedIn)
	t.Run("NoId", testReportGameResultNoId)
	t.Run("IdNotInt", testReportGameResultIdNotInt)
	t.Run("NoReportPermissions", testReportGameResultNoReportPermissions)
	t.Run("DatabaseError", testReportGameResultDatabaseError)
	t.Run("MalformedBody", testReportGameResultMalformedBody)
	t.Run("GameDoesNotExist", testReportGameResultGameDoesNotExist)
	t.Run("CorrectReportDatabaseError", testReportGameResultCorrectReportDatabaseError)
	t.Run("CorrectReport", testReportGameResultCorrectReport)

}



