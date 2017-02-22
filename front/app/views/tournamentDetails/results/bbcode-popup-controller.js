'use strict';

app.controller('ResultsBBCodeCtrl', function ($uibModalInstance, score, scopeParent) {
    var scope = this;
    this.copySuccess = false;
    this.results = score;
    this.copy = function () {
        if (window.getSelection) {
            if (window.getSelection().empty) {  // Chrome
                window.getSelection().empty();
            } else if (window.getSelection().removeAllRanges) {  // Firefox
                window.getSelection().removeAllRanges();
            }
        } else if (document.selection) {  // IE?
            document.selection.empty();
        }
        if (document.selection) {
            var range = document.body.createTextRange();
            range.moveToElementText(document.getElementById("results_bbcode"));
            range.select();
            document.execCommand("Copy");

        } else if (window.getSelection) {
            var range = document.createRange();
            range.selectNode(document.getElementById("results_bbcode"));
            window.getSelection().addRange(range);
            document.execCommand("Copy");
        }
        scope.copySuccess = true;
    };

    this.cancel = function () {
        $uibModalInstance.dismiss('cancel');
    };
});