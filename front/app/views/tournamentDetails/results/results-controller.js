'use strict';

app.controller('ResultsCtrl', ["$rootScope", "$routeParams", "$uibModal", "uuid", "TournamentService", "UtilsService", function ($rootScope, $routeParams, $uibModal, uuid, tournamentService, utilsService) {
    var scope = this;
    scope.error = undefined;
    scope.players = {};

    $rootScope.$on("UpdateResult", function(){
        scope.error = null;
        tournamentService.getResults($routeParams.id).then(function(players){
            scope.players = players;
        }).catch(function(err){
            scope.error = err;
        });
    });

    this.compare = function(a, b) {
        if (a.value.victory_points === b.value.victory_points) {
            if (a.value.sos === b.value.sos){
                if (a.value.scenario_points === b.value.scenario_points) {
                    if (a.value.destruction_points < b.value.destruction_points){
                        return 1
                    }else{
                        return -1
                    }
                } else {
                    if (a.value.scenario_points < b.value.scenario_points){
                        return 1
                    }else{
                        return -1
                    }
                }
            } else {
                if( a.value.sos < b.value.sos){
                    return 1
                }else{
                    return -1
                }
            }
        } else {
            if (a.value.victory_points < b.value.victory_points){
                return 1
            }else{
                return -1
            }
        }
    };

    this.bbCodeResults = function(score) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: 'views/tournamentDetails/results/bbcode-popup.html',
            controller: 'ResultsBBCodeCtrl',
            controllerAs: 'ResultsBBCodeCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                score: function () {
                    return score;
                },
                scopeParent: function(){
                    return scope;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };
}]);
