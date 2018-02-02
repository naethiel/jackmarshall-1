'use strict';

app.controller('AssignCtrl', ['$rootScope', '$routeParams', '$localStorage', 'TournamentService', function($rootScope, $routeParams, $localStorage, tournamentService) {

    if($localStorage.currentUser == null){
        $location.path( "/auth/login" );
    }

    var scope = this;
    scope.assignements = [];
    scope.roundNumber = {};
    scope.error = false;

    tournamentService.get($routeParams.id).then(function(tournament){
        scope.roundNumber = tournament.rounds.length;
        if (scope.roundNumber > 0) {
            var games = tournament.rounds[scope.roundNumber-1].games
            games.sort(function (a, b) {
                return a.table.name.localeCompare(b.table.name);
            });
            while (games.length > 0){
                scope.assignements.push(games.splice(0,12))
            }
        }
    }).catch(function(err){
        console.error(err)
        scope.error = err;
    });

}]);
