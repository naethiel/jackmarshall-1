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
                throw "CreateTournamentError";
            });
		},
		getAll : function(){
            return $http.get(tournament_endpoint + '/tournaments')
            .then(function(res) {
                return res.data;
            })
            .catch(function (err){
                console.error("Unable to get tournaments list : ", err);
                throw "GetAllTournamentError";
            });
		}
	};
});
