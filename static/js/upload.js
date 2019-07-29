//工具定义
String.prototype.formatUnicorn = String.prototype.formatUnicorn ||
function () {
    "use strict";
    var str = this.toString();
    if (arguments.length) {
        var t = typeof arguments[0];
        var key;
        var args = ("string" === t || "number" === t) ?
            Array.prototype.slice.call(arguments)
            : arguments[0];

        for (key in args) {
            str = str.replace(new RegExp("\\{" + key + "\\}", "gi"), args[key]);
        }
    }

    return str;
};

//定义调用文件浏览器的来源
var callFileExplorerSrcFunc = function(url) {
    console.log("no src call file explorer ! ",url);
}

function initWebUpload() {
    var thumbnailWidth = 100; 
    var thumbnailHeight = 100;
    // 初始化Web Uploader

    var uploader = WebUploader.create({

        // 选完文件后，是否自动上传。
        auto: true,

        // swf文件路径
        swf: '/static/js/webupload/Uploader.swf',

        // 文件接收服务端。
        server: '/attachment/upload',

        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#uploadFileLabel',

        // 只允许选择图片文件。
        accept: {
            title: 'Images',
            extensions: 'gif,jpg,jpeg,bmp,png',
            mimeTypes: 'image/*'
        }
    });

    // 当有文件添加进来的时候
    uploader.on( 'fileQueued', function( file ) {
        console.log("file add in");
        console.log(file);
        var $li = $(
                '<div id="' + file.id + '" class="file-item thumbnail">' +
                    '<img>' +
                    '<div class="info">' + file.name + '</div>' +
                '</div>'
                ),
            $img = $li.find('img');


        // $list为容器jQuery实例
        //$list.append( $li );

        // 创建缩略图
        // 如果为非图片文件，可以不用调用此方法。
        // thumbnailWidth x thumbnailHeight 为 100 x 100
        
        uploader.makeThumb( file, function( error, src ) {
            if ( error ) {
                $img.replaceWith('<span>不能预览</span>');
                return;
            }

            $img.attr( 'src', src );
        }, thumbnailWidth, thumbnailHeight );
        
        //显示
        $("#fileList").append($li);
    });


    // 文件上传过程中创建进度条实时显示。
    uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
            $percent = $li.find('.progress span');

        // 避免重复创建
        if ( !$percent.length ) {
            $percent = $('<p class="progress"><span></span></p>')
                    .appendTo( $li )
                    .find('span');
        }

        $percent.css( 'width', percentage * 100 + '%' );
    });

    // 文件上传成功，给item添加成功class, 用样式标记上传成功。
    uploader.on( 'uploadSuccess', function( file ) {
        console.log("upload success");
        $( '#'+file.id ).addClass('upload-state-done');
        reloadFileList();
    });

    // 文件上传失败，显示上传出错。
    uploader.on( 'uploadError', function( file ) {
        var $li = $( '#'+file.id ),
            $error = $li.find('div.error');

        // 避免重复创建
        if ( !$error.length ) {
            $error = $('<div class="error"></div>').appendTo( $li );
        }

        $error.text('上传失败');
    });

    // 完成上传完了，成功或者失败，先删除进度条。
    uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').remove();
    });


 
    console.log("init uploader");
    $('#uploadFileLabel div:nth-child(1)').attr('style','position: absolute; top: 0px; left: 0px; width: 100px; height: 30px; overflow: hidden; bottom: auto; right: auto;');
    return uploader;
}



var templateHtml = '<div class="file" data-name="{0}"> ' +
    '<div class="img"> ' +
        '<img src="{1}" title="{2}" data-url="{1}" data-type="image" data-name="{0}">     ' +                                   
        '<div id="file-select-{2}" data-url="{1}" class="select">选择使用</div> ' +
        '<div class="action"> ' +
            '<a href="javascript:delFile({2});" data-action="delete" title="删除" data-path="{1}" data-name="{0}"> ' +
                '<i class="iconfont"></i> ' +
            '</a> ' +
        '</div> ' +
    '</div> ' +
    '<p class="name" title="{2}" data-name="{0}" data-path="{1}"> ' +
        '<span>{0}</span> ' +
    '</p> ' +
'</div>';


function bindFileSelectEvent(id) {
    $("#file-select-"+id).on('click',function(e){
        console.log("select ",e);
        console.log(e.currentTarget.dataset.url);
        callFileExplorerSrcFunc(e.currentTarget.dataset.url)
    });
}

function reloadFileList() {
    $.get( "/attachment/list", function( data ) {
        console.log(data); 
        //data.data[0]
        if(data.ret_code == 0) {
            $(".file-list").html('');
            data.data.forEach(function(elment){
                var html = templateHtml.formatUnicorn(elment.filename,elment.filepath,elment.id)
                $(".file-list").append(html);
                bindFileSelectEvent(elment.id);
            });
        }//end if
    });
};



 

function loadFileExplorer(fn) {
    $.get( "/attachment/explorer", function( data ) {
        $("#file-explorer-content").append(data);
        fn();

        $('#icon-refresh').on('click',function(){
            reloadFileList();
        }); 
            
    });
};


function delFile(id) {
    console.log("del file ",id);
}


$(function() {
    console.log( "ready!" );

    //jquery
    loadFileExplorer(function(){
        var uploader = initWebUpload();
        reloadFileList() ;
    });
    

});









