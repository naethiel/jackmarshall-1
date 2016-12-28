'use strict';

app.directive('playerList', function(){
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/players/player-list.html",
        scope: {},
        controller: 'PlayersCtrl',
        controllerAs: 'PlayersCtrl',
        bindToController: {
            tournament: '=tournament'
        }
    };
});
