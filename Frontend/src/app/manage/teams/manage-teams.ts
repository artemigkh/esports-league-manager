import {Component, Inject} from "@angular/core";
import {LeagueService} from "../../httpServices/leagues.service";
import {Team} from "../../interfaces/Team";
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from "@angular/material";
import {Game} from "../../interfaces/Game";
import {WarningPopup} from "../warningPopup/warning-popup";
import {TeamsService} from "../../httpServices/teams.service";
import {Action} from "../actions";
import {Id} from "../../httpServices/api-return-schemas/id";
import {ManageComponentInterface} from "../manage-component-interface";
import {UserService} from "../../httpServices/user.service";
import {TeamPermissions, UserPermissions} from "../../httpServices/api-return-schemas/permissions";
import {FormBuilder, FormGroup} from "@angular/forms";

class TeamData {
    title: string;
    action: Action;
    team: Team;
    caller: ManageTeamsComponent;
}

@Component({
    selector: 'app-manage-teams',
    templateUrl: './manage-teams.html',
    styleUrls: ['./manage-teams.scss'],
})
export class  ManageTeamsComponent implements ManageComponentInterface {
    displayedColumns: string[] = ['team'];
    teams: Team[];
    constructor(private leagueService: LeagueService,
                private teamsService: TeamsService,
                private userService: UserService,
                public dialog: MatDialog) {
        console.log("team service created: ", this.teamsService);
        this.teamsService.getTeamSummary().subscribe(
            teamSummary => {
                let teams = teamSummary;
                this.userService.getUserPermissions().subscribe(
                    (next: UserPermissions) => {
                        this.teams = [];
                        teams.forEach((team: Team) => {
                            if(next.leaguePermissions.administrator || next.leaguePermissions.editTeams) {
                                this.teams.push(team);
                            } else {
                                next.teamPermissions.forEach((teamPermission: TeamPermissions) => {
                                    if(team.id == teamPermission.id &&
                                        (teamPermission.administrator || teamPermission.information)) {
                                        this.teams.push(team);
                                    }
                                });
                            }
                        });
                    }, error => {
                        console.log(error);
                    }
                );
            }, error => {
                console.log(error);
        });
    }

    newTeamPopup(): void {
        const dialogRef = this.dialog.open(ManageTeamPopup, {
            width: '500px',
            data: {
                title: "Create New Team",
                action: Action.Create,
                team: {
                    name: "",
                    tag: "",
                    description: ""
                },
                caller: this
            },
            autoFocus: false
        });
    }

    editTeamPopup(team: Team): void {
        console.log(team);
        this.teamsService.getTeamInformation(team.id).subscribe(
            (next: Team) => {
                next.id = team.id;
                const dialogRef = this.dialog.open(ManageTeamPopup, {
                    width: '500px',
                    data: {
                        title: "Edit Team Information",
                        action: Action.Edit,
                        team: next,
                        caller: this
                    },
                    autoFocus: false
                });
            }, error => {
                console.log(error);
            }
        );
    }

    warningPopup(team: Team): void {
        const dialogRef = this.dialog.open(WarningPopup, {
            width: '500px',
            data: {
                entity: "team",
                name: team.name,
                Id: team.id,
                caller: this
            },
            autoFocus: false
        });
    }

    private updateTeamsList(): void {
        this.teamsService.getTeamSummary().subscribe(
            teamSummary => {
                this.teams = teamSummary;
            }, error => {
                console.log(error);
            });
    }

    notifyCreateSuccess(id: number): void {
        this.updateTeamsList();
    }

    notifyUpdateSuccess(id: number): void {
        this.updateTeamsList();
    }

    notifyDelete(id: number): void {
        console.log("attempt delete team with id ", id);
        this.teamsService.deleteTeam(id).subscribe(
            next => {
                console.log("deleted team with id ", id);
                this.updateTeamsList();
            }, error => {
                console.log('failed to delete team, reason:', error);
            }
        )
    }
}

@Component({
    selector: 'manage-teams-popup',
    templateUrl: 'manage-teams-popup.html',
    styleUrls: ['./manage-teams-popup.scss'],
})
export class ManageTeamPopup {
    teams: Team[];
    action: Action;
    name: string;
    tag: string;
    description: string;
    id: number;
    iconContainer: FormGroup;
    constructor(
        public dialogRef: MatDialogRef<ManageTeamPopup>,
        @Inject(MAT_DIALOG_DATA) public data: TeamData,
        private leagueService: LeagueService,
        private teamsService: TeamsService,
        private _formBuilder: FormBuilder,) {
        this.action = data.action;
        this.name = data.team.name;
        this.tag = data.team.tag;
        this.description = data.team.description;
        this.id = data.team.id;
        this.iconContainer = this._formBuilder.group({icon: null});
        console.log(this.data.team);
        this.teamsService.getTeamSummary().subscribe(
            teamSummary => {
                this.teams = teamSummary;
            }, error => {
                console.log(error);
            });
    }

    onFileChange(event) {
        if(event.target.files.length > 0) {
            this.iconContainer.value.icon = event.target.files[0];
        }
    }

    OnCancel(): void {
        this.dialogRef.close();
    }

    OnConfirm(): void {
        let form = new FormData();
        form.append("name", this.name);
        form.append("tag", this.tag);
        form.append("description", this.description);
        form.append("icon", this.iconContainer.value.icon);

        if(this.action == Action.Create) {
            this.teamsService.createNewTeamWithIcon(form).subscribe(
                (next: Id) => {
                    console.log("successfully created team");
                    this.data.caller.notifyCreateSuccess(next.id);
                    this.dialogRef.close();
                }, error => {
                    console.log("error during team creation");
                    console.log(error);
                    this.dialogRef.close();
                }
            )
        } else if(this.action = Action.Edit) {
            this.teamsService.updateTeamWithIcon(this.id, form).subscribe(
                next => {
                    console.log("successfully updated team");
                    this.data.caller.notifyUpdateSuccess(this.id);
                    this.dialogRef.close();
                }, error => {
                    console.log("error during team update");
                    console.log(error);
                    this.dialogRef.close();
                }
            )
        }
    }
}

