'use strict';

angular.module('tournamentDetails', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/tournaments/:id', {
        templateUrl: 'tournaments/views/tournamentDetails/tournament-details.html',
        controller: 'TournamentsEditCtrl'
    });
}])

.controller('TournamentsEditCtrl', ['$http', '$routeParams', function($http, $routeParams) {
    var scope = this;
    scope.tournament = {};
    scope.player = {};
    scope.player.lists = ["",""];
    scope.table = {};
    scope.round = {};
    scope.score = [];
    $http.get('/api/tournaments/'+$routeParams.id).success(function(data){
        data.date = moment(data.date, 'YYYY-MM-DDThh:mm:ssZ').format('DD/MM/YYYY');
        scope.tournament = data;
    });

    $http.get('/api/tournaments/'+$routeParams.id+ '/results').success(function(data){
        scope.score = data;
    });


    this.getNextRound = function(){
        $http.get('/api/tournaments/'+scope.tournament.id+'/round').success(function(data){
            scope.round = data;
            scope.tournament.rounds[data.number] = data;
        });
    };

    this.addPlayer = function(){
        console.error(scope.tournament);
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.players.push(scope.player);
        temp.date = moment(temp.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            scope.tournament.players.push(scope.player);
            scope.player = {};
            scope.player.lists = ["",""];
        });
    };

    this.addTable = function(){
        console.error(scope.tournament);
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.tables.push(scope.table);
        temp.date = moment(temp.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            scope.tournament.tables.push(scope.table);
            scope.table = {};
        });
    };

    this.deletePlayer = function(player){
        console.error(scope.tournament);
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.players.splice(temp.players.indexOf(player), 1);
        temp.date = moment(temp.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            scope.tournament.players.splice(scope.tournament.players.indexOf(player), 1);
        });
    };

    this.deleteTable = function(table){
        console.error(scope.tournament);
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.tables.splice(temp.tables.indexOf(table), 1);
        temp.date = moment(temp.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            scope.tournament.tables.splice(scope.tournament.tables.indexOf(table), 1);
        });
    };

    this.updateTournament = function(){
        console.error(scope.tournament);
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.date = moment(temp.date, 'DD/MM/YYYY').format('YYYY-MM-DDThh:mm:ssZ');
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
        });
    };

}])

.directive("roundTabs", function() {
    return {
        restrict: "E",
        templateUrl: "/tournaments/views/tournamentDetails/rounds/round-tabs.html",
        controller: function() {
            this.tab = 1;

            this.isSet = function(checkTab) {
                return this.tab === checkTab;
            };

            this.setTab = function(activeTab) {
                this.tab = activeTab;
            };
        },
        controllerAs: "tab"
    };
})

.directive('tournamentDescription', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/tournament-description.html"
    };
})

.directive('addPlayer', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/views/tournamentDetails/players/add-player.html"
    };
})

.directive('editPlayer', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/views/tournamentDetails/players/edit-player.html"
    };
})

.directive('playersList', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/views/tournamentDetails/players/players-list.html"
    };
})

.directive('addTable', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/views/tournamentDetails/tables/add-table.html"
    };
})

.directive('editTable', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/views/tournamentDetails/tables/edit-table.html"
    };
})

.directive('tablesList', function(){
    return {
        restrict: "E",
        templateUrl: "/tournaments/views/tournamentDetails/tables/tables-list.html"
    };
})

.directive("roundList", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentDetails/rounds/round-list.html"
    };
})

.directive("editRound", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentDetails/rounds/edit-round.html"

    };
})

.directive("editGame", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentDetails/rounds/edit-game.html"
    };
})

.directive("tournamentResults", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentDetails/tournament/results.html"
    };
})

;
