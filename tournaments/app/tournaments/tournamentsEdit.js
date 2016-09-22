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
        console.error(data);
        console.error(data.date);
        console.error(moment(data.date, 'YYYY-MM-DDThh:mm:ssZ').format('DD/MM/YYYY'));
        data.date = moment(data.date, 'YYYY-MM-DDThh:mm:ssZ').format('DD/MM/YYYY');

        scope.tournament = data;

    });
}])

.directive('tournamentDescription', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/tournament-description.html"
    };
})

;
