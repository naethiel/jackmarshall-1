'use strict';

app.controller('GamesCtrl', ["$rootScope", "TournamentService", function ($rootScope, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.game = {};

    this.setWin = function(game, player_index, opponent_index){
        game.results[player_index].victory_points = 1;
        game.results[opponent_index].victory_points = 0;
    };

    this.setLoss = function(game, player_index, opponent_index){
        game.results[player_index].victory_points = 0;
        game.results[opponent_index].victory_points = 1;
    };

    this.onDropComplete=function(source, destination){
        var sourceTemp = JSON.parse(JSON.stringify(source));
        source.name = destination.name;
        source.faction = destination.faction;
        source.payed_fee = destination.payed_fee;
        source.lists = destination.lists;
        source.leave = destination.leave;
        destination.name = sourceTemp.name;
        destination.faction = sourceTemp.faction;
        destination.payed_fee = sourceTemp.payed_fee;
        destination.lists = sourceTemp.lists;
        destination.leave = sourceTemp.leave;
        tournamentService.verifyRound(scope.tournament, scope.roundNumber);
    };
}]);
