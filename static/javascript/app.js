var wallApp = angular.module('wallApp', []);

wallApp.controller('WallController', function($scope, $http) {
	$http.get("/all").success(function(data) {
		$scope.collection = data.Collection;
		console.log($scope.collection);
	});

	$scope.newOne = function() {
		var added = {
			top: 10,
			left: 10,
			width: 100,
			height: 100,
			body: 'input hear.'
		};
		var index = $scope.collection.push($scope.collection, [added]);
		$scope.edit(index);
	};

	$scope.edit = function(index) {
		console.log(index);
		$scope.active = $scope.collection[index];
	}

	$scope.flush = function(index) {
		console.log(index);
	}

});
