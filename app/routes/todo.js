import Ember from 'ember';

export default Ember.Route.extend({
  model:function(){
    return [
      {title: "", isCompleted: false},
      {title: "", isCompleted: false}
   ];
  }
});
