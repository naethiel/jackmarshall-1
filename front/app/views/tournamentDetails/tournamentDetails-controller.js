'use strict';

app.controller('TournamentCtrl', ['$rootScope', '$routeParams', 'TournamentService', function($rootScope, $routeParams, tournamentService) {

    if($localStorage.currentUser == null){
        $location.path( "/auth/login" );
    }

    var scope = this;
    scope.tournament = {};
    scope.error = undefined;

    tournamentService.get($routeParams.id).then(function(tournament){
        scope.tournament = tournament;
        $rootScope.$emit("SetTab", scope.tournament.rounds.length -1);
        scope.tournament.rounds.forEach(function(round){
            tournamentService.verifyRound(scope.tournament, scope.tournament.rounds.length -1)
        });
    }).catch(function(err){
        scope.error = err;
    });
}]);
