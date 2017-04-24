'use strict';

app.controller('DeleteRoundCtrl', function ($uibModalInstance, $rootScope, tournament, round, scopeParent, tournamentService) {
    var scope = this;
    scope.error = undefined

    this.ok = function () {
        scope.error = null;
        var temp = JSON.parse(JSON.stringify(scopeParent.tournament));
        var i = scopeParent.tournament.rounds.indexOf(round);
        temp.rounds.splice(i, 1);
        for (var i=0; i<temp.rounds.length; i++){
            temp.rounds[i].number = i;
        }
        tournamentService.update(temp).then(function(){
            scopeParent.tournament = JSON.parse(JSON.stringify(temp));
            $rootScope.$emit("UpdateResult");
            scopeParent.tournament.rounds.forEach(function(rnd){
                tournamentService.verifyRound(scopeParent.tournament, rnd.number)
            });
            $uibModalInstance.close();
        }).catch(function(err){
            console.error(err);
            scope.error = true;
        });
    };
    this.cancel = function () {
        $uibModalInstance.dismiss('cancel');
    };
});
