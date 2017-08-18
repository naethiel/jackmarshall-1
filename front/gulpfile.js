var gulp = require('gulp'),
    sass = require('gulp-sass'),
    postcss = require('gulp-postcss'),
    autoprefixer = require('gulp-autoprefixer'),
    plumber = require('gulp-plumber'),
    uglify = require('gulp-uglify'),
    concat = require('gulp-concat');
    uncss = require('gulp-uncss');

var path = {
	'vendors' : './bower_components/',
	'app' : './app/',
	'images' : './images/',
	'data' : './data/',
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
	// .pipe(uglify())
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('config', function(){
	gulp.src([
		path.app + 'config.js',
	])
    .pipe(plumber())
    .pipe(gulp.dest('./dist/js/'));
})

gulp.task('images', function(){
	gulp.src([
		path.images + '**/*',
	])
    .pipe(plumber())
    .pipe(gulp.dest('./dist/images/'));
})

gulp.task('data', function(){
	gulp.src([
		path.data + '**/*',
	])
    .pipe(plumber())
    .pipe(gulp.dest('./dist/data/'));
})

gulp.task('jquery', function(){
	gulp.src([
		path.vendors + 'jquery/dist/jquery.min.js',
	])
    .pipe(plumber())
    .pipe(gulp.dest('./dist/js/'));
})

gulp.task('timerDeps', function(){
	gulp.src([
		path.vendors + 'jquery.countdown/dist/jquery.countdown.js',
		path.vendors + 'moment/min/moment.min.js',
	])
    .pipe(plumber())
    .pipe(concat('timer.vendors.js'))
    .pipe(uglify())
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('app', function(){
	gulp.src([
		path.app + '**/*.js',
        "!" + path.app + "config.js",
	])
    .pipe(plumber())
	.pipe(concat('app.min.js'))
	.pipe(uglify())
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('views', function(){
	gulp.src([
		path.app + '**/*.html',
	])
    .pipe(plumber())
	.pipe(gulp.dest('./dist/'));
})

gulp.task('style', function(){
	gulp.src([
		path.style + 'jm.scss',
	])
    .pipe(plumber())
    .pipe(sass())
    .pipe(autoprefixer())
    .pipe(postcss())
    // .pipe(uncss({html: ['app/*index.html', 'app/**/*.html']}))
	.pipe(gulp.dest('./dist/style/'));
})

gulp.task('fonts', function(){
	gulp.src([
        path.vendors + 'bootstrap/fonts/*',
        path.vendors + 'fontawesome/fonts/*',
	])
    .pipe(plumber())
	.pipe(gulp.dest('./dist/fonts/'));
})

gulp.task('app-dev', function(){
	gulp.src([
		path.app + '**/*.js',
        "!" + path.app + "config.js",
	])
    .pipe(plumber())
	.pipe(concat('app.js'))
	.pipe(gulp.dest('./dist/js/'));
})

gulp.task('watch', function () {
	gulp.watch(path.app + '**/*.js', ['app', 'app-dev', 'config']);
	gulp.watch(path.app + '**/*.html', ['views']);
	gulp.watch(path.data + '**/*', ['data']);
	gulp.watch(path.style + '**/*.css', ['style']);
	gulp.watch(path.style + '**/*.scss', ['style']);
});

gulp.task('default', ['vendors','app', 'app-dev', 'views', 'style', 'fonts', 'jquery', 'timerDeps', 'config', 'data', 'images']);
