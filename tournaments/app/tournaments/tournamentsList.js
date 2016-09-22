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
    $http.get('/api/tournaments').success(function(data){
        scope.tournaments = data;
    });

    scope.tournament = {};
    this.createTournament = function(){
        scope.tournament.date = moment(scope.tournament.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
        $http.post('/api/tournaments', scope.tournament).success(function(data){
            scope.tournament.id = data;
            scope.tournaments.push(scope.tournament);
            scope.tournament = {};
        });
    };
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
        templateUrl: "tournaments/create-tournament.html"
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
