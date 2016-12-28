'use strict';

app.controller('ListCtrl', ['$uibModal', 'TournamentService', function($uibModal, tournamentService) {
    var scope = this;
    scope.tournaments = [];
    scope.error = undefined;

    scope.futureTournamentCollapsed = false;
    scope.pastTournamentCollapsed = true;

    tournamentService.getAll().then(function(tournaments){
        scope.tournaments = tournaments;
    }).catch(function(){
        scope.error = true;
    });

    this.confirmDelete = function (tournament) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: '/views/tournamentList/tournamentDelete/tournament-delete-popup.html',
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
