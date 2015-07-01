import DS from 'ember-data';

export default DS.Model.extend({
  title:DS.atrr('string'),
  age:DS.atrr('number'),
});
