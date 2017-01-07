'use strict';

app.directive("editGame", function(){
    return {
        restrict: "E",
        templateUrl: "views/tournamentDetails/games/game-edit.html",
        scope: {},
        controller: 'GamesCtrl',
        controllerAs: 'GamesCtrl',
        bindToController: {
            tournament: '=tournament',
            game: '=game',
            roundNumber: '=roundnumber'
        }
    };
});
