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
        game.results[player_index].victory_points = 1-game.results[player_index].victory_points;
        game.results[opponent_index].victory_points = 1-game.results[player_index].victory_points;
        var reverse = function() {
            game.results[0].victory_points = 1-game.results[0].victory_points;
            game.results[1].victory_points = 1-game.results[1].victory_points;
        }
        this.updateGame(game, reverse)
    };

    this.onDropComplete=function(source, destination){
        var temp = source.player;
        source.player = destination.player;
        destination.player = temp;
        var reverse = function() {
            var temp = source.player;
            source.player = destination.player;
            destination.player = temp;
        }
        this.updateGame(reverse)
        tournamentService.verifyRound(scope.tournament, scope.roundNumber);
    };

    this.updateGame = function(reverse){
        scope.errorUpdate = null;
        tournamentService.update(scope.tournament).then(function(){
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            reverse();
            scope.errorUpdate = true;
        })
    };
}]);
