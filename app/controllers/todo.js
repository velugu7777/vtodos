import Ember from 'ember';

export default Ember.ArrayController.extend({
ktodo:"",
 actions:{
  newTodo:function(post){
  //  this.sendAction("internalAction");
  //  var k = {title: kkk, isCompleted: false};
    //console.log(k);
  //  newTodo.push("Todos");
  //  console.log(newTodo);
  }
}
});
App.PostController = Ember.objectController.extend({
  actions:{
    newTodo:function(post){
      console.log(post.get("title"));
    }

  }
});
