var alphabet = "68ABDEFGHJLMNQRTYabdefghijmnqrty";


var alphabetMap = [];
var amap = [];


function init() {
    alphabetMap[0] =  "68ABCDEFGHJKLMnpQrsTUWXYZi234579";
    alphabetMap[1] =  "68aBcdEfghjkLmNpqRstUwxYzi234579";
    alphabetMap[2] =  "68ABCdefGhjkLMnPqRStuWxyZi234579";
    alphabetMap[3] =  "68aBcDeFgHJKLmNPQrSTuwXyZi234579";
    
    // 重新排序，生成新的序列
    var list = [7, 5, 2, 18, 8, 11, 24, 25, 9, 12, 10, 13, 29, 31, 21, 1, 15, 20, 30, 28, 27, 17, 22, 26, 19, 4, 14, 3, 0, 23, 6, 16];
    
    for(var i=0;i<alphabetMap.length;i++) {
        var n = [];
        for(var j=0;j<list.length;j++) {
            n[j] = alphabetMap[i][list[j]];
        }
    
        alphabetMap[i] = n.toString().replace(/,/g,"");
        console.log("i : ",i," -- ", alphabetMap[i]);
    }
    
    
    var alphabet = alphabetMap[0].toUpperCase()
    
    for(var i=0;i<alphabetMap.length;i++) {
        if(alphabet != alphabetMap[i].toUpperCase()) {
            console.log("alphabet bug ",alphabet," with ",alphabetMap[i].toUpperCase())
            return 
        }
    }
    
    
    for(var i=0;i<alphabet.length;i++) {
        //console.log(alphabet[i]);
        amap[ alphabet[i] ] = i;
    }
    //console.log(alphabet);
    //console.log(amap);
    //console.log(amap.length);
    
    /*
    if(amap.length != 32) {
        console.log("len(amap) != 32");
        return 
    }
    */
    
    
}


function encode(uid) {
    var alphabet = alphabetMap[uid%4];
    var index = 5;
    var buf = [];
    buf[0] = 0;
    
    for(;uid > 0 && index >= 0; ) {
        var b1 = uid & 0x1f;
        uid = uid >> 5;
        buf[index] = b1;
        console.log(index,"--",b1 ," uid ",uid);
        index-- ;
    }

    console.log(alphabet);
    console.log(buf[0]);
    for(var i=0;i<buf.length;i++) {
        buf[i] = alphabet[buf[i]];
    }

    console.log(buf);
}

 
init();
encode(10531401);