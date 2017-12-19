new Vue({
	el: '#container',
	delimiters: ['%%', '%%'],
	data: {
		db_name: "",
		db_engine: "",
		db_charset: "",
		viewable: {
			dbname: true
		}
	},
	created: function(){
		var self = this;
		self.db_name = $("#db_name").val();
	},
	methods: {
		Changed: function(t) {
			var self = this;
			self.viewable[t] = !self.viewable[t];
		},
	}
});
