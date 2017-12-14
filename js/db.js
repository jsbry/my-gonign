new Vue({
	el: '#container',
	data: {
		viewable: {
			dbname: true
		}
	},
	created: function(){
		var self = this;
	},
	methods: {
		Editable: function(t) {
			var self = this;
			self.viewable[t] = false;
		}
	}
});
