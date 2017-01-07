'use strict';

app.service('UtilsService', function($http){
	return {
		getCasters : function(){
			return $http.get('/data/casters.json')
            .then(function(res) {
				return res.data;
            })
            .catch(function (err){
                console.error("Unable to get casters list : ", err);
                throw err
            });
		}
	};
});
