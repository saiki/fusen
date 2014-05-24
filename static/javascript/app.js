var wallController = angular.module('wallController', []);

wallController.controller('WallController', function($scope, $http) {
	$http.get("/all").success(function(data) {
		$scope.collection = data;
	});

	$scope.add = function() {
		$scope.collection.push({});
		alert($scope.collection);
		console.log($scope.collection);
	};
});
