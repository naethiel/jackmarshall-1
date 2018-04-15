'use strict';

app.controller('NewUserCtrl', ["$localStorage", "$http", "$location", "AuthService", function($localStorage, $http, $location, authService) {

    if($localStorage.currentUser != null){
        $location.path( "/tournament/list" );
    }

    var scope = this;
    scope.user = {};
    scope.password_confirm = "";
    scope.error = null;

    scope.create = function(){
        scope.error = null;
        authService.create(scope.user).then(function(){
            authService.login(scope.user.name, scope.user.password).then(function(){
                $location.path("/tournament/list");
            })
        }).catch(function(err){
            scope.error = err
        })
    };
}]);
