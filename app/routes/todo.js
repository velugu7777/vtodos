import Ember from 'ember';

export default Ember.Route.extend({
  model:function(){
    return [
      {title: "hi", isCompleted: true},
      {title: "kkkk", isCompleted: true}
    ]//this.store.find("todo");
  }
});
