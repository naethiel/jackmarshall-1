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
		}
	};
});
