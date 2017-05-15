'use strict';

app.controller('GamesCtrl', ["$rootScope", "TournamentService", function ($rootScope, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.game = {};
    scope.errorUpdate = undefined;

    this.changeRes = function(game, player_index, opponent_index){
        game.results[player_index].victory_points = !game.results[player_index].victory_points;
        game.results[opponent_index].victory_points = !game.results[player_index].victory_points;
        this.updateGame()
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
        this.updateGame()
        tournamentService.verifyRound(scope.tournament, scope.roundNumber);
    };

    this.updateGame = function(){
        scope.errorUpdate = null;
        tournamentService.update(scope.tournament).then(function(id){
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };
}]);
