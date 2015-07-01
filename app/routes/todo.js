import Ember from 'ember';

export default Ember.Route.extend({

model:function(){
  return[
    {title:'anji',age:25},
    {title:'velugu',age:24},
  ];
}
});
