new Vue({
	el: '#main',
	delimiters: ['%%', '%%'],
	data: {
		projectlist: []
	},
	created: function(){
		var self = this;
		self.getProjects();
	},
	methods: {
		getProjects: function() {
			var self = this;
			axios.get(
				"/projectlist",{
			}).then(function(res){
				if(res.data.code == 200){
					self.projectlist = res.data.result.projectlist;
				}
				console.log(res.data);
			}).catch(function(error){
				console.log(error);
			});
		},
		locationDB: function(id) {
			location.href = "/db/" + id;
		}
	}
});
