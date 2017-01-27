var api = new Vue({
	el: '#app',
	data: function() {
		return {
			messages: [],
			error: false
		}
	},
	methods: {
		refreshMessages: function() {
			var self = this;
			jQuery.get("/api/twitter/list", function(data) {
				self.messages = data;
			})
			.fail(function() {
				self.error = "API request failed: /api/twitter/list";
			});
		},
		postMessage: function() {
			var self = this;
			var message = $(".message-text-input");
			jQuery.post("/api/twitter/add", {
				message: message.val()
			})
			.fail(function() {
				self.error = "API request failed: /api/twitter/add";
			})
			.done(function() {
				self.refreshMessages();
				message.val('');
			});
		}
	},
	mounted: function () {
		this.refreshMessages();
	}
});
