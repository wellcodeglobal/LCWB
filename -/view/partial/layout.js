  (function() {
    'use strict';
//     Person
    var action_person = document.querySelector('.action_person');
    var action_person_dialog = document.querySelector('#action_person_dialog');
    if (! action_person_dialog.showModal) {
      dialogPolyfill.registerDialog(action_person_dialog);
    }
    action_person.addEventListener('click', function() {
       action_person_dialog.showModal();
    });
    action_person_dialog.querySelector('button:not([disabled])')
    .addEventListener('click', function() {
      action_person_dialog.close();
    });
    
//     Group
    var action_group = document.querySelector('.action_group');
    var action_group_dialog = document.querySelector('#action_group_dialog');
    if (! action_group_dialog.showModal) {
      dialogPolyfill.registerDialog(action_group_dialog);
    }
    action_group.addEventListener('click', function() {
       action_group_dialog.showModal();
    });
    action_group_dialog.querySelector('button:not([disabled])')
    .addEventListener('click', function() {
      action_group_dialog.close();
    });
  }());