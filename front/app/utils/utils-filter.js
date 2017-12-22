'use strict';

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

app.filter('trim', function () {
    return function(value) {
        if(!angular.isString(value)) {
            return value;
        }
        return value.replace(/ +/g, "").toLowerCase();
    };
});

app.filter('fdate', function() {
    return function(input, format) {
        if (!moment.isMoment(input)){
            return "invalid moment object";
        }
        return input.format(format);
    };
});

app.filter('toArray', function () {
    return function (obj) {
        if (!angular.isObject(obj)) return obj;
        return Object.keys(obj).map(function(key) {
            return obj[key];
        });
    };
});

app.filter('nbKeys', function() {
    return function(object) {
        if (object===undefined ||object===undefined){
            return null
        }
        return Object.keys(object).length;
    }
});
