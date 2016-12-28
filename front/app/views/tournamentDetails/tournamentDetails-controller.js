'use strict';

app.controller('TournamentCtrl', ['$routeParams', 'TournamentService', function($routeParams, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.error = undefined;

    tournamentService.get($routeParams.id).then(function(tournament){
        scope.tournament = tournament;
    }).catch(function(err){
        scope.error = err;
    });
}]);
