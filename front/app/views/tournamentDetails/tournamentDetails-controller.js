'use strict';

app.controller('TournamentCtrl', ['$rootScope', '$routeParams', '$localStorage', 'TournamentService', function($rootScope, $routeParams, $localStorage, tournamentService) {

    if($localStorage.currentUser == null){
        $location.path( "/auth/login" );
    }

    var scope = this;
    scope.tournament = {};
    scope.error = false;

    tournamentService.get($routeParams.id).then(function(tournament){
        scope.tournament = tournament;
        $rootScope.$emit("SetTab", scope.tournament.rounds.length -1);
        $rootScope.$emit("UpdateRounds", scope.tournament.rounds.length);
        scope.tournament.rounds.forEach(function(round){
            tournamentService.verifyRound(scope.tournament, round.number)
        });
    }).catch(function(err){
        console.error(err)
        scope.error = err;
    });

}]);
