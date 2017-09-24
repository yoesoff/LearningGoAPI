var User = Backbone.Model.extend({

});

var Employee = User.extend({

});

var user = new User();
var employe = new Employee();

// Destroy object
user.destroy({
    success: function () {
        console.log("The model has been destroyed successfully");
    }
});

