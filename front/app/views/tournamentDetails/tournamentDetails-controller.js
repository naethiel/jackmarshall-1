'use strict';

app.controller('TournamentCtrl', ['$rootScope', '$routeParams', 'TournamentService', function($rootScope, $routeParams, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.error = undefined;

    tournamentService.get($routeParams.id).then(function(tournament){
        scope.tournament = tournament;
        $rootScope.$emit("SetTab", scope.tournament.rounds.length -1);
    }).catch(function(err){
        scope.error = err;
    });
}]);
