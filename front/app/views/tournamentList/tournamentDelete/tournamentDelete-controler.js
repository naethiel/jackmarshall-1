'use strict';

app.controller('DeleteTournamentCtrl', function ($uibModalInstance, tournament, scopeParent, tournamentService) {
    var scope = this;
    scope.error = undefined
    this.ok = function () {
        scope.error = null;
        tournamentService.delete(tournament.id).then(function(){
            scopeParent.tournaments.splice(scopeParent.tournaments.indexOf(tournament), 1);
            $uibModalInstance.close();
        }).catch(function(){
            scope.error = true;
        });
    };
    this.cancel = function () {
        $uibModalInstance.dismiss('cancel');
    };
});
