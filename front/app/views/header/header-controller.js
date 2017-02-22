'use strict';

app.controller('HeaderCtrl', ["$localStorage", "$location", "$http", "$scope", "AuthService", function ($localStorage, $location, $http, $scope, authService) {

    var scope = this;

    if ($localStorage.currentUser != null) {
        $http.defaults.headers.common.Authorization = $localStorage.currentUser.token;
    }

    $scope.$watch(function () { return $localStorage.currentUser; },function(newVal,oldVal){
        if(newVal != null){
            scope.user = newVal.username;
        }
    })

    this.logout = function(){
        scope.error = null;
        $localStorage.currentUser = null;
        this.user = null;
        $http.defaults.headers.common.Authorization = null;
        $location.path( "/auth/login" );
    };
}]);
