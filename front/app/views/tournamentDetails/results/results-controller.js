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
            scope.players = players;
        }).catch(function(err){
            scope.error = err;
        });
    });

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
