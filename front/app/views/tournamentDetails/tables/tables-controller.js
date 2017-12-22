'use strict';

app.controller('TablesCtrl', ["$scope","$rootScope", "$route", "uuid", "TournamentService", "UtilsService", function ($scope,$rootScope, $route, uuid, tournamentService, utilsService) {
    var scope = this;
    scope.tournament = {};
    scope.table = {};
    scope.errorAdd = undefined;
    scope.errorUpdate = undefined;
    scope.errorDelete = undefined;
    scope.tablesCollapsed = false;
    scope.scenarios = []

    utilsService.getFileData('/data/scenarios.json').then(function(scenarios){
        scope.scenarios = scenarios;
        console.log(scope.scenarios);
    })

    $rootScope.$on("UpdateRounds", function(e, nb_round){
        scope.tablesCollapsed = (nb_round > 0);
    });

    this.addTable = function(){
        scope.errorAdd = false;
        scope.table.id = uuid.v4();
        scope.tournament.tables[scope.table.id] = scope.table
        tournamentService.update(scope.tournament).then(function(){
            scope.table = {};
            document.getElementById("add_table_name").focus()
            $scope.addTableForm.$setUntouched();
        }).catch(function(err){
            delete scope.tournament.tables[scope.table.id]
            scope.errorAdd = true;
        })
    };

    this.deleteTable = function(table){
        scope.errorDelete = false;
        delete scope.tournament.tables[table.id]
        tournamentService.update(scope.tournament).then(function(){
        }).catch(function(err){
            scope.tournament.tables[table.id] = table
            scope.errorDelete = true;
        })
    };

    this.updateTable = function(table){
        scope.errorUpdate = false;
        tournamentService.update(scope.tournament).then(function(){
            table.detailsVisible=false;
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };

    this.compare = function(a, b) {
        return naturalSort(a.value, b.value);
    };

}]);
