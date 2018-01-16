new Vue({
	el: '#main',
	delimiters: ['%%', '%%'],
	data: {
		projectlist: []
	},
	created: function () {
		var self = this;
		self.getProjects();
	},
	methods: {
		newRegisterModal: function () {
			var self = this;
			self.$prompt('プロジェクト名を入力してください', '新規プロジェクト', {
				confirmButtonText: '登録',
				cancelButtonText: '閉じる',
			}).then(function (value) {
				self.$message({
					type: 'success',
					message: 'Your email is:' + value
				});
			}).catch(function () {
				self.$message({
					type: 'info',
					message: 'Input canceled'
				});
			});
		},
		getProjects: function () {
			var self = this;
			axios.get(
				"/projectlist", {
				}).then(function (res) {
					if (res.data.code == 200) {
						self.projectlist = res.data.result.projectlist;
					} else {
						self.networkConnectError();
					}
				}).catch(function (error) {
					self.networkConnectError();
				});
		},
		networkConnectError: function () {
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
		locationDB: function (id) {
			location.href = "/db/" + id;
		},
		locationAPI: function (id) {
			location.href = "/api/" + id;
		}
	}
});
