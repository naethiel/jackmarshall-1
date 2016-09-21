'use strict';

angular.module('tournamentsList', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/tournaments/list', {
        templateUrl: 'tournaments/tournaments-list.html',
        controller: 'TournamentsListCtrl'
    });
}])

.controller('TournamentsListCtrl', ['$http', function($http) {
    var scope = this;
    scope.tournaments = [];
    $http.get('/tournaments').success(function(data){
        scope.tournaments = data;
    });
}])

.directive("futureTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/future-tournaments.html"
    };
})

.directive("pastTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/past-tournaments.html"

    };
})

.directive("createTournament", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/create-tournament.html",
        controller: function($http){
            var scope = this;
            scope.tournament = {};
            this.createTournament = function(){

                scope.tournament.date = moment(scope.tournament.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
                console.error(scope.tournament);
                $http.post('/tournaments', scope.tournament).success(function(data){
                    scope.tournament.id = data;
                    //TODO : add scope.tounament to tounaments list or redirect
                    scope.tournament = {};
                });
            }
        },
        controllerAs: "CreateCtrl"

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
