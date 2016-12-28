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
		}
	};
});
