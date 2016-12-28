'use strict';

app.controller('ListCtrl', ['TournamentService', function(tournamentService) {
    var scope = this;
    scope.tournaments = [];
    scope.tournament = {};
    scope.error = undefined;

    scope.newTournamentCollapsed = false;
    scope.futureTournamentCollapsed = false;
    scope.pastTournamentCollapsed = true;

    this.createTournament = function(){
        scope.error = null
        tournamentService.create(scope.tournament).then(function(id){
            scope.tournament.id = id;
            scope.tournaments.push(scope.tournament);
            scope.tournament = {};
            scope.newTournamentCollapsed = true;
        }).catch(function(err){
            scope.error = err;
        });
    };
}]);
