'use strict';

app.controller('LoginCtrl', ["$localStorage", "$http", "$location", "AuthService", function($localStorage, $http, $location, authService) {

    if($localStorage.currentUser != null){
        $location.path( "/tournament/list" );
    }

    var scope = this;
    scope.username = "";
    scope.password = "";
    scope.error = null;

    scope.login = function(){
        scope.error = null;
        authService.login(scope.username, scope.password).then(function(){
            $location.path( "/tournament/list");
        }).catch(function(err){
            scope.error = err;
        })
    };
}]);
