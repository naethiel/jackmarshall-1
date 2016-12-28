'use strict';

app.directive("futureTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "/views/tournamentList/tournament-future.html"
    };
});

app.directive("pastTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "/views/tournamentList/tournament-past.html"

    };
});

app.filter('isFuture', function() {
    return function(items, dateFieldName) {
        return items.filter(function(item){
            return moment(item[dateFieldName || 'date']).isSameOrAfter(new Date(),'day');
        })
    }
});

app.filter('isPast', function() {
    return function(items, dateFieldName) {
        return items.filter(function(item){
            return moment(item[dateFieldName || 'date']).isBefore(new Date(), 'day');
        })
    }
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
