'use strict';

app.controller('PlayersCtrl', ["$scope", "$rootScope", "$route", "uuid", "TournamentService", "UtilsService", function ($scope, $rootScope, $route, uuid, tournamentService, utilsService) {
    var scope = this;
    scope.tournament = {};
    scope.player = {};
    scope.errorAdd = undefined;
    scope.errorUpdate = undefined;
    scope.errorDelete = undefined;
    scope.playersCollapsed = false;

    $rootScope.$on("UpdateRounds", function(e, nb_round){
        scope.playersCollapsed = (nb_round > 0);
    });

    this.addPlayer = function(){
        scope.errorAdd = false;
        scope.player.id = uuid.v4();
        scope.tournament.players[scope.player.id] = scope.player
        tournamentService.update(scope.tournament).then(function(){
            scope.player = {};
            document.getElementById("add_player_name").focus()
            $scope.addPlayerForm.$setUntouched();
        }).catch(function(err){
            delete scope.tournament.players[scope.player.id]
            scope.errorAdd = true;
        })
    };

    this.deletePlayer = function(player){
        scope.errorDelete = false;
        delete scope.tournament.players[player.id]
        tournamentService.update(scope.tournament).then(function(){
        }).catch(function(err){
            scope.tournament.players[player.id] = player
            scope.errorDelete = true;
        })
    };

    this.updatePlayer = function(player){
        scope.errorUpdate = false;
        tournamentService.update(scope.tournament).then(function(){
            player.detailsVisible=false;
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };

    this.changeStatus = function(player){
        scope.errorUpdate = false;
        player.leave = !player.leave;
        tournamentService.update(scope.tournament).then(function(){
        }).catch(function(err){
            player.leave=!player.leave;
            scope.errorUpdate = true;
        })
    }

    this.compare = function(a, b) {
        return naturalSort(a.value, b.value);
    };


}]);
