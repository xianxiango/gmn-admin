function compomentProductSeletor() {
    this.cataid = 0;

    this.productCompomentHtml = '<div class="product" data-name="{0}"> ' +
    '<div class="img"> ' +
        '<img src="{1}" title="{2}" data-url="{1}" data-type="image" data-name="{0}">     ' +                                   
        '<div id="product-select-{2}" data-id="{2}" data-name="{0}" data-url="{1}" class="select product-select">选择使用</div> ' +
        '<div class="action"> ' +
            '<a href="javascript:;" data-action="delete" title="编辑" data-path="{1}" data-name="{0}"> ' +
                '<i class="iconfont">&#58882</i> ' +
            '</a> ' +
        '</div> ' +
    '</div> ' +
    '<p class="name" title="{2}" data-name="{0}" data-path="{1}"> ' +
        '<span>{0}</span> ' +
    '</p> ' +
    '</div>';

    this.catagoryCompomentHtml = '<li class="{1}" data-id="{2}" data-type="">{0}</li>';


    this.bindCatagoryClick = function() {
        var _this = this;
        $(".product-panel-tab ul li").on('click',function(e){
            var btnTarget = $(this);
            console.log(e.target.dataset.id);
            var cataid = e.target.dataset.id;
            _this.getProductListByCataId(cataid,function(ret){
                if(ret.ret_code == 0) {
                    $(".product-list").html('');
                    ret.data.products.forEach(function(elment){
                        var html = _this.productCompomentHtml.formatUnicorn(elment.name,elment.images,elment.id)
                        $(".product-list").append(html);
                    });
                }//end if
    
                btnTarget.addClass("current");
                btnTarget.siblings().removeClass("current");    
                _this.bindProductSelect();        
            });
        });
    }


    this.bindProductSelect = function () {
        var _this = this;
        $(".product-select").on('click',function(e){
            console.log("select ",e);
            _this.afterSelectedEvent(e);
        });
    }

    this.getProductListByCataId = function(cataid,callback) {
        var _this = this;
        $.get( "/compoment/product_list_bycata",{
            cataid:cataid
        }, function( ret ) {
            callback(ret);
        });
    };

    this.reloadProductList = function() {
        var _this = this;
        $.get( "/compoment/product_list", function( ret ) {
            if(ret.ret_code == 0) {
                $(".product-list").html('');
                $(".product-panel-tab ul").html('');
                ret.data.products.forEach(function(elment){
                    var html = _this.productCompomentHtml.formatUnicorn(elment.name,elment.images,elment.id)
                    $(".product-list").append(html);
                });

                ret.data.catagorys.forEach(function(elment){
                    var addClass = "";
                    if(elment.id == ret.data.selected) 
                        addClass = "current";
                    var html = _this.catagoryCompomentHtml.formatUnicorn(elment.name,addClass,elment.id);
                    $(".product-panel-tab ul").append(html);
                });

                _this.bindCatagoryClick();
                _this.bindProductSelect();
    
            }//end if
        });
    };

    this.loadProductSeletor = function() {
        var _this = this;
        
        $.get( "/compoment/product_selector", function( data ) {
            $("#product-selector").append(data);
            _this.reloadProductList();
    
            $('#product-selector #icon-refresh').on('click',function(){
                console.log("click selector refresh button");
                _this.reloadProductList();
            }); 
                
        });
    };

    this.afterSelectedEvent = function(e) {

    };

    this.setSeleftEvent = function(fn) {
        this.afterSelectedEvent = fn;
    }

}



var productSeletor = new compomentProductSeletor();
productSeletor.loadProductSeletor();