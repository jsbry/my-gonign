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
		Changed: function(t) {
			var self = this;
			self.viewable[t] = !self.viewable[t];
		},
	}
});
