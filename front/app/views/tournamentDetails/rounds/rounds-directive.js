'use strict';

app.directive("editRound", function(){
    return {
        restrict: "E",
        templateUrl: "views/tournamentDetails/rounds/round-edit.html",
        scope: {},
        controller: 'RoundsCtrl',
        controllerAs: 'RoundsCtrl',
        bindToController: {
            tournament: '=tournament',
            round: '=round'
        }
    };
});
