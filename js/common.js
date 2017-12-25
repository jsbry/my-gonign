ELEMENT.locale(ELEMENT.lang.ja);

new Vue({
	el: '#navbar',
	delimiters: ['%%', '%%'],
	data: {
		isCollapse: true,
	},
	methods: {
		headerCollapse: function(){
			var self = this;
			self.isCollapse = !self.isCollapse;
			if(self.isCollapse == true){
				$("#main").css("width","calc(100% - 65px)");
				$(".el-aside").css("width", "65px");
			}else{
				$("#main").css("width","calc(100% - 200px)");
				$(".el-aside").css("width", "200px");
			}
		}
	}
})
