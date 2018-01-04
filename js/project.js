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
				}else{
					networkConnectError();
				}
			}).catch(function(error){
				networkConnectError();
			});
		},
		networkConnectError: function(){
			var self = this;
			self.$alert('通信エラーが発生しました', 'エラー', {
				confirmButtonText: 'OK',
				callback: action => {
					self.$message({
						type: 'warning',
						message: `しばらく時間をおいてから再度お試しください`
					});
				}
			});
		},
		locationDB: function(id) {
			location.href = "/db/" + id;
		},
		locationAPI: function(id) {
			location.href = "/api/" + id;
		}
	}
});
