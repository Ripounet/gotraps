(function($) {
	
	$("#title-stdout").click(function() {
		$("#stdout").collapse('toggle');
	});

	$("#title-discussion").click(function() {
		$("#discussion").collapse('toggle');
	});
	
	function resetView(){
		$("#gotrap-code").html("");
		$("#stdout").html("");
		$("#stdout").collapse("hide");
		$("#compile-errors").html("");
		$("#compile-errors").collapse("hide");
	}
	
	$("a.trap-link").click(function() {
		resetView();
		var item = $(this).attr("href");
		item = item.substring(1);  // Remove the #
		var codepath = "content/" + item + ".code";
		$.get(codepath, function(data) {
			var code= data;
			$("#gotrap-code").html(code);
			$("#gotrap-code-hidden").html(code);
		});
	});
	
	$("#btn-run").click(function() {
		var code=$("#gotrap-code-hidden").html();
		//var compileUrl = "http://play.golang.org/compile"; 
		var compileUrl = "/compile";  // <- this is a custom proxy
		//alert(code);
		$.post( compileUrl, 
				{ version: "2", body: code },
				function(data) {
					//alert(data);
					var events = data.Events;
					var event = events[0];
					var output = event.Message;
					$("#stdout").html(output);
					$("#stdout").collapse("show");
					$("#compile-errors").collapse("show");
				},
				"json");
		// "json" fails due to same-origin policy
		// "jsonp" fails because can't POST, and service does not accept GET
		// todo: a small handler proxyCompile
	});
	
})(jQuery);