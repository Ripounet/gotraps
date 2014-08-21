(function($) {
	
	function switchToZone(zone){
		["#overview", "#trap-zone", "#see-also"].forEach(function(z){
			if( z == zone ){
				if( ! $(z).is(":visible") )
					$(z).collapse("show");
			}else{
				if( $(z).is(":visible") )
					$(z).collapse("hide");
			}
		});
	}
	
	function resetTrapView(){
		switchToZone("#trap-zone");
		
		$("#gotrap-code").html("");
		$("#stdout").html("");
		$("#stdout").collapse("hide");
		$("#compile-errors").html("");
		$("#compile-errors").collapse("hide");
		$("#trap-discussion").html("");
		$("#trap-discussion").collapse("hide");
	}
	
	$("#overview-link").click(function() {
		switchToZone("#overview");
	});
	
	$("a.trap-link").click(function() {
		resetTrapView();
		var item = $(this).attr("href");
		item = item.substring(1);  // Remove the #
		
		// Get the snippet
		var codepath = "content/" + item + ".code";
		$.get(codepath, function(data) {
			var code= data;
			$("#gotrap-code").html(code);
			$("#gotrap-code-hidden").html(code);
		});
		
		// Get the trap metadata
		var metapath = "content/" + item + ".json";
		$.getJSON(metapath, function(data, status, xhr) {
			var trap= data.trap;
			$("#trap-title").html(trap.title);
			$("#trap-intro").html(trap.intro);
			$("#trap-discussion").html(trap.discussion);
		});
	});
	
	$("#btn-run").click(function() {
		var code=$("#gotrap-code-hidden").html();
		//var compileUrl = "http://play.golang.org/compile"; 
		// ^^ would fails due to same-origin policy!
		var compileUrl = "/compile";  // <- this is a custom proxy
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
	});

	$("#title-discussion").click(function() {
		$("#trap-discussion").collapse("toggle");
	});
	
	$("#see-also-link").click(function() {
		switchToZone("#see-also");
	});
	
	$(".sidebar li a").click(function() {
		$(".sidebar li").removeClass("active");
		$(this).parent().addClass("active");
	});

	function log(msg){
		console.log(msg);
	}
})(jQuery);