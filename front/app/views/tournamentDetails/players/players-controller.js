'use strict';

app.controller('PlayersCtrl', ["$rootScope", "$route", "uuid", "TournamentService", "UtilsService", function ($rootScope, $route, uuid, tournamentService, utilsService) {
    var lists = JSON.stringify([{
        caster: "",
        theme: "",
        played: false,
        list: ""
    },{
        caster: "",
        theme: "",
        played: false,
        list: ""
    }]);
    var scope = this;
    scope.tournament = {};
    scope.player = {};
    scope.player.lists = JSON.parse(lists);
    scope.errorAdd = undefined;
    scope.errorUpdate = undefined;
    scope.errorDelete = undefined;
    scope.playersCollapsed = false;
    scope.casters = []

    utilsService.getCasters().then(function(casters){
        scope.casters = casters;
    })

    this.addPlayer = function(){
        scope.errorAdd = null;
        scope.player.id = uuid.v4();
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.players.push(scope.player);
        tournamentService.update(temp).then(function(id){
            scope.tournament.players.push(scope.player);
            scope.player = {};
            scope.player.lists = JSON.parse(lists);
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            scope.errorAdd = true;
        })
    };

    this.deletePlayer = function(player){
        scope.errorDelete = null;
        var temp = JSON.parse(JSON.stringify(scope.tournament));
        temp.players.splice(temp.players.indexOf(player), 1);
        tournamentService.update(temp).then(function(id){
            scope.tournament.players.splice(scope.tournament.players.indexOf(player), 1);
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            scope.errorDelete = true;
        })
    };

    this.updatePlayer = function(player){
        scope.errorUpdate = null;
        scope.tournament.rounds.forEach(function(round){
            round.games.forEach(function(game){
                game.results.forEach(function(result){
        			if (result.player.id === player.id){
                        result.player = player;
                    }
        		});
    		});
		});
        tournamentService.update(scope.tournament).then(function(id){
            $rootScope.$emit("UpdateResult");
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };
    //FIXME then empty
    this.dropPlayer = function(player){
        player.leave = true;
        tournamentService.update(scope.tournament).then(function(id){
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };
    //FIXME then empty

    this.rejoinPlayer = function(player){
        player.leave = false;
        tournamentService.update(scope.tournament).then(function(id){
        }).catch(function(err){
            scope.errorUpdate = true;
        })
    };

}]);
