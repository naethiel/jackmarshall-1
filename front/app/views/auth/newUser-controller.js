'use strict';

app.controller('NewUserCtrl', ["$localStorage", "$http", "$location", "AuthService", function($localStorage, $http, $location, authService) {

    if($localStorage.currentUser != null){
        $location.path( "/tournament/list" );
    }

    var scope = this;
    scope.user = {};
    scope.error = undefined;

    this.create = function(){
        scope.error = null;
        authService.create(scope.user).then(function(){
            authService.login(scope.user.login, scope.user.password).then(function(){
                $location.path( "/tournament/list" );
            })
        }).catch(function(err){
            scope.error = err
        })
    };

    this.toLogin = function(){
        $location.path( "/auth/login" );
    };

}]);

app.directive('equalsTo', [function () {
       return {
           restrict: 'A',
           scope: true,
           require: 'ngModel',
           link: function (scope, elem, attrs, control) {
               var check = function () {
               var v1 = scope.$eval(attrs.ngModel);
               var v2 = scope.$eval(attrs.equalsTo).$viewValue;
               return v1 == v2;
           };
           scope.$watch(check, function (isValid) {
               control.$setValidity("equalsTo", isValid);
           });
       }
   };
}]);
