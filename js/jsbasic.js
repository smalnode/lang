var me = {
    firstName: "Wang",
    lastName: "Deqin",
    hi: function() {
        console.log("Hi, I'm", this.firstName, this.lastName);
    },
    verify: function() {
        console.log("this === me:", this === me);
    }
}

var refMe = me;
refMe.firstName = "Smal"
refMe.lastName = "Acnt"
refMe.hi();
refMe.verify();

function anymous() {
    var any = "a word";
    console.log("Am I called by global:", this == global);
    console.log("Called by:", this);
    return { any };
}

anymous();

var a = new anymous();
console.log(typeof a, a, a.any);

