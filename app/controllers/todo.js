import Ember from 'ember';

export default Ember.ArrayController.extend({
ktodo:"",
 actions:{
  newTodo:function(kkk){

    var k = {title: kkk, isCompleted: false};
  }
  }

});
