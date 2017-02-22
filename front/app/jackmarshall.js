'use strict';

var tournament_endpoint = "http://localhost:8080/api";
var auth_endpoint = "http://localhost:8081";

var app = angular.module('jackmarshall', [
    'ngRoute',
    'ui.bootstrap',
    'ngAnimate',
    'ngDraggable',
    'ngStorage',
    'angular-uuid'
]);

app.config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
    $routeProvider.when('/auth/login', {
        templateUrl: 'views/auth/login.html',
        controller: 'LoginCtrl'
    });
    $routeProvider.when('/auth/new', {
        templateUrl: 'views/auth/new-user.html',
        controller: 'NewUserCtrl'
    });
    $routeProvider.when('/tournament/list', {
        templateUrl: 'views/tournamentList/tournament-list.html',
        controller: 'ListCtrl'
    });
    $routeProvider.when('/tournament/:id', {
        templateUrl: '/views/tournamentDetails/tournament-details.html',
        controller: 'TournamentCtrl'
    });
    $routeProvider.otherwise({redirectTo: '/auth/login'});
}]);
