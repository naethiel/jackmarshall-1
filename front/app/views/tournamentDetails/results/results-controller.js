'use strict';

app.controller('ResultsCtrl', ["$rootScope", "$routeParams", "$uibModal", "uuid", "TournamentService", "UtilsService", function ($rootScope, $routeParams, $uibModal, uuid, tournamentService, utilsService) {
    var scope = this;
    scope.error = undefined;
    scope.score = [];

    tournamentService.getResults($routeParams.id).then(function(score){
        scope.score = score;
    }).catch(function(err){
        scope.error = err;
    });


    $rootScope.$on("UpdateResult", function(){
        scope.error = null;
        tournamentService.getResults($routeParams.id).then(function(score){
            scope.score = score;
        }).catch(function(err){
            scope.error = err;
        });
    });

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
