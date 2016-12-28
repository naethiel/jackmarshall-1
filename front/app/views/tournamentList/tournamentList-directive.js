'use strict';

app.directive("createTournament", function(){
    return {
        restrict: 'E',
        templateUrl: "/views/tournamentList/tournament-create.html"
    };
});


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
