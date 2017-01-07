'use strict';

app.service('TournamentService', function($http){
	return {
		create : function(tournament){
			return $http.post(tournament_endpoint + '/tournaments', tournament)
			.then(function(res) {
				return res.data;
			})
			.catch(function (err){
				console.error("Unable to create tournament : ", err);
				throw err.status;
			});
		},
		getAll : function(){
			return $http.get(tournament_endpoint + '/tournaments')
			.then(function(res) {
				return res.data;
			})
			.catch(function (err){
				console.error("Unable to get tournaments list : ", err);
				throw err.status;
			});
		},
		delete : function(id){
			return $http.delete(tournament_endpoint + '/tournaments/' + id)
			.then(function() {
				return;
			})
			.catch(function (err){
				console.error("Unable to delete tournament " + id + " : ", err);
				throw err.status;
			});
		},
		get : function(id){
			return $http.get(tournament_endpoint + '/tournaments/' + id)
			.then(function(res) {
				return res.data;
			})
			.catch(function (err){
				console.error("Unable to get tournament " + id + " : ", err);
				throw err.status;
			});
		},
		update : function(tournament){
			return $http.put(tournament_endpoint + '/tournaments/' + tournament.id, tournament)
			.then(function(res) {
				return res.data;
			})
			.catch(function (err){
				console.error("Unable to update tournament " + id + " : ", err);
				throw err.status;
			});
		},
		getResults : function(id){
			return $http.get(tournament_endpoint + '/tournaments/' + id + "/results")
			.then(function(res) {
				return res.data;
			})
			.catch(function (err){
				console.error("Unable to get results for tournament " + id + " : ", err);
				throw err.status;
			});
		},
		getNextRound : function(id){
			return $http.get(tournament_endpoint + '/tournaments/' + id + "/round")
			.then(function(res) {
				return res.data;
			})
			.catch(function (err){
				console.error("Unable to get results for tournament " + id + " : ", err);
				throw err.status;
			});
		},
		verifyRound : function(tournament, roundNumber){
			var index;
			for (var i=0; i<tournament.rounds.length; i++){
				if (tournament.rounds[i].number === roundNumber){
					index = i;
					break;
				}
			}
			tournament.rounds[index].games.forEach(function(game){
				verifyParing(tournament, game, index);
				verifyTable(tournament, game, index);
				verifyList(tournament, game.results[0].player, index);
				verifyList(tournament, game.results[1].player, index);
			});
		}
	};
});




function verifyParing(tournament, g, index){
	g.errorPairing = false;
	for (var i=0; i < index; i++){
		var round = tournament.rounds[i];
		if (round.number===index){
			return;
		}
		round.games.forEach(function(game){
			if ((g.results[0].player.name === game.results[0].player.name && g.results[1].player.name === game.results[1].player.name) ||
			(g.results[0].player.name === game.results[1].player.name && g.results[1].player.name === game.results[0].player.name)) {
				g.errorPairing = true;
			}
		});
	}
};

function verifyList(tournament, player, index){
	for (var i=0; i < index; i++){
		var round = tournament.rounds[i];
		round.games.forEach(function(game){
			if (game.results[0].player.name === player.name && game.results[0].list != "") {
				player.lists.forEach(function(list) {
					if (list.caster === game.results[0].list){
						list.played = true;
					}
				});

			} else if (game.results[1].player.name === player.name && game.results[1].list != "") {
				player.lists.forEach(function(list) {
					if (list.caster === game.results[1].list){
						list.played = true;
					}
				});
			}
		});
	}
};

function verifyTable(tournament, g, index){

	for (var i=0; i < index; i++){
		var round = tournament.rounds[i];
		if (round.number===index){
			return;
		}
		round.games.forEach(function(game){
			if (g.results[0].player.name === game.results[0].player.name || g.results[0].player.name === game.results[1].player.name) {
				if (g.table.name === game.table.name) {
					g.results[0].errorTable = true;
				} else if (g.table.scenario === game.table.scenario){
					g.results[0].errorScenario = true;
				}
			} else if (g.results[1].player.name === game.results[0].player.name || g.results[1].player.name === game.results[1].player.name) {
				if (g.table.name === game.table.name) {
					g.results[1].errorTable = true;
				} else if (g.table.scenario === game.table.scenario){
					g.results[1].errorScenario = true;
				}
			}
		});
	}
};
