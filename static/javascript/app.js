ko.bindingHandlers.draggable = {
  init: function (element, valueAccessor, allBindingsAccessor, viewModel) {
   $(element).draggable();
  },
  update: function (element, valueAccessor, allBindingsAccessor, viewModel) {
   if (valueAccessor()) {
    $(element).draggable('enable');
   } else {
    $(element).draggable('disable');
   }
  }
};

ko.bindingHandlers.resizable = {
  init: function (element, valueAccessor, allBindingsAccessor, viewModel) {
   $(element).resizable();
  },
  update: function (element, valueAccessor, allBindingsAccessor, viewModel) {
   if (valueAccessor()) {
    $(element).resizable('enable');
   } else {
    $(element).resizable('disable');
   }
  }
};

var ViewModel = function() {
	var self = this;
	self.collection = ko.observableArray();
	$.getJSON("/all", function(data) {
		if ( data == null || data.Collection == null ) {
			return;
		}
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

	self.newOne = function(event) {
		console.log("newOne");
		console.log(event);
		var added = {
			top: ko.observable(event.clientY),
			left: ko.observable(event.clientX),
			width: ko.observable(200),
			height: ko.observable(300),
			body: ko.observable(''),
			color: ko.observable("#FFFFFF")
		};
		var index = self.collection.push(added);
		console.log(index);
		console.log(self.collection());
		self.edit(index);
	};

	self.edit = function(index) {
		self.active = self.collection[index];
	}

	self.flush = function(index) {
		console.log(index);
	}

	self.remove = function(data) {
		if ( event.target == "#document" ) {
			return;
		}
		console.log("remove start.");
		var removed = self.collection.remove(data);
		console.log(removed);
		console.log("remove end.");
	}

};
$(function() {
	var vm = new ViewModel();
	ko.applyBindings(vm);
});
