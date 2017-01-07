'use strict';

app.controller('TablesCtrl', ["$route", "uuid", "TournamentService", function ($route, uuid, tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.table = {};
    scope.errorAdd = undefined;
    scope.errorUpdate = undefined;
    scope.errorDelete = undefined;
    scope.tablesCollapsed = false;

    this.addTable = function(){
        scope.errorAdd = null;
        scope.table.id = uuid.v4();
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.tables.push(scope.table);
        tournamentService.update(temp).then(function(id){
            scope.tournament.id = id;
            $route.updateParams({id:id});
            scope.tournament.tables.push(scope.table);
            scope.table = {};
        }).catch(function(err){
            scope.errorAdd = true;
        })
    };

    this.deleteTable = function(table){
        console.log("TOTO");
        scope.errorDelete = null;
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.tables.splice(temp.tables.indexOf(table), 1);
        tournamentService.update(temp).then(function(id){
            scope.tournament.id = id
            $route.updateParams({id:id});
            scope.tournament.tables.splice(scope.tournament.tables.indexOf(table), 1);
        }).catch(function(err){
            scope.errorDelete = true;
        })
    };

    this.updateTable = function(){
        scope.errorUpdate = null;
        tournamentService.update(scope.tournament).then(function(id){
            scope.tournament.id = id
            $route.updateParams({id:id});
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };

}]);
