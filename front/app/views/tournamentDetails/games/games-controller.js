'use strict';

app.controller('GamesCtrl', ["$rootScope", "TournamentService", function ($rootScope, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.game = {};
    scope.errorUpdate = undefined;

    this.Player = function(index){
        return scope.tournament.players[scope.game.results[index].player]
    };
    this.Table = function(){
        return scope.tournament.tables[scope.game.table]
    };

    this.changeRes = function(game, player_index, opponent_index){
        if (game.results[player_index].victory_points===0 ||game.results[player_index].victory_points===false){
            game.results[player_index].victory_points = 1;
            game.results[opponent_index].victory_points = 0;
        } else {
            game.results[player_index].victory_points = 0;
            game.results[opponent_index].victory_points = 1;
        }
        // game.results[player_index].victory_points = !game.results[player_index].victory_points;
        // game.results[opponent_index].victory_points = !game.results[player_index].victory_points;
        this.updateGame()
    };

    this.onDropComplete=function(source, destination){
        var temp = source.player;
        source.player = destination.player;
        destination.player = temp;
        this.updateGame()
        tournamentService.verifyRound(scope.tournament, scope.roundNumber);
    };

    this.updateGame = function(){
        scope.errorUpdate = null;
        tournamentService.update(scope.tournament).then(function(){
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };
}]);
