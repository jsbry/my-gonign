new Vue({
	el: '#main',
	delimiters: ['%%', '%%'],
	data: {
		id: "",
		updated: "",
		db_name: "",
		db_engine: "",
		db_charset: "",
		viewable: {
			dbname: true
		}
	},
	created: function(){
		var self = this;
		params = location.href.split("/");
		self.id = params[params.length -1];
		self.getDbinfo();
	},
	methods: {
		Changed: function(t) {
			var self = this;
			self.viewable[t] = !self.viewable[t];
		},
		getDbinfo: function() {
			var self = this;
			axios.get(
				"/dbinfo/" + self.id,{
			}).then(function(res){
				console.log(res.data);
			}).catch(function(error){
				console.log(error);
			});
		},
	}
});
