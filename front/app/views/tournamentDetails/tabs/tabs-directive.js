'use strict';

app.directive("tabs", function() {
    return {
        restrict: "E",
        templateUrl: "/views/tournamentDetails/tabs/tab.html",
        scope: {},
        controller: 'TabsCtrl',
        controllerAs: 'TabsCtrl',
        bindToController: {
            tournament: '=tournament'
        }
    };
});
