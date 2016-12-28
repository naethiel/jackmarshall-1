'use strict';

app.service('AuthService', function($http, $localStorage){
	return {
		login : function(username, password){
            return $http.post(auth_endpoint + '/login', { login: username, password: password })
            .then(function(res) {
                $localStorage.currentUser = {username: username, password: password, token: res.data};
                $http.defaults.headers.common.Authorization = 'Bearer ' + res.data;
                return;
            })
            .catch(function (err){
                console.error("Unable to login : ", err);
                throw err.status
            });
		}
	}
});
