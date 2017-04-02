'use strict';

app.service('AuthService', function($http, $localStorage, $q, $location){
	return {
		login : function(username, password){
            return $http.post(auth_endpoint + '/login', { login: username, password: password })
            .then(function(res) {
                $localStorage.currentUser = {username: username, token: res.data.token, expiration: res.data.expiration, refresh_token: res.data.refresh_token};
                $http.defaults.headers.common.Authorization = res.data.token;
                return;
            })
            .catch(function (err){
                console.error("Unable to login : ", err);
                throw err.status
            });
		},
		logout : function(){
	        $localStorage.currentUser = null;
	        $http.defaults.headers.common.Authorization = null;
	        $location.path( "/auth/login" );
		},
		create : function(user){
            return $http.post(auth_endpoint + '/organizer', user)
            .then(function(res) {
                $localStorage.currentUser = {username: user.login, token: res.data.token, expiration: res.data.expiration, refresh_token: res.data.refresh_token};
                $http.defaults.headers.common.Authorization = res.data.token;
                return;
            })
            .catch(function (err){
                console.error("Unable to create user : ", err);
                throw err.status
            });
		},

		refresh : function() {
          if ($localStorage.currentUser.token && $localStorage.currentUser.expiration > Math.floor( Date.now() / 1000 )) {
			  return $q.when();
		  } else {
			  return $http.post(auth_endpoint + '/refresh', $localStorage.currentUser.refresh_token)
              .then(function(res) {
                  $localStorage.currentUser.token = res.data.token;
                  $localStorage.currentUser.expiration = res.data.expiration;
                  $http.defaults.headers.common.Authorization = res.data.token;
                  return;
              })
              .catch(function (err){
                  console.error("Unable to refresh token : ", err);
				  $localStorage.currentUser = null;
				  $http.defaults.headers.common.Authorization = null;
				  $location.path( "/auth/login" );
				  throw err.status
              });
            }
    	}
	};
});
