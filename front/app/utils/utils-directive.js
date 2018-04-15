'use strict';

app.directive("dateFormat", function(){
    return {
        restrict: 'A',
        require: 'ngModel',
        link: function(scope, elem, attrs, ctrl){
            var dateFormat = attrs.dateFormat;
            attrs.$observe('dateFormat', function (newValue) {
                if (dateFormat == newValue || !ctrl.$modelValue) return;
                dateFormat = newValue;
                ctrl.$modelValue = new Date(ctrl.$setViewValue);
            });

            ctrl.$formatters.unshift(function (modelValue) {
                scope = scope;
                if (!dateFormat || !modelValue) return "";
                var retVal = moment(modelValue).format(dateFormat);
                return retVal;
            });

            ctrl.$parsers.unshift(function (viewValue) {
                scope = scope;
                var date = moment(viewValue, dateFormat);
                return (date && date.isValid()) ? date.toDate() : "";
            });
        }
    };
});

app.directive("timeFormat", ['moment', function(moment){
    return {
        restrict: 'A',
        require: 'ngModel',
        link: function(scope, elem, attrs, ctrl){
            var timeFormat = attrs.timeFormat;
            attrs.$observe('timeFormat', function (newValue) {
                if (timeFormat == newValue || !ctrl.$modelValue) return;
                timeFormat = newValue;
                ctrl.$modelValue = ctrl.$setViewValue;
            });

            ctrl.$formatters.unshift(function (modelValue) {
                scope = scope;
                if (!timeFormat || !modelValue) return "";
                var retVal = moment(modelValue).format(timeFormat);
                return retVal;
            });

            ctrl.$parsers.unshift(function (viewValue) {
                scope = scope;
                var date = moment(viewValue, timeFormat);
                return (date && date.isValid()) ? date : "";
            });
        }
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
