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


app.directive('addPlayer', function(){
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/players/player-add.html"
    };
});

app.directive('editPlayer', function(){
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/players/player-edit.html"
    };
});
