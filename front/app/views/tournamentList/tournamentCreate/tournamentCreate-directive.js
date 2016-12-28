'use strict';

app.directive("createTournament", function(){
    return {
        restrict: 'E',
        templateUrl: "/views/tournamentList/tournamentCreate/tournament-create.html",
        scope: {},
        controller: 'CreateTournamentCtrl',
        controllerAs: 'CreateCtrl',
        bindToController: {
            tournaments: '=tournaments'
        }
    };
});
