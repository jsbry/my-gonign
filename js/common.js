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
				$("section").css("width","calc(100% - 85px)");
			}else{
				$("section").css("width","calc(100% - 221px)");
			}
		}
	}
})
