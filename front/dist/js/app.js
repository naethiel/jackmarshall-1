'use strict';

var api_endpoint = "localhost:8080";
var auth_endpoint = "localhost:8081";

var app = angular.module('jackmarshall', [
    'ngRoute',
    'ui.bootstrap',
    'ngAnimate',
    'ngDraggable',
    'ngStorage'
]);

app.config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
    $routeProvider.when('/auth', {
        templateUrl: 'views/auth/auth.html',
        controller: 'AuthCtrl'
    });
    $routeProvider.when('/tournament/list', {
        templateUrl: 'views/tournamentList/tournament-list.html',
        controller: 'TournamentListCtrl'
    });
    $routeProvider.otherwise({redirectTo: '/tournament/list'});
}]);
