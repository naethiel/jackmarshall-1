'use strict';

angular.module('tournamentsEdit', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/tournaments/:id', {
        templateUrl: 'tournaments/tournament-edit.html',
        controller: 'TournamentsEditCtrl'
    });
}])

.controller('TournamentsEditCtrl', ['$http', '$routeParams', function($http, $routeParams) {
    var scope = this;
    scope.tournament = {};
    $http.get('/api/tournaments/'+$routeParams.id).success(function(data){
        scope.tournament = data;
    });
}])

;
