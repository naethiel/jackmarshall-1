'use strict';

app.controller('AuthCtrl', ["$localStorage", "$http", "$location", "AuthService", function($localStorage, $http, $location, authService) {
    var scope = this;
    scope.username = "";
    scope.password = "";
    scope.error = undefined;

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
}]);
