'use strict';

app.controller('CreateTournamentCtrl', ["$scope", "TournamentService", function ($scope, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.error = undefined;

    this.createTournament = function(){
        scope.error = null
        tournamentService.create(scope.tournament).then(function(id){
            scope.tournament.id = id;
            scope.tournaments.push(scope.tournament);
            scope.tournament = {};
            document.getElementById("tournament_name").focus()
            $scope.tournamentCreateForm.$setUntouched();
        }).catch(function(err){
            scope.error = err;
        });
    };
}]);
