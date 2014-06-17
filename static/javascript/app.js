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

ko.bindingHandlers.colorpicker = {
  init: function (element, valueAccessor, allBindingsAccessor, viewModel) {
   $(element).tinycolorpicker();
//  },
//  update: function (element, valueAccessor, allBindingsAccessor, viewModel) {
//   if (valueAccessor()) {
//    $(element).resizable('enable');
//   } else {
//    $(element).resizable('disable');
//   }
  }
};

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
		var index = self.collection.push(added);
		self.edit(index);
	};

	self.edit = function(index) {
		self.active = self.collection[index];
	}

	self.flush = function(index) {
		console.log(index);
	}

};
$(function() {
	ko.applyBindings(new ViewModel());
});
