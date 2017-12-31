'use strict';

app.filter('filterGame', function() {
    return function(items, str, players, tables) {
        return items.filter(function(item){
            var table = tables[item.table].name
            var scenario = tables[item.table].scenario
            var p0 = players[item.results[0].player].name
            var p1 = players[item.results[1].player].name
            return (str==null||table.includes(str)||scenario.includes(str)||p0.includes(str)||p1.includes(str));
        })
    }
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
