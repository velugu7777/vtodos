import DS from 'ember-data';

var Todo= DS.Model.extend({
  title:DS.atrr('string'),
  isCompleted:DS.atrr('boolean'),
});
export default Todo;
