(function($) {
	
	function show(id){
		var elt = $(id);
		if( ! elt.is(":visible") )
			elt.collapse("show");
	}
	
	function hide(id){
		var elt = $(id);
		if( elt.is(":visible") )
			elt.collapse("hide");
	}
	
	function switchToZone(zone){
		["#overview", "#trap-zone", "#see-also"].forEach(function(z){
			if( z == zone )
				show(z);
			else
				hide(z);
		});
	}
	
	function resetTrapView(){
		switchToZone("#trap-zone");
		
		$("#trap-title").html("...");
		$("#trap-intro").html("<i>loading...</i>");
		$("#gotrap-code").html("<i>loading...</i>");
		$("#stdout").html("(No output before you click the [Run] button)");
		$("#compile-errors").html("(No errors before you click the [Run] button)");
		hide("#collapse-output");
		$("#compile-errors").html("");
		hide("#compile-errors");
		$("#trap-discussion").html("");
		hide("#collapse-discussion");
	}
	
	$("#overview-link").click(function() {
		switchToZone("#overview");
	});
	
	$("a.trap-link").click(function() {
		resetTrapView();
		var item = $(this).attr("href");
		item = item.substring(1);  // Remove the #
		
		loadTrap(item);
	});
	
	function loadTrap(item){
		//console.log("loadTrap("+item+")");
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
	}
	
	$("#btn-run").click(function() {
		var code=$("#gotrap-code-hidden").html();
		//var compileUrl = "http://play.golang.org/compile"; 
		// ^^ would fail due to same-origin policy!
		var compileUrl = "/compile";  // <- this is a custom proxy
		show("#collapse-output");
		show("#stdout");
		// Todo spinner?
		$("#stdout").html("<i>compiling...</i>");
		$.post( compileUrl, 
				{ version: "2", body: code },
				function(data) {
					//alert(data);
					var events = data.Events;
					var output = "";
					events.forEach(function(event){
						output += event.Message;
					});
					$("#stdout").html(output);
					hide("#compile-errors");
				},
				"json")
			    .fail( function(xhr, textStatus, errorThrown) {
			        var response = xhr.responseJSON;
			        $("#compile-errors").html(response.Errors);
					show("#compile-errors");
					hide("#stdout");
			    });
	});

	$("#see-also-link").click(function() {
		switchToZone("#see-also");
	});
	
	$(".sidebar li a").click(function() {
		$(".sidebar li").removeClass("active");
		$(this).parent().addClass("active");
	});
	
	// Bookmarkable anchors
	if( window.location.hash.indexOf("#") != -1 ){
		$(".sidebar li").removeClass("active");
		var hash = window.location.hash;
		hash = hash.substr(1);  // Remove #
		if( hash == "overview" ){
			switchToZone("#overview");
			$("#overview-link").parent().addClass("active");
		}else if( hash == "see-also" ){
			switchToZone("#see-also");
			$("#see-also-link").parent().addClass("active");
		}else{
			resetTrapView();
			loadTrap(hash);
			$(".trap-link").each(function(){
				if( $(this).attr("href") == "#"+hash )
					$(this).parent().addClass("active");
			});
		}
	}

	function log(msg){
		console.log(msg);
	}
})(jQuery);