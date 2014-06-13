
var ViewModel = function() {
	var self = this;
	self.collection = ko.observableArray();
	$.getJSON("/all", function(data) {
		for ( var i = 0; i < data.Collection.length; i++ ) {
			var fusen = {
				top: ko.observable(data.Collection[i].top),
				left: ko.observable(data.Collection[i].left),
				width: ko.observable(data.Collection[i].width),
				height: ko.observable(data.Collection[i].height),
				body: ko.observable(data.Collection[i].body),
				color: ko.observable(data.Collection[i].color)
			}
			self.collection.push(fusen);
		}
	});

	self.newOne = function() {
		var added = {
			top: ko.observable(10),
			left: ko.observable(10),
			width: ko.observable(100),
			height: ko.observable(100),
			body: ko.observable(''),
			color: ko.observable("#FFFFFF")
		};
		var index = self.collection.push(self.collection, added);
		console.log(self.collection);
		self.edit(index);
	};

	self.edit = function(index) {
		self.active = self.collection[index];
	}

	self.flush = function(index) {
		console.log(index);
	}

};
var vm = new ViewModel();

ko.applyBindings(vm, document.body);
