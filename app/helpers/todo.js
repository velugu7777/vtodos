import Ember from 'ember';

export function todo(params/*, hash*/) {
  return params;
}

export default Ember.HTMLBars.makeBoundHelper(todo);
