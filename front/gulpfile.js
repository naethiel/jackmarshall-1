var gulp = require('gulp'),
    uglify = require('gulp-uglify'),
    concat = require('gulp-concat');

var path = {
	'vendors' : './bower_components/',
	'app' : './app/',
    'style' : './style/'
};

gulp.task('vendors', function(){
	gulp.src([
		path.vendors + 'angular/angular.min.js',
		path.vendors + 'angular-route/angular-route.min.js',
		path.vendors + 'angular-animate/angular-animate.min.js',
		path.vendors + 'angular-bootstrap/ui-bootstrap-tpls.min.js',
		path.vendors + 'moment/min/moment.min.js',
		path.vendors + 'ngDraggable/ngDraggable.js',
		path.vendors + 'ngstorage/ngStorage.min.js',
		path.vendors + 'angular-uuids/angular-uuid.js',
	])
	.pipe(concat('vendors.js'))
	.pipe(uglify())
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('jquery', function(){
	gulp.src([
		path.vendors + 'jquery/dist/jquery.min.js',
	])
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('app', function(){
	gulp.src([
		path.app + '**/*.js',
	])
	.pipe(concat('app.min.js'))
	.pipe(uglify())
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('views', function(){
	gulp.src([
		path.app + '**/*.html',
	])
	.pipe(gulp.dest('./dist/'));
})

gulp.task('style', function(){
	gulp.src([
		path.style + 'jm.css',
        path.vendors + 'bootstrap/dist/css/bootstrap.min.css',
        path.vendors + 'fontawesome/css/font-awesome.min.css',
	])
	.pipe(gulp.dest('./dist/style/'));
})

gulp.task('fonts', function(){
	gulp.src([
        path.vendors + 'bootstrap/fonts/*',
        path.vendors + 'fontawesome/fonts/*',
	])
	.pipe(gulp.dest('./dist/fonts/'));
})

gulp.task('app-dev', function(){
	gulp.src([
		path.app + '**/*.js',
	])
	.pipe(concat('app.js'))
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('watch', function () {
	gulp.watch(path.app + '**/*.js', ['app', 'app-dev']);
	gulp.watch(path.app + '**/*.html', ['views']);
	gulp.watch(path.style + '**/*.css', ['style']);
});

gulp.task('default', ['vendors','app', 'app-dev', 'views', 'style', 'fonts', 'jquery']);
