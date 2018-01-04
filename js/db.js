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
		},
		editableTabsValue: '2',
		editableTabs: [{
			title: 'Tab 1',
			name: '1',
			content: 'Tab 1 content'
		}, {
			title: 'Tab 2',
			name: '2',
			content: 'Tab 2 content'
		}],
		tabIndex: 2
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
				if(res.data.code == 200){
					self.updated = res.data.result.updated;
					self.db_name = res.data.result.db_name;
					self.db_engine = res.data.result.db_engine;
					self.db_charset = res.data.result.db_charset;
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
		handleTabsEdit(targetName, action) {
			if (action === 'add') {
				let newTabName = ++this.tabIndex + '';
				this.editableTabs.push({
					title: 'New Tab',
					name: newTabName,
					content: 'New Tab content'
				});
				this.editableTabsValue = newTabName;
			}
			if (action === 'remove') {
				let tabs = this.editableTabs;
				let activeName = this.editableTabsValue;
				if (activeName === targetName) {
					tabs.forEach((tab, index) => {
						if (tab.name === targetName) {
							let nextTab = tabs[index + 1] || tabs[index - 1];
							if (nextTab) {
								activeName = nextTab.name;
							}
						}
					});
				}

				this.editableTabsValue = activeName;
				this.editableTabs = tabs.filter(tab => tab.name !== targetName);
			}
		}
	}
});
