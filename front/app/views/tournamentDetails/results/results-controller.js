'use strict';

app.controller('ResultsCtrl', ["$rootScope", "$routeParams", "$uibModal", "uuid", "TournamentService", "UtilsService", function ($rootScope, $routeParams, $uibModal, uuid, tournamentService, utilsService) {
    var scope = this;
    scope.error = undefined;
    scope.players = {};
    scope.sortType = 'victory_points';
    scope.sortFields = {
        'victory_points':     ['result.victory_points','result.sos','result.scenario_points','result.destruction_points'],
        'sos':                ['result.sos','result.victory_points','result.scenario_points','result.destruction_points'],
        'scenario_points':    ['result.scenario_points','result.victory_points','result.sos','result.destruction_points'],
        'destruction_points': ['result.destruction_points','result.victory_points','result.sos','result.scenario_points'],
    }
    scope.sortOrder = true;

    $rootScope.$on("UpdateResult", function(){
        scope.error = null;
        tournamentService.getResults($routeParams.id).then(function(players){
            scope.players = players
            scope.addBest(scope.players)
        }).catch(function(err){
            scope.error = err;
        });
    });

    scope.addBest = function(players){
        var vp = {index: -1,max: 0}
        var sos = {index: -1,max: 0}
        var sp = {index: -1,max: 0}
        var dp = {index: -1,max: 0}
        var ck = {index: -1,max: 0}

        angular.forEach(players, function(player, id) {
            if (player.result.victory_points >= vp.max) {
                vp.index = id
                vp.max = player.result.victory_points
            }
            if (player.result.sos >= sos.max) {
                sos.index = id
                sos.max = player.result.sos
            }
            if (player.result.scenario_points >= sp.max) {
                sp.index = id
                sp.max = player.result.scenario_points
            }
            if (player.result.destruction_points >= dp.max) {
                dp.index = id
                dp.max = player.result.destruction_points
            }
            if (player.result.caster_kills >= ck.max) {
                ck.index = id
                ck.max = player.result.caster_kills
            }
        });

        players[vp.index].result.best_vp = true
        players[sos.index].result.best_sos = true
        players[sp.index].result.best_sp = true
        players[dp.index].result.best_dp = true
        players[ck.index].result.best_ck = true

    }

    scope.bbCodeResults = function(players) {
        var params = {
            animation: false,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: 'views/tournamentDetails/results/bbcode-popup.html',
            controller: 'ResultsBBCodeCtrl',
            controllerAs: 'ResultsBBCodeCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                players: function () {
                    return players;
                },
                scopeParent: function(){
                    return scope;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };
}]);
