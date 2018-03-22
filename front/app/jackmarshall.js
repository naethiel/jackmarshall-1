'use strict';

var app = angular.module('jackmarshall', [
    'ngRoute',
    'ui.bootstrap',
    'ngAnimate',
    'ngDraggable',
    'ngStorage',
    'angular-uuid',
    'angularMoment'
]);

app.config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
    $locationProvider.html5Mode(true);
    
    $routeProvider.when('/auth/login', {
        templateUrl: 'views/auth/login.html',
    });
    $routeProvider.when('/auth/new', {
        templateUrl: 'views/auth/new-user.html',
    });
    $routeProvider.when('/tournament/list', {
        templateUrl: 'views/tournamentList/tournament-list.html',
    });
    $routeProvider.when('/tournament/:id', {
        templateUrl: '/views/tournamentDetails/tournament-details.html',
    });
    $routeProvider.when('/timer', {
        templateUrl: '/views/timer/timer.html',
    });
    $routeProvider.when('/tournament/:id/assignements', {
        templateUrl: '/views/assignements/assignements.html',
    });
    $routeProvider.otherwise({redirectTo: '/auth/login'});
}]);
