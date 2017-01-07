'use strict';

app.directive('tableList', function(){
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/tables/table-list.html",
        scope: {},
        controller: 'TablesCtrl',
        controllerAs: 'TablesCtrl',
        bindToController: {
            tournament: '=tournament'
        }
    };
});

app.directive('addTable', function(){
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/tables/table-add.html"
    };
});

app.directive('editTable', function(){
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/tables/table-edit.html"
    };
});
