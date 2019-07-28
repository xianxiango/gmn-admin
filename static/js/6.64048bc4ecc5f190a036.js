webpackJsonp([6],{"6RD7":function(e,t){},ZY7b:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var n=i("lbHh"),a=i.n(n),s=i("gyMJ"),r={name:"admin",data:function(){return{loading:!1,adminList:[],permitList:[],deleteUserName:"",deleteUserIndex:0,centerDialogVisible:!1,createdUsername:"",createdPassword:""}},methods:{fetchPermitGroups:function(){var e=this;Object(s.a)("fetchPermitGroup").then(function(t){t&&(e.permitList=t)})},fetcgmnList:function(){var e=this;this.loading=!0,Object(s.a)("fetcgmnList").then(function(t){e.loading=!1,t&&t.length&&(t.map(function(e){return e.disabled=!0}),e.adminList=t)})},resetPermit:function(e){this.adminList[e].disabled&&(this.adminList[e].disabled=!1)},submitResetPermit:function(e){var t=this;Object(s.a)("setAdminPermit",{uid:this.adminList[e].uid,groupid:this.adminList[e].permit_groupid}).then(function(i){i&&(ELEMENT.Message.success("操作成功"),t.adminList[e].disabled=!0)})},handleDelete:function(e){this.deleteUserName=this.adminList[e].username,this.deleteUserIndex=e,this.centerDialogVisible=!0},confirmDelete:function(){var e=this;if(0===this.deleteUserIndex)return ELEMENT.Message.error("无法删除"),this.centerDialogVisible=!1,!1;Object(s.a)("removeAdmin",{id:this.adminList[this.deleteUserIndex].uid}).then(function(t){t&&(e.centerDialogVisible=!1,e.adminList.splice(e.deleteUserIndex,1),ELEMENT.Message.success("操作成功"))})},resetPassword:function(e){var t=this;ELEMENT.MessageBox.prompt("请输入新密码(6~12位数字或字母组合)","修改密码",{confirmButtonText:"确定",cancelButtonText:"取消"}).then(function(i){var n=i.value;n&&n.length>5?Object(s.a)("resetPassword",{id:t.adminList[e].uid,old_psw:t.adminList[e].password,new_psw:n}).then(function(i){i&&(a.a.get("username")===t.adminList[e].username?(a.a.remove("auth"),window.location.href=window.location.origin+"/login"):ELEMENT.Message.success("操作成功"))}):ELEMENT.Message.error("密码不符合要求")}).catch(function(e){return console.log(e)})},createdAdmin:function(){var e=this;ELEMENT.MessageBox.prompt("请输入用户名","创建账号",{confirmButtonText:"确定",cancelButtonText:"取消"}).then(function(t){var i=t.value;i?Object(s.a)("createdAdmin",{username:i}).then(function(t){t&&(e.createdUsername=t.username,e.createdPassword=t.password)}):ELEMENT.Message.error("账号不符合要求")}).catch(function(e){return console.log(e)})}},created:function(){this.fetcgmnList(),this.fetchPermitGroups()}},o={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",{staticClass:"app-container public-table"},[i("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{data:e.adminList,border:"",height:"calc(100% - 120px)"}},[i("el-table-column",{attrs:{label:"创建时间",align:"center",width:"180"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("i",{staticClass:"el-icon-time"}),e._v(" "),i("span",{staticStyle:{"margin-left":"10px"}},[e._v(e._s(t.row.crtime))])]}}])}),e._v(" "),i("el-table-column",{attrs:{label:"登录IP",align:"center",width:"160"},scopedSlots:e._u([{key:"default",fn:function(t){return t.row.last_login_ip.length>4?[i("span",[e._v(e._s(t.row.last_login_ip))])]:void 0}}])}),e._v(" "),i("el-table-column",{attrs:{label:"用户名",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("span",[e._v(e._s(t.row.username))])]}}])}),e._v(" "),i("el-table-column",{attrs:{label:"所属权限组",align:"center",width:"360"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("el-select",{attrs:{placeholder:"未分配",disabled:t.row.disabled},on:{change:function(i){e.submitResetPermit(t.$index)}},model:{value:t.row.permit_groupid,callback:function(i){e.$set(t.row,"permit_groupid",i)},expression:"scope.row.permit_groupid"}},e._l(e.permitList,function(e){return i("el-option",{key:e.groupid,attrs:{label:e.name,value:e.groupid}})}))]}}])}),e._v(" "),i("el-table-column",{attrs:{label:"操作",align:"center",width:"380"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("el-button",{attrs:{size:"mini"},on:{click:function(i){e.resetPermit(t.$index)}}},[e._v("修改权限组")]),e._v(" "),i("el-button",{attrs:{size:"mini"},on:{click:function(i){e.resetPassword(t.$index)}}},[e._v("修改密码")]),e._v(" "),i("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(i){e.handleDelete(t.$index)}}},[e._v("删除账号")])]}}])})],1),e._v(" "),i("div",{staticClass:"created-admin"},[e.createdUsername&&e.createdPassword?i("b",[i("span",{staticStyle:{color:"#409eff","margin-right":"20px"}},[e._v("账号："+e._s(e.createdUsername))]),i("span",{staticStyle:{color:"#ED1C24","margin-right":"20px"}},[e._v("密码："+e._s(e.createdPassword))])]):e._e(),e._v(" "),e.createdUsername&&e.createdPassword?i("el-tooltip",{attrs:{effect:"dark",content:"复制账号信息",placement:"top"}},[i("el-button",{directives:[{name:"copy",rawName:"v-copy",value:"账号："+e.createdUsername+"、密码："+e.createdPassword,expression:"`账号：${createdUsername}、密码：${createdPassword}`"}],staticStyle:{"margin-right":"30px"},attrs:{type:"info",icon:"el-icon-share",circle:""}})],1):e._e(),e._v(" "),i("el-button",{staticStyle:{width:"300px"},attrs:{type:"primary"},on:{click:e.createdAdmin}},[e._v("生成一个管理员账号")])],1),e._v(" "),i("el-dialog",{attrs:{title:"提示",visible:e.centerDialogVisible,width:"30%",center:"","append-to-body":!0},on:{"update:visible":function(t){e.centerDialogVisible=t}}},[i("span",[e._v("是否删除管理员\n      "),i("a",{staticStyle:{color:"red"},attrs:{href:"javascript:;"}},[e._v(" "+e._s(e.deleteUserName)+" ")]),e._v("账号？")]),e._v(" "),i("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[i("el-button",{on:{click:function(t){e.centerDialogVisible=!1}}},[e._v("取 消")]),e._v(" "),i("el-button",{attrs:{type:"primary"},on:{click:e.confirmDelete}},[e._v("确 定")])],1)])],1)},staticRenderFns:[]};var c=i("VU/8")(r,o,!1,function(e){i("6RD7")},"data-v-4b155a12",null);t.default=c.exports}});
//# sourceMappingURL=6.64048bc4ecc5f190a036.js.map