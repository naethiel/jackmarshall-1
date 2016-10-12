'use strict';

angular.module('tournamentsList', ['ngRoute', 'ui.bootstrap', 'ngAnimate'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/tournaments/list', {
        templateUrl: 'tournaments/views/tournamentsList/tournaments-list.html',
        controller: 'TournamentsListCtrl'
    });
}])
.controller('PopupCtrl', function ($uibModalInstance, tournament, scopeParent) {

    this.ok = function () {
        scopeParent.deleteTournament(tournament);
        $uibModalInstance.close();
    };

    this.cancel = function () {
        $uibModalInstance.dismiss('cancel');
    };
})


.controller('TournamentsListCtrl', ['$http', '$uibModal', function($http, $uibModal) {
    var scope = this;
    scope.tournaments = [];
    scope.tournament = {};

    scope.newTournamentCollapsed = false;
    scope.futureTournamentsCollapsed = false;
    scope.pastTournamentsCollapsed = false;

    $http.get('/api/tournaments').success(function(data){
        scope.tournaments = data;
    });


    this.createTournament = function(){
        scope.tournament.date = moment(scope.tournament.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
        $http.post('/api/tournaments', scope.tournament).success(function(data){
            scope.tournament.id = data;
            scope.tournaments.push(scope.tournament);
            scope.tournament = {};
            scope.newTournamentCollapsed = true;
            scope.futureTournamentsCollapsed = false;
        });
    };

    this.deleteTournament = function(tournament){
        $http.delete('/api/tournaments/'+tournament.id).success(function(data){
            scope.tournaments.splice(scope.tournaments.indexOf(tournament), 1);
        });
    };

    this.confirmDelete = function (tournament) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: 'tournaments/views/tournamentsList/delete_popup.html',
            controller: 'PopupCtrl',
            controllerAs: 'PopupCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                tournament: function () {
                    return tournament;
                },
                scopeParent: function(){
                    return scope;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };


}])

.directive("futureTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentsList/future-tournaments.html"
    };
})

.directive("pastTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentsList/past-tournaments.html"

    };
})

.directive("createTournament", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentsList/create-tournament.html"
    };
})

.filter('isFuture', function() {
    return function(items, dateFieldName) {
        return items.filter(function(item){
            return moment(item[dateFieldName || 'date']).isSameOrAfter(new Date(),'day');
        })
    }
})

.filter('isPast', function() {
    return function(items, dateFieldName) {
        return items.filter(function(item){
            return moment(item[dateFieldName || 'date']).isBefore(new Date(), 'day');
        })
    }
})

;
