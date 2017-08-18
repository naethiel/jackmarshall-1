'use strict';

app.controller('TabsCtrl', ["$rootScope", "$route", "TournamentService", function ($rootScope, $route, tournamentService) {
    var scope = this;
    scope.tab = -1;
    scope.error = undefined

    $rootScope.$on("SetTab", function(event, tab){
        scope.tab = tab;
    });

    this.isSet = function(tab) {
        return scope.tab === tab;
    };

    this.setTab = function(tab) {
        scope.tab = tab;
    };

    this.getNextRound = function(){
        scope.roundLoading = true;
        tournamentService.update(scope.tournament).then(function(){
            tournamentService.getNextRound(scope.tournament.id).then(function(tournament){
                scope.tournament = tournament;
                scope.round = scope.tournament.rounds[scope.tournament.rounds.length - 1];
                scope.tab = scope.tournament.rounds.length - 1;
                tournamentService.verifyRound(scope.tournament, scope.round.number);
                scope.roundLoading = false;
                $rootScope.$emit("UpdateRounds", scope.tournament.rounds.length);
            }).catch(function(err){
                scope.roundLoading = false;
            })
        }).catch(function(err){
            scope.roundLoading = false;
        })
    };
}]);
