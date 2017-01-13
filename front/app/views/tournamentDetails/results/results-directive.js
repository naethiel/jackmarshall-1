'use strict';

app.directive('results', function(){
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/results/results.html",
        scope: {},
        controller: 'ResultsCtrl',
        controllerAs: 'ResultsCtrl',
        bindToController: {
            tournament: '=tournament'
        }
    };
});
