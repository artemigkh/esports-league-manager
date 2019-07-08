import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {httpOptions, httpOptionsForm} from "./http-options";
import {Observable} from "rxjs/Rx";
import {TeamCoreWithIcon, TeamId, TeamPermissionsCore, TeamWithPlayers, TeamWithRosters} from "../interfaces/Team";
import {PlayerCore, PlayerId} from "../interfaces/Player";

// function sortMainRosterByPosition(team: Team) {
//     let sortedRoster = [];
//     ['top', 'jungle', 'middle', 'support', 'bottom'].forEach((role: string) => {
//         team.players.forEach((player: Player) => {
//             if(player.position.toLowerCase() == role) {
//                 sortedRoster.push(player);
//             }
//         });
//     });
//     team.players = sortedRoster;
// }

@Injectable()
export class TeamsService {
    constructor(private http: HttpClient) {
    }

    public getLeagueTeams(): Observable<TeamWithPlayers[]> {
        return this.http.get<TeamWithPlayers[]>('http://localhost:8080/api/v1/teams', httpOptions)
    }

    public getLeagueTeamsWithRosters(): Observable<TeamWithRosters[]> {
        return this.http.get<TeamWithRosters[]>('http://localhost:8080/api/v1/teamsWithRosters', httpOptions)
    }

    public getTeamWithRosters(teamId: string) {
        return this.http.get<TeamWithRosters>(
            'http://localhost:8080/api/v1/teams/' + teamId + "/withRosters", httpOptions)
    }

    public createTeam(form: FormData): Observable<TeamId> {
        return this.http.post<TeamId>('http://localhost:8080/api/v1/teams', form, httpOptionsForm)
    }

    public updateTeam(teamId: number, form: FormData): Observable<null> {
        return this.http.put<null>('http://localhost:8080/api/v1/teams/' + teamId, form, httpOptionsForm)
    }

    public deleteTeam(teamId: number): Observable<null> {
        return this.http.delete<null>('http://localhost:8080/api/v1/teams/' + teamId, httpOptions)
    }

    public createPlayer(teamId: number, player: PlayerCore): Observable<PlayerId> {
        return this.http.post<PlayerId>('http://localhost:8080/api/v1/teams/' + teamId + '/players',
            player, httpOptions)
    }

    public updatePlayer(teamId: number, playerId: number, player: PlayerCore): Observable<null> {
        return this.http.put<null>('http://localhost:8080/api/v1/teams/' + teamId + '/players/' + playerId,
            player, httpOptions)
    }

    public deletePlayer(teamId: number, playerId: number,): Observable<null> {
        return this.http.delete<null>('http://localhost:8080/api/v1/teams/' + teamId + '/players/' + playerId,
            httpOptions)
    }

    public updateTeamManagerPermissions(teamId: number, userId: number, permissions: TeamPermissionsCore) {
        return this.http.put<null>('http://localhost:8080/api/v1/teams/' + teamId + '/permissions/' + userId,
            permissions, httpOptions)
    }

    //
    // public getTeamManagers(): Observable<any> {
    //     return this.http.get('http://localhost:8080/api/leagues/teamManagers', httpOptions);
    // }
    //
    // public updateManagerPermissions(teamId: number, userId: number, administrator: boolean, information: boolean,
    //                                 players: boolean, reportResults: boolean) {
    //     return this.http.put('http://localhost:8080/api/teams/updatePermissions', {
    //         teamId: teamId,
    //         userId : userId,
    //         administrator: administrator,
    //         information: information,
    //         players: players,
    //         reportResults: reportResults
    //     }, httpOptions)
    // }
    //
    // public getTeamInformation(teamId: number): Observable<Object> {
    //     let url = "";
    //     switch(this.leagueService.getGame()) {
    //         case 'leagueoflegends': {
    //             url = 'http://localhost:8080/api/league-of-legends/teams/';
    //             break;
    //         }
    //         default: {
    //             url = 'http://localhost:8080/api/teams/';
    //         }
    //     }
    //     return new Observable(observer => {
    //         this.http.get(url + teamId, httpOptions).subscribe(
    //         (next: Team) => {
    //                 console.log(next);
    //                 let players = next.players;
    //                 console.log(players);
    //                 let team = next;
    //                 team.substitutes = [];
    //                 team.players = [];
    //                 if(players) {
    //                     players.forEach((player: any)=> {
    //                         if(player.mainRoster) {
    //                             team.players.push(player);
    //                         } else {
    //                             team.substitutes.push(player);
    //                         }
    //                     });
    //                 }
    //                 if(this.leagueService.getGame() == 'leagueoflegends') {
    //                     sortMainRosterByPosition(team);
    //                 }
    //                 team.id = teamId;
    //                 observer.next(team);
    //                 observer.complete();
    //             }, error => {
    //                 observer.error(error);
    //                 console.log(error);
    //             }
    //         );
    //     });
    // }
    //
    // public addPlayerInformationToTeam(team: Team): Observable<Team> {
    //     return new Observable(observer => {
    //         this.http.get('http://localhost:8080/api/teams/' + team.id, httpOptions).subscribe(
    //             (next: GtiTeam) => {
    //                 if(next.players) {
    //                     next.players.forEach(player=> {
    //                         if(player.mainRoster) {
    //                             team.players.push(player);
    //                         } else {
    //                             team.substitutes.push(player);
    //                         }
    //                     });
    //                 } else {
    //                     team.players = [];
    //                     team.substitutes = [];
    //                 }
    //
    //
    //                 observer.next(team)
    //             }, error => {
    //                 observer.error(error);
    //                 console.log(error);
    //             }
    //         );
    //     });
    // }
    //
    // public getTeamSummary(): Observable<Team[]> {
    //     return new Observable(observer => {
    //         this.http.get('http://localhost:8080/api/leagues/teamSummary', httpOptions).subscribe(
    //             (next: Team[]) => {
    //                 if(next == null) {
    //                     observer.next([]);
    //                 } else {
    //                     let teams = next;
    //                     teams.forEach(team => {
    //                         team.players = [];
    //                         team.substitutes = [];
    //                     });
    //                     observer.next(teams)
    //                 }
    //             }, error => {
    //                 console.log(error);
    //                 observer.error(error);
    //             }
    //         );
    //     });
    // }
    //
    // public addTeamInformation(games: Game[], teams: Team[]) {
    //     games.forEach(game => {
    //         teams.forEach(team => {
    //             if(game.team1Id == team.id) {
    //                 game.team1 = team;
    //             } else if (game.team2Id == team.id) {
    //                 game.team2 = team;
    //             }
    //         })
    //     })
    // }
}
