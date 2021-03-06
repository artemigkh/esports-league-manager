import random
from datetime import timedelta, datetime

from faker import Faker
from faker.providers import internet
from faker.providers import lorem

from .team import Team
from .availability import Availability
from .game import Game

fake = Faker()
fake.add_provider(internet)
fake.add_provider(lorem)

valid_game_strings = [
    "genericsport",
    "basketball",
    "curling",
    "football",
    "hockey",
    "rugby",
    "soccer",
    "volleyball",
    "waterpolo",
    "genericesport",
    "csgo",
    "leagueoflegends",
    "overwatch",
]


class League:
    def __init__(self, t):
        self.managers = []
        self.teams = []
        self.availabilities = []
        self.games = []

        self.public_view = True
        self.public_join = True

        current_time = datetime.utcnow()
        self.signup_start = current_time - timedelta(weeks=1)
        self.signup_end = current_time + timedelta(weeks=1)
        self.league_start = current_time + timedelta(weeks=3)
        self.league_end = current_time + timedelta(weeks=4)

        self.name = fake.slug()
        self.description = fake.text(max_nb_chars=500)
        self.game = random.choice(valid_game_strings)
        r = t.http.post("http://localhost:8080/api/v1/leagues", json={
            "name": self.name,
            "description": self.description,
            "game": self.game,
            "publicView": self.public_view,
            "publicJoin": self.public_join,
            "signupStart": int(self.signup_start.timestamp()),
            "signupEnd": int(self.signup_end.timestamp()),
            "leagueStart": int(self.league_start.timestamp()),
            "leagueEnd": int(self.league_end.timestamp())
        })

        t.assertEqual(201, r.status_code)
        self.league_id = r.json()["leagueId"]

    def update_to_middle_of_competition_time(self, t):
        current_time = datetime.utcnow()
        self.signup_start = current_time - timedelta(weeks=7)
        self.signup_end = current_time - timedelta(weeks=6)
        self.league_start = current_time - timedelta(weeks=5)
        self.league_end = current_time + timedelta(weeks=5)
        r = t.http.put("http://localhost:8080/api/v1/leagues", json={
            "name": self.name,
            "description": self.description,
            "game": self.game,
            "publicView": True,
            "publicJoin": True,
            "signupStart": int(self.signup_start.timestamp()),
            "signupEnd": int(self.signup_end.timestamp()),
            "leagueStart": int(self.league_start.timestamp()),
            "leagueEnd": int(self.league_end.timestamp())
        })

        t.assertEqual(200, r.status_code)

    def update_permissions(self, t, public_join, public_view=True):
        self.public_join = public_join
        self.public_view = public_view
        r = t.http.put("http://localhost:8080/api/v1/leagues", json={
            "name": self.name,
            "description": self.description,
            "game": self.game,
            "publicView": public_view,
            "publicJoin": public_join,
            "signupStart": int(self.signup_start.timestamp()),
            "signupEnd": int(self.signup_end.timestamp()),
            "leagueStart": int(self.league_start.timestamp()),
            "leagueEnd": int(self.league_end.timestamp())
        })
        t.assertEqual(200, r.status_code)

    def create_team(self, t, manager, name=None, tag=None):
        new_team = Team(t, self, manager, random.randint(0, 100), name, tag)
        self.teams.append(new_team)
        return new_team

    def create_availability(self, t, league, weekday, hour, minute, duration_minutes):
        new_availability = Availability(t, league, weekday, hour, minute, duration_minutes)
        self.availabilities.append(new_availability)
        return new_availability

    def create_game(self, t, team1_id, team2_id, game_time):
        new_game = Game(t, team1_id, team2_id, game_time)
        self.games.append(new_game)
        return new_game

    def get_team(self, team_id):
        return next((t for t in self.teams if t.team_id == team_id), None)

    def get_game(self, game_id):
        return next((g for g in self.games if g.game_id == game_id), None)

    def assert_server_data_consistent(self, t):
        r = t.http.get("http://localhost:8080/api/v1/leagues")
        t.assertEqual(200, r.status_code)
        self.assert_equal_json(t, r.json())

    def assert_equal_json(self, t, json):
        t.assertEqual(self.league_id, json["leagueId"])
        t.assertEqual(self.name, json["name"])
        t.assertEqual(self.description, json["description"])
        t.assertEqual(self.game, json["game"])
        t.assertEqual(self.public_view, json["publicView"])
        t.assertEqual(self.public_join, json["publicJoin"])
        t.assertEqual(int(self.signup_start.timestamp()), json["signupStart"])
        t.assertEqual(int(self.signup_end.timestamp()), json["signupEnd"])
        t.assertEqual(int(self.league_start.timestamp()), json["leagueStart"])
        t.assertEqual(int(self.league_end.timestamp()), json["leagueEnd"])

    def assert_teams_equal_json(self, t, json):
        for json_team in json:
            team = self.get_team(json_team["teamId"])
            team.assert_equal_json(t, json_team)

    def assert_managers_equal_json(self, t, json):
        for json_team in json:
            team = self.get_team(json_team["teamId"])
            team.assert_display_equal_json(t, json_team)
            for json_manager in json_team["managers"]:
                manager = next((m for m in team.managers if m.user_id == json_manager["userId"]), None)
                print("checking manager for team with id " + str(team.team_id))
                t.assertEqual(manager.user_id,  json_manager["userId"])
                t.assertEqual(manager.email, json_manager["email"])
                t.assertEqual(True, json_manager["administrator"])
                t.assertEqual(True, json_manager["information"])
                t.assertEqual(True, json_manager["games"])

    def assert_games_equal_json(self, t, json):
        for json_game in json:
            game = self.get_game(json_game["gameId"])
            game.assert_equal_json(t, json_game, self.teams)
