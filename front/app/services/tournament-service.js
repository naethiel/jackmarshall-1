'use strict';

app.service('TournamentService', ['$http', 'AuthService', function($http, authService){
	return {
		create : function(tournament){
			return authService.refresh().then(function(){
				return $http.post(tournament_endpoint + '/tournaments', tournament)
				.then(function(res) {
					return res.data;
				})
				.catch(function (err){
					console.error("Unable to create tournament : ", err);
					throw err.status;
				});
			});
		},
		getAll : function(){
			return authService.refresh().then(function(){
				return $http.get(tournament_endpoint + '/tournaments')
				.then(function(res) {
					return res.data;
				})
				.catch(function (err){
					console.error("Unable to get tournaments list : ", err);
					throw err.status;
				});
			});
		},
		delete : function(id){
			return authService.refresh().then(function(){
				return $http.delete(tournament_endpoint + '/tournaments/' + id)
				.then(function() {
					return;
				})
				.catch(function (err){
					console.error("Unable to delete tournament " + id + " : ", err);
					throw err.status;
				});
			});
		},
		get : function(id){
			return authService.refresh().then(function(){
				return $http.get(tournament_endpoint + '/tournaments/' + id)
				.then(function(res) {
					return res.data;
				})
				.catch(function (err){
					console.error("Unable to get tournament " + id + " : ", err);
					throw err.status;
				});
			});
		},
		update : function(tournament){
			return authService.refresh().then(function(){
				return $http.put(tournament_endpoint + '/tournaments/' + tournament.id, tournament)
				.then(function(res) {
					return res.data;
				})
				.catch(function (err){
					console.error("Unable to update tournament " + tournament.id + " : ", err);
					throw err.status;
				});
			});
		},
		getResults : function(id){
			return authService.refresh().then(function(){
				return $http.get(tournament_endpoint + '/tournaments/' + id + "/results")
				.then(function(res) {
					return res.data;
				})
				.catch(function (err){
					console.error("Unable to get results for tournament " + id + " : ", err);
					throw err.status;
				});
			});
		},
		getNextRound : function(id){
			return authService.refresh().then(function(){
				return $http.get(tournament_endpoint + '/tournaments/' + id + "/round")
				.then(function(res) {
					return res.data;
				})
				.catch(function (err){
					console.error("Unable to get results for tournament " + id + " : ", err);
					throw err.status;
				});
			});
		},
		verifyRound : function(tournament, roundNumber){
			for (var k in tournament.players) {
				tournament.players[k].victory_points=0;
			}
			var index;
			for (var i=0; i<tournament.rounds.length; i++){
				if (tournament.rounds[i].number === roundNumber){
					index = i;
					break;
				}
			}
			verifyVictoryPoint(tournament, index);

			tournament.rounds[index].games.forEach(function(game){
				verifyParing(tournament, game, index);
				// verifyTable(tournament, game, index);
				verifyOrigin(tournament,game,index);
			});
		}
	};
}]);




function verifyParing(tournament, g, index){
	g.errorPairing = false;
	for (var i=0; i < index; i++){
		tournament.rounds[i].games.forEach(function(game){
			if ((g.results[0].player === game.results[0].player && g.results[1].player === game.results[1].player) ||
			(g.results[0].player === game.results[1].player && g.results[1].player === game.results[0].player)) {
				g.errorPairing = true;
			}
		});
	}
};

function verifyVictoryPoint(tournament, index){
	for (var i=0; i < index; i++){
		tournament.rounds[i].games.forEach(function(game){
			tournament.players[game.results[0].player].victory_points += game.results[0].victory_points
			tournament.players[game.results[1].player].victory_points += game.results[1].victory_points
		});
	}
	tournament.rounds[index].games.forEach(function(game){
		game.errorVP = false;
		if (tournament.players[game.results[0].player].victory_points != tournament.players[game.results[1].player].victory_points){
			game.errorVP = true;
		}
	});
};


function verifyOrigin(tournament,g, index){
	g.errorOrigin = (tournament.players[g.results[0].player].origin == tournament.players[g.results[1].player].origin && tournament.players[g.results[0].player].origin != "");
};

function getNum(val) {
	if (isNaN(val)) {
		return 0;
	}
	return val;
}

function verifyTable(tournament, g, index){
	g.results[0].errorTable = false;
	g.results[1].errorTable = false;
	g.results[0].errorScenario = false;
	g.results[1].errorScenario = false;
	for (var i=0; i < index; i++){
		tournament.rounds[i].games.forEach(function(game){
			if (g.results[0].player === game.results[0].player || g.results[0].player === game.results[1].player) {
				if (g.table === game.table) {
					g.results[0].errorTable = true;
				} else if (tournament.tables[g.table].scenario === tournament.tables[game.table].scenario){
					g.results[0].errorScenario = true;
				}
			} else if (g.results[1].player.id === game.results[0].player.id || g.results[1].player.id === game.results[1].player.id) {
				if (g.table.id === game.table.id) {
					g.results[1].errorTable = true;
				} else if (g.table.scenario === game.table.scenario){
					g.results[1].errorScenario = true;
				}
			}
		});
	}
};
