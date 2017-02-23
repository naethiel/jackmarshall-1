'use strict';

app.controller('TournamentCtrl', ['$rootScope', '$routeParams', '$localStorage', '$scope', 'TournamentService', function($rootScope, $routeParams, $localStorage, $scope, tournamentService) {

    if($localStorage.currentUser == null){
        $location.path( "/auth/login" );
    }

    var scope = this;
    scope.tournament = {};
    scope.error = false;

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
