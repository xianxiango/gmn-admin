
 

function changeAdtype(adtype) {
	$("#ad-choose").siblings().hide();
	if (adtype == "ads") { $("#ad-ads").show();}
	if (adtype == "banner") { $("#ad-banner").show();}
	if (adtype == "product") { $("#ad-product").show();}
	if (adtype == "goods") { $("#ad-goods").show();}
}
 
function bindOperation() {
	$(".goods-edit").on("click",function(e){
		console.log("edit",e);
	})

	$(".goods-del").on("click",function(e){
		var advertId = e.currentTarget.dataset.id;
		$.post( "/admin/advert/deladvert",{
			id:advertId,
		},function( data ) {
			$( ".result" ).html( data );
		});

		//移除元素
		$(this).parent().parent("li").remove();
	})
}

$(function () {
	changeAdtype("product");

    //添加页面
    $("#btn-chose-file").on('click',function(){
        callFileExplorerSrcFunc = function(url) {
            $('#article-img').attr("value",url);
            $('#img-display').attr("src",url);
            $('#img-display').show()
            $('#file-explorer-modal-dialog').modal('toggle');
        }
        $('#file-explorer-modal-dialog').modal('toggle');
    });


   $(".btn-chose-file-ex").on('click',function(){
	   var _this = $(this);
	   var parent = _this.parent();
        callFileExplorerSrcFunc = function(url) {
			var img = parent.find("img");
			var input = parent.find(".input_thumb");
			img.attr("src",url);
			input.attr("value",url);
			img.show();
            $('#file-explorer-modal-dialog').modal('toggle');
        }
        $('#file-explorer-modal-dialog').modal('toggle');
	});
	
	$(".btn-choose-product").on('click',function(e){
		var adtype = e.currentTarget.dataset.type;
		var dlg = $('#product-selector-modal-dialog');
		productSeletor.setSeleftEvent(function(e){
			var html = "" ;
			if(adtype == "ad-goods"){
				html = '<li>'+
					'<input type="text"   style="display:none;"  name="goods-goods-id[]" value="{0}"  />'+
					'<img  src="{2}" alt="横幅缩略图" style="display:block;" width="110" height="110">'+
					'<p>{1}</p>'+
					'<p><a class="goods-edit" href="javascript:;">修改</a> <a class="goods-del" href="javascript:;">删除</a></p>'+
					'</li>';
				var dataset = e.currentTarget.dataset;
				var liHtml = html.formatUnicorn(dataset.id,dataset.name,dataset.url);
				$(".goods-goods ul").append(liHtml);
			}else {
				html = '<li>'+
					'<input type="text"   style="display:none;"  name="product-goods-id[]" value="{0}"  />'+
					'<img  src="{2}" alt="横幅缩略图" style="display:block;" width="110" height="110">'+
					'<p>{1}</p>'+
					'<p><a class="goods-edit" href="javascript:;">修改</a> <a class="goods-del" href="javascript:;">删除</a></p>'+
					'</li>';
				var dataset = e.currentTarget.dataset;
				var liHtml = html.formatUnicorn(dataset.id,dataset.name,dataset.url);
				$(".product-goods ul").append(liHtml);
			}
			bindOperation();
			dlg.modal('toggle');
		})

		dlg.modal('toggle');
        
	});
	
	
	

	$("#combo-adtype").change(function (e) {
        var adtype = $(this).val();
		changeAdtype(adtype);
    });
});
