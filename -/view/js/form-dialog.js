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


//     Edit User
    var action_edit_user = document.querySelector('.action_edit_user');
    var action_edit_user_dialog = document.querySelector('#action_edit_user_dialog');
    if (! action_edit_user_dialog.showModal) {
      dialogPolyfill.registerDialog(action_edit_user_dialog);
    }
    action_edit_user.addEventListener('click', function() {
       action_edit_user_dialog.showModal();
    });
    action_edit_user_dialog.querySelector('button:not([disabled])')
    .addEventListener('click', function() {
      action_edit_user_dialog.close();
    });


  }());