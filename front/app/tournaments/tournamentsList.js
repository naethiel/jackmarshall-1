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
    $http.get('/tournaments/tournamentsList.json').success(function(data){
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
