'use strict';

app.service('UtilsService', function($http){
	return {
		getFileData : function(path){
			return $http.get(path)
            .then(function(res) {
				return res.data;
            })
            .catch(function (err){
                console.error("Unable to get data from file : ", err);
                throw err
            });
		}
	};
});
