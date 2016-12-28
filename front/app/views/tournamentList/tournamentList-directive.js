'use strict';

app.directive("futureTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "/views/tournamentList/tournament-future.html"
    };
});

app.directive("pastTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "/views/tournamentList/tournament-past.html"

    };
});
