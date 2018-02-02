'use strict';

app.controller('TimerCtrl', ['$interval', 'moment', function ($interval, moment) {
    var scope = this;
    scope.duration = moment('1970-01-01T02:15:00');
    scope.remaining = moment('1970-01-01T00:00:00');
    scope.end = 0;
    scope.pauseTime = undefined;
    scope.timer = {};
    scope.started=false;

    this.start = function(){
        if (scope.started) {
            return;
        }
        if (scope.pauseTime != null) {
            scope.end.add(moment().diff(scope.pauseTime, 'seconds'), 'seconds')
            scope.pauseTime = null;
        } else {
            scope.end = moment().add(scope.duration.hours(), 'h')
            .add(scope.duration.minutes(), 'm')
            .add(scope.duration.seconds(), 's');
        }
        this.tick();
        scope.timer = $interval(this.tick,1000);
        scope.started=true;
    };

    this.pause = function(){
        $interval.cancel(scope.timer);
        if (scope.pauseTime == null && scope.started){
            scope.pauseTime = moment();
            scope.started=false;
        }
    };

    this.stop = function(){
        $interval.cancel(scope.timer);
        scope.remaining = moment('1970-01-01T00:00:00');
        scope.pauseTime = null;
        scope.started=false;
    };

    this.tick = function() {
        var now=moment();
        if (scope.end.diff(now, 'seconds') <= 0) {
            $interval.cancel(scope.timer);
        }
        scope.remaining = moment().hours(scope.end.diff(now, 'hours')%24)
        .minutes(scope.end.diff(now, 'minutes')%60)
        .seconds(scope.end.diff(now, 'seconds')%60);
    };
}]);
