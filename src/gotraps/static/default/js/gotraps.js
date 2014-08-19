(function($) {
	
	$("#title-stdout").click(function() {
		$("#stdout").collapse('toggle');
	});

	$("#title-discussion").click(function() {
		$("#discussion").collapse('toggle');
	});
	
	$("a.trap-link").click(function() {
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
		alert(code);
		$.post("http://play.golang.org/compile", 
				{ version: "2", body: code },
				function(data) {
					alert(data);
				},
				"json");
		// "json" fails due to same-origin policy
		// "jsonp" fails because can't POST, and service does not accept GET
		// todo: a small handler proxyCompile
	});
	
})(jQuery);