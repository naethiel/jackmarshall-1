'use strict';

app.controller('PlayersCtrl', ["TournamentService", function (tournamentService) {
    var scope = this;
    scope.tournament = {};
    scope.player = {};
    scope.error = undefined;
    scope.playersCollapsed = false;

    // this.addPlayer = function(){
    //     scope.player.id = uuid.v4();
    //     var temp = JSON.parse(JSON.stringify(scope.tournament));
    //     temp.players.push(scope.player);
    //     $http.put('/api/tournaments/'+scope.tournament.id, temp).success(function(data){
    //         scope.tournament.id = data
    //         $route.updateParams({id:data});
    //         scope.tournament.players.push(scope.player);
    //         scope.player = {};
    //         scope.player.lists = ["",""];
    //     });
    // };
}]);
