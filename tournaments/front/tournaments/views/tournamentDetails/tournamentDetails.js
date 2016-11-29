'use strict';

angular.module('tournamentDetails', ['ngRoute', 'ngDraggable'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/tournaments/:id', {
        templateUrl: 'tournaments/views/tournamentDetails/tournament-details.html',
        controller: 'TournamentsEditCtrl'
    });
}])

.controller('PopupResultCtrl', function ($uibModalInstance, score, scopeParent) {
    var scope = this;
    this.copySuccess = false;
    this.results = score;
    this.copy = function () {
        if (window.getSelection) {
            if (window.getSelection().empty) {  // Chrome
                window.getSelection().empty();
            } else if (window.getSelection().removeAllRanges) {  // Firefox
                window.getSelection().removeAllRanges();
            }
        } else if (document.selection) {  // IE?
            document.selection.empty();
        }
        if (document.selection) {
            var range = document.body.createTextRange();
            range.moveToElementText(document.getElementById("results_bbcode"));
            range.select();
            document.execCommand("Copy");

        } else if (window.getSelection) {
            var range = document.createRange();
            range.selectNode(document.getElementById("results_bbcode"));
            window.getSelection().addRange(range);
            document.execCommand("Copy");
        }
        scope.copySuccess = true;
    };

    this.cancel = function () {
        $uibModalInstance.dismiss('cancel');
    };
})

.controller('PopupRoundCtrl', function ($uibModalInstance, round, scopeParent) {
    var scope = this;
    this.copySuccess = false;
    this.round = round;
    this.copy = function () {
        if (window.getSelection) {
            if (window.getSelection().empty) {  // Chrome
                window.getSelection().empty();
            } else if (window.getSelection().removeAllRanges) {  // Firefox
                window.getSelection().removeAllRanges();
            }
        } else if (document.selection) {  // IE?
            document.selection.empty();
        }
        if (document.selection) {
            var range = document.body.createTextRange();
            range.moveToElementText(document.getElementById("results_bbcode"));
            range.select();
            document.execCommand("Copy");

        } else if (window.getSelection) {
            var range = document.createRange();
            range.selectNode(document.getElementById("results_bbcode"));
            window.getSelection().addRange(range);
            document.execCommand("Copy");
        }
        scope.copySuccess = true;
    };

    this.cancel = function () {
        $uibModalInstance.dismiss('cancel');
    };
})

.controller('TournamentsEditCtrl', ['$http', '$routeParams', '$route', '$uibModal', function($http, $routeParams, $route, $uibModal) {
    var scope = this;
    scope.tournament = {};
    scope.player = {};
    scope.player.lists = ["",""];
    scope.table = {};
    scope.round = {};
    scope.score = [];
    $http.get('/api/tournaments/'+$routeParams.id).success(function(data){
        scope.tournament = data;
        scope.tournament.rounds.forEach(function(round){
            verifyRound(round.number);
        });

    });

    $http.get('/api/tournaments/'+$routeParams.id+ '/results').success(function(data){
        scope.score = data;
    });

    this.bbCode = function (score) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: 'tournaments/views/tournamentDetails/tournament/result_popup.html',
            controller: 'PopupResultCtrl',
            controllerAs: 'PopupResultCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                score: function () {
                    return score;
                },
                scopeParent: function(){
                    return scope;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };

    this.bbCodeRound = function (round) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: 'tournaments/views/tournamentDetails/rounds/result_popup.html',
            controller: 'PopupRoundCtrl',
            controllerAs: 'PopupRoundCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                round: function () {
                    return round;
                },
                scopeParent: function(){
                    return scope;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };

    this.getNextRound = function(){
        scope.roundLoading = true;
        $http.get('/api/tournaments/'+scope.tournament.id+'/round').success(function(data){
            scope.round = data;
            scope.tournament.rounds[data.number] = data;
            verifyRound(data.number);
            scope.updateTournament();
            scope.roundLoading = false;
        })
        .fail(function(){
            scope.roundLoading = false;
        });
    };

    this.addPlayer = function(){
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.players.push(scope.player);
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
            scope.tournament.players.push(scope.player);
            scope.player = {};
            scope.player.lists = ["",""];
        });
    };

    this.addTable = function(){
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.tables.push(scope.table);
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
            scope.tournament.tables.push(scope.table);
            scope.table = {};
        });
    };

    this.deletePlayer = function(player){
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.players.splice(temp.players.indexOf(player), 1);
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
            scope.tournament.players.splice(scope.tournament.players.indexOf(player), 1);
        });
    };


    this.dropPlayer = function(player){
        player.leave = true;
        $http.put('/api/tournaments/'+scope.tournament.id, scope.tournament).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
        });
    };

    this.rejoinPlayer = function(player){
        player.leave = false;
        $http.put('/api/tournaments/'+scope.tournament.id, scope.tournament).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
        });
    };


    this.deleteTable = function(table){
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.tables.splice(temp.tables.indexOf(table), 1);
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
            scope.tournament.tables.splice(scope.tournament.tables.indexOf(table), 1);
        });
    };

    this.deleteRound = function(round){
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.rounds.splice(temp.rounds.indexOf(round), 1);
        $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
            scope.tournament.rounds.splice(scope.tournament.rounds.indexOf(round), 1);
        });
    };

    this.updateTournament = function(){
        $http.put('/api/tournaments/'+scope.tournament.id, scope.tournament).success(function(data){
            scope.tournament.id = data
            $route.updateParams({id:data});
        });
    };


    this.onDropComplete=function(source, destination, roundIndex){

        var sourceTemp = JSON.parse(JSON.stringify(source));

        source.name = destination.name;
        source.faction = destination.faction;
        source.payed_fee = destination.payed_fee;
        source.lists = destination.lists;
        source.leave = destination.leave;

        destination.name = sourceTemp.name;
        destination.faction = sourceTemp.faction;
        destination.payed_fee = sourceTemp.payed_fee;
        destination.lists = sourceTemp.lists;
        destination.leave = sourceTemp.leave;

        verifyRound(roundIndex);

    }

    function verifyRound(index){
        scope.tournament.rounds[index].games.forEach(function(game){
            verifyParing(game, index);
            verifyTable(game, index);
            verifyList(game, index);
        });
    }

    function verifyParing(g, index){
        g.errorPairing = false;
        scope.tournament.rounds.forEach(function(round){
            if (round.number===index){
                return;
            }
            round.games.forEach(function(game){
                if ((g.results[0].player.name === game.results[0].player.name && g.results[1].player.name === game.results[1].player.name) ||
                (g.results[0].player.name === game.results[1].player.name && g.results[1].player.name === game.results[0].player.name)) {
                    g.errorPairing = true;
                }
            });
        });
    }

    function verifyList(g, index){
        g.results[0].listFree = isListFree(g.results[0].player);
        g.results[1].listFree = isListFree(g.results[1].player);
    }

    function isListFree(player){
        //console.error("isFree ",player.name);
        var played = [];
        scope.tournament.rounds.forEach(function(round){
            round.games.forEach(function(game){
                if (game.results[0].player.name === player.name && game.results[0].list != "") {
                    console.error(player.name, "found game", game.results[0].list);
                    console.error("logged ? ", played.indexOf(game.results[0].list));
                    if (played.indexOf(game.results[0].list) === -1){
                        played.push(game.results[0].list);
                    }
                } else if (game.results[1].player.name === player.name && game.results[1].list != "") {
                    console.error(player.name, "found game", game.results[1].list);
                    console.error("logged ? ", played.indexOf(game.results[1].list));
                    if (played.indexOf(game.results[1].list) === -1){
                        played.push(game.results[1].list);
                    }
                }
            });
        });
        console.error(player.name, " ", played);
        console.error(played.length === player.lists.length);
        return played.length === player.lists.length;
    }

    function verifyTable(g, index){
        scope.tournament.rounds.forEach(function(round){
            if (round.number===index){
                return;
            }
            round.games.forEach(function(game){
                if (g.results[0].player.name === game.results[0].player.name || g.results[0].player.name === game.results[1].player.name) {
                    if (g.table.name === game.table.name) {
                        g.results[0].errorTable = true;
                    } else if (g.table.scenario === game.table.scenario){
                        g.results[0].errorScenario = true;
                    }
                } else if (g.results[1].player.name === game.results[0].player.name || g.results[1].player.name === game.results[1].player.name) {
                    if (g.table.name === game.table.name) {
                        g.results[1].errorTable = true;
                    } else if (g.table.scenario === game.table.scenario){
                        g.results[1].errorScenario = true;
                    }
                }
            });
        });
    }

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
