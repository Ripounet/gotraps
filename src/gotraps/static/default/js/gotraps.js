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
		});
	});
	
})(jQuery);