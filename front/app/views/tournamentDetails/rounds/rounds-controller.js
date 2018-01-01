'use strict';

app.controller('RoundsCtrl', ["$rootScope", "$route", "$uibModal", "$scope", "TournamentService", function ($rootScope, $route, $uibModal, $scope, tournamentService) {
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

    this.bbCodeRound = function(round) {
        var params = {
            animation: false,
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

    this.confirmDelete = function (round) {
        var params = {
            animation: false,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: '/views/tournamentDetails/rounds/round-delete-popup.html',
            controller: 'DeleteRoundCtrl',
            controllerAs: 'DeleteRoundCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                tournament: function () {
                    return scope.tournament;
                },
                round: function () {
                    return round;
                },
                scopeParent: function(){
                    return scope;
                },
                tournamentService: function(){
                    return tournamentService;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };

    this.openAssignements = function(id){
        window.open('/#!/tournament/'+id+'/assignements');;
    }

    this.compare = function(a, b) {
        return naturalSort(scope.tournament.tables[a.value].name, scope.tournament.tables[b.value].name);
    };
    // this.openAssignements = function(id){
    //     window.open('views/tournamentDetails/rounds/assignements.html?id='+id);
    // }
}]);
