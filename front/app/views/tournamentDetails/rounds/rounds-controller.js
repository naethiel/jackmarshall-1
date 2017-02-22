'use strict';

app.controller('RoundsCtrl', ["$rootScope", "$route", "$uibModal", "TournamentService", function ($rootScope, $route, $uibModal, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.round = {};
    scope.errorDelete = undefined;
    scope.errorUpdate = undefined;
    scope.successUpdate = undefined;

    this.updateRound = function(){
        scope.errorUpdate = null;
        scope.succesUpdate = null;
        tournamentService.update(scope.tournament).then(function(id){
            scope.successUpdate = true;
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };

    this.deleteRound = function(round){
        scope.errorDelete = null;
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.rounds.splice(temp.rounds.indexOf(round), 1);
        tournamentService.update(temp).then(function(id){
            scope.tournament.rounds.splice(scope.tournament.rounds.indexOf(round), 1);
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            scope.errorDelete = true;
        })
    };

    this.bbCodeRound = function(round) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: 'views/tournamentDetails/rounds/bbcode-popup.html',
            controller: 'RoundBBCodeCtrl',
            controllerAs: 'RoundBBCodeCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                round: function () {
                    return round;
                },
                scopeParent: function(){
                    return scope;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };

    this.openAssignements = function(id){
        window.open('views/tournamentDetails/rounds/assignements.html?id='+id);
    }
}]);
