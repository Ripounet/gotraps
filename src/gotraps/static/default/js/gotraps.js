(function($) {
	
	function show(selector){
		var elt = $(selector);
		if( ! elt.is(":visible") )
			elt.collapse("show");
	}
	
	function hide(selector){
		var elt = $(selector);
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
	
	$(".overview-link").click(function() {
		switchToZone("#overview");
		hide(".navbar-collapse.trap-list");
	});
	
	$("a.trap-link").click(function() {
		resetTrapView();
		var item = $(this).attr("href");
		item = item.substring(1);  // Remove the #
		
		hide(".navbar-collapse.trap-list");
		loadTrap(item);
	});
	
	function loadTrap(item){
		//console.log("loadTrap("+item+")");
		// Get the snippet
		var codepath = "content/" + item + "/code.go";
		$.get(codepath, function(data) {
			var code= data;
			$("#gotrap-code").html(code);
		});
		
		// Get the trap metadata
		var metapath = "content/" + item + "/trap.json";
		$.getJSON(metapath, function(data, status, xhr) {
			var trap= data.trap;
			$("#trap-title").html(trap.title);
			$("#trap-intro").html(trap.intro);
		});

		// Get the discussion text
		var discussionpath = "content/" + item + "/discussion.html";
		$.get(discussionpath, function(data) {
			var discussion= data;
			$("#trap-discussion").html(discussion);
		});
	}
	
	$("#btn-run").click(function() {
		var code=$("#gotrap-code").text();
		var trapname=$("#trap-title").html();
		//var compileUrl = "http://play.golang.org/compile"; 
		// ^^ would fail due to same-origin policy!
		var compileUrl = "/compile";  // <- this is a custom proxy
		show("#collapse-output");
		show("#stdout");
		// Todo spinner?
		$("#stdout").html("<i>compiling...</i>");
		$.post( compileUrl, 
				{ version: "2", body: code, trapname: trapname },
				function(data) {
					//alert(data);
					var events = data.Events;
					var errors = data.Errors;
					if(events){
						var output = "";
						events.forEach(function(event){
							output += event.Message;
						});
						$("#stdout").text(output);
					}else{
						hide("#stdout");
					}
					if(errors){
				        $("#compile-errors").text(errors);
				        show("#compile-errors");
					}else{
						hide("#compile-errors");
					}
				},
				"json")
			    .fail( function(xhr, textStatus, errorThrown) {
			        var response = xhr.responseJSON;
			        $("#compile-errors").html(response.Errors);
					show("#compile-errors");
					hide("#stdout");
			    });
	});

	$(".see-also-link").click(function() {
		switchToZone("#see-also");
		hide(".navbar-collapse.trap-list");
	});
	
	
	function highlightTrapInMenus(){
		$(".sidebar li").removeClass("active");
		$(".trap-list li").removeClass("active");
		$(this).parent().addClass("active");
	}
	
	$(".trap-list li a").click(highlightTrapInMenus);
	$(".sidebar li a").click(highlightTrapInMenus);
	
	$(".type-green").attr("title","Beginner");
	$(".type-orange").attr("title","Tricky");
	$(".type-yellow").attr("title","FYI");
	
	// Bookmarkable anchors
	if( window.location.hash.indexOf("#") != -1 ){
		$(".sidebar li").removeClass("active");
		var hash = window.location.hash;
		hash = hash.substr(1);  // Remove #
		if( hash == "overview" ){
			switchToZone("#overview");
			$(".overview-link").parent().addClass("active");
		}else if( hash == "see-also" ){
			switchToZone("#see-also");
			$(".see-also-link").parent().addClass("active");
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

	var lefteye = $("img.left-eye");
	var righteye = $("img.right-eye");
	$("body").mousemove(function(e) {
		// Homepage illustration animation
		lefteye.css({ transform: "rotate( " + (180 - e.pageY/3) + "deg)" });
		righteye.css({ transform: "rotate( " + (e.pageX/3) + "deg)" });
	})

})(jQuery);