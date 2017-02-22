'use strict';

app.directive("jmHeader", function(){
    return {
        restrict: "E",
        templateUrl: "/views/header/header.html",
        scope: {},
        controller: 'HeaderCtrl',
        controllerAs: 'HeaderCtrl'
    };
});
