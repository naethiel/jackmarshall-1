'use strict';

angular.module('tournamentsList', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/tournaments/list', {
    templateUrl: 'tournaments/tournaments-list.html',
    controller: 'TournamentsListCtrl'
  });
}])

.controller('TournamentsListCtrl', [function() {

}]);
