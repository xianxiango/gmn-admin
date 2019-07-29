//产品缩略图代码

function fixEditor() {
   $("#btn-submit").on('click',function(){
        console.log('click');
        var editor_data = CKEDITOR.instances['TextArea1'].getData();
        console.log(editor_data);
        $('#TextArea1').val(editor_data);
        
    });
}

function showViewImage(url) {
    $('#product-preview > div.thumb.hasImg > img').attr("src",url);
    $('#product-preview > div.thumb.hasImg').show();
}

function hideViewImage() {
    $('#product-preview > div.thumb.hasImg > img').attr("src","");
    $('#product-preview > div.thumb.hasImg').hide();
}

function bindViewImageEvent() {
   $(".ui-sortable").on('click','.view',function(){
        console.log($(".ui-sortable-handle"));
        var $this = $(this).parent('li');
        $this.siblings().removeClass("active");
        $this.addClass("active");

        var url = $this.find('img').attr("src");        
        showViewImage(url);
    });
}

function bindControlIconEvent() {
   $(".ui-sortable").on('click',".iconfont",function(){
        var $this = $(this).parent('div').parent('li');
        var childCount = $('#product-preview > div.list > ul > li').length;
        if (childCount == 1) {
            hideViewImage();
        }
        $this.remove();
    });
}

$(function () {
	fixEditor();

    var imageListHtml = 
    '<li class="ui-sortable-handle">' +
    '    <div class="view"> ' +
    '      <img src="{0}" alt="产品图片"> ' +
    '    </div> ' +
    '    <input type="hidden" name="img[]" value="{0}"> ' +
    '    <div class="action"> ' +
    '      <i class="iconfont" data-action="delete" title="删除"></i> ' +
    '    </div> ' +
    '</li>';
    
    $("#btn-chose-file").on('click',function(){
        callFileExplorerSrcFunc = function(url) {
            var html = imageListHtml.formatUnicorn(url)
            $('.ui-sortable').append(html);
            $('.list').show();
            $('#file-explorer-modal-dialog').modal('toggle');

            showViewImage(url);
        }
        $('#file-explorer-modal-dialog').modal('toggle');
    });


    bindViewImageEvent();
    bindControlIconEvent();
});



///