'use strict';

app.controller('LoginCtrl', ["$localStorage", "$http", "$location", "AuthService", function($localStorage, $http, $location, authService) {

    if($localStorage.currentUser != null){
        $location.path( "/tournament/list" );
    }

    var scope = this;
    scope.username = "";
    scope.password = "";
    scope.error = undefined;

    scope.newUser = {};

    this.login = function(){
        scope.error = null;
        authService.login(scope.username, scope.password).then(function(){
            $location.path( "/tournament/list" );
        }).catch(function(err){
            scope.error = err
        })
    };

    this.logout = function(){
        scope.error = null;
        $localStorage.currentUser = null;
        $http.defaults.headers.common.Authorization = null;
    };

    this.toNewAccount = function(){
        $location.path( "/auth/new" );
    };

}]);
