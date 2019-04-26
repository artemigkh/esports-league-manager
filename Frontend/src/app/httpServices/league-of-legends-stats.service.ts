import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {LeagueService} from "./leagues.service";
import {Observable} from "rxjs";
import {httpOptions} from "./http-options";

@Injectable()
export class LeagueOfLegendsStatsService {
    constructor(private http: HttpClient, private leagueService: LeagueService) {}

    public getPlayerStats(): Observable<any> {
        return this.http.get('http://localhost:8080/api/league-of-legends/stats/player', httpOptions);
    }

    public getTeamStats(): Observable<any> {
        return this.http.get('http://localhost:8080/api/league-of-legends/stats/team', httpOptions);
    }

    public getChampionStats(): Observable<any> {
        return this.http.get('http://localhost:8080/api/league-of-legends/stats/champion', httpOptions);
    }
}
