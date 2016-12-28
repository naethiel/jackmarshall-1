'use strict';

app.controller('ListCtrl', ['$uibModal', 'TournamentService', function($uibModal, tournamentService) {
    var scope = this;
    scope.tournaments = [];
    scope.tournament = {};
    scope.error = undefined;
    scope.errorGetAll = undefined;

    scope.newTournamentCollapsed = false;
    scope.futureTournamentCollapsed = false;
    scope.pastTournamentCollapsed = true;

    tournamentService.getAll().then(function(tournaments){
        scope.tournaments = tournaments;
    }).catch(function(){
        scope.errorGetAll = true;
    });

    this.createTournament = function(){
        scope.error = null
        tournamentService.create(scope.tournament).then(function(id){
            scope.tournament.id = id;
            scope.tournaments.push(scope.tournament);
            scope.tournament = {};
            scope.newTournamentCollapsed = true;
        }).catch(function(err){
            scope.error = err;
        });
    };

    this.confirmDelete = function (tournament) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: '/views/tournamentList/tournament-delete-popup.html',
            controller: 'DeleteTournamentCtrl',
            controllerAs: 'DeleteCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                tournament: function () {
                    return tournament;
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

}]);
