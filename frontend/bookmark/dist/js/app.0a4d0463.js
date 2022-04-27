(function(e){function t(t){for(var r,o,c=t[0],i=t[1],l=t[2],d=0,m=[];d<c.length;d++)o=c[d],Object.prototype.hasOwnProperty.call(s,o)&&s[o]&&m.push(s[o][0]),s[o]=0;for(r in i)Object.prototype.hasOwnProperty.call(i,r)&&(e[r]=i[r]);u&&u(t);while(m.length)m.shift()();return n.push.apply(n,l||[]),a()}function a(){for(var e,t=0;t<n.length;t++){for(var a=n[t],r=!0,c=1;c<a.length;c++){var i=a[c];0!==s[i]&&(r=!1)}r&&(n.splice(t--,1),e=o(o.s=a[0]))}return e}var r={},s={app:0},n=[];function o(t){if(r[t])return r[t].exports;var a=r[t]={i:t,l:!1,exports:{}};return e[t].call(a.exports,a,a.exports,o),a.l=!0,a.exports}o.m=e,o.c=r,o.d=function(e,t,a){o.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:a})},o.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},o.t=function(e,t){if(1&t&&(e=o(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var a=Object.create(null);if(o.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)o.d(a,r,function(t){return e[t]}.bind(null,r));return a},o.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return o.d(t,"a",t),t},o.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},o.p="";var c=window["webpackJsonp"]=window["webpackJsonp"]||[],i=c.push.bind(c);c.push=t,c=c.slice();for(var l=0;l<c.length;l++)t(c[l]);var u=i;n.push([0,"chunk-vendors"]),a()})({0:function(e,t,a){e.exports=a("56d7")},"00a9":function(e,t,a){"use strict";a("9456")},"1a52":function(e,t,a){"use strict";a("3aff")},"3aff":function(e,t,a){},"56d7":function(e,t,a){"use strict";a.r(t);a("e260"),a("e6cf"),a("cca6"),a("a79d");var r=a("2b0e"),s=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{attrs:{id:"app"}},[a("router-view")],1)},n=[],o={name:"App",data:function(){return{}}},c=o,i=a("2877"),l=Object(i["a"])(c,s,n,!1,null,null,null),u=l.exports,d=a("f309");r["a"].use(d["a"]);var m=new d["a"]({}),p=a("bc3a"),v=a.n(p),f=a("8c4f"),g=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-app",{attrs:{app:""}},[a("v-main",{staticClass:"grey lighten-3"},[a("v-alert",{staticClass:"mx-auto",style:e.alertMessage.alertStyle,attrs:{type:e.alertMessage.alertType}},[e._v(e._s(e.alertMessage.text))]),a("v-card",{staticClass:"mx-auto my-6 pa-5",attrs:{elevation:"2",width:"330",height:"450"}},[a("v-form",{ref:"formRef",model:{value:e.valid,callback:function(t){e.valid=t},expression:"valid"}},[e.status?a("div",{staticClass:"grey--text"},[e._v("注册")]):a("div",{staticClass:"grey--text"},[e._v("登录")]),a("v-text-field",{attrs:{label:"username",required:"",rules:e.nameRules,messages:e.messages.nameErrMessage},model:{value:e.formData.username,callback:function(t){e.$set(e.formData,"username",t)},expression:"formData.username"}}),e.status?a("v-text-field",{attrs:{label:"E-mail",required:"",rules:e.emailRules},model:{value:e.formData.email,callback:function(t){e.$set(e.formData,"email",t)},expression:"formData.email"}}):e._e(),e.status?a("v-row",{staticClass:"py-0"},[a("v-col",{staticClass:"py-0"},[a("v-text-field",{attrs:{label:"验证码",required:""},model:{value:e.formData.veriCode,callback:function(t){e.$set(e.formData,"veriCode",t)},expression:"formData.veriCode"}})],1),a("v-col",{staticClass:"py-0"},[a("v-btn",{staticClass:"ml-2 mt-5",attrs:{color:"primary",small:""},on:{click:e.getVerificationCode}},[e._v("发送验证码")]),a("p",{staticClass:"mt-1 ml-5 blue--text small"},[e._v(e._s(e.messages.codeSendMessage))])],1)],1):e._e(),a("v-text-field",{attrs:{label:"password",type:"password",required:"",rules:e.pwdRules},model:{value:e.formData.password,callback:function(t){e.$set(e.formData,"password",t)},expression:"formData.password"}}),e.status?a("v-text-field",{attrs:{label:"confirm password",type:"password",required:"",rules:e.pwdRules,messages:e.messages.pwdErrMessage},on:{blur:e.checkSame},model:{value:e.formData.confirmPassword,callback:function(t){e.$set(e.formData,"confirmPassword",t)},expression:"formData.confirmPassword"}}):e._e(),a("v-row",{staticStyle:{position:"absolute",left:"30px",top:"360px"}},[a("v-col",{attrs:{cols:"5"}},[e.status?a("v-btn",{staticClass:"mr-4 mt-5",attrs:{color:"primary"},on:{click:e.register}},[e._v("注册")]):a("v-btn",{staticClass:"mr-4 mt-5",attrs:{color:"primary"},on:{click:e.login}},[e._v("登录")])],1),a("v-col",{staticClass:"pr-0",attrs:{cols:"7"}},[e.status?a("div",{staticClass:"mt-6 ml-4 grey--text",staticStyle:{cursor:"pointer"},on:{click:function(t){return e.setStatus(0)}}},[e._v("已有账号？登录")]):a("div",{staticClass:"mt-6 ml-4 grey--text",staticStyle:{cursor:"pointer"},on:{click:function(t){return e.setStatus(1)}}},[e._v("没有账号？注册")])])],1)],1)],1)],1)],1)},h=[],b=a("1da1"),k=(a("96cf"),a("ac1f"),a("00b4"),{data:function(){return{alertMessage:{alertStyle:{width:"300px",height:"50px",visibility:"hidden"},alertType:"success",text:""},messages:{pwdErrMessage:"",nameErrMessage:"",codeSendMessage:""},valid:!0,status:1,formData:{username:"",email:"",password:"",confirmPassword:"",veriCode:""},nameRules:[function(e){return!!e||"username is required"}],emailRules:[function(e){return!!e||"E-mail is required"},function(e){return/.+@.+\..+/.test(e)||"E-mail must be valid"}],pwdRules:[function(e){return!!e||"password is required"},function(e){return e.length>5||"password is too short"}]}},methods:{checkSame:function(){return this.formData.password!=this.formData.confirmPassword?(this.messages.pwdErrMessage="两次输入密码不一致",!1):(this.messages.pwdErrMessage="",!0)},getVerificationCode:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,e.$http.post("verificationCode",{username:e.formData.username,email:e.formData.email});case 2:a=t.sent,r=a.data,200!=r.code?e.messages.nameErrMessage=r.msg:(e.messages.nameErrMessage="",e.messages.codeSendMessage="验证码已发送");case 5:case"end":return t.stop()}}),t)})))()},register:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(!e.$refs.formRef.validate()||!e.checkSame()){t.next=6;break}return t.next=3,e.$http.post("register/".concat(e.formData.veriCode),{username:e.formData.username,password:e.formData.password,email:e.formData.email});case 3:a=t.sent,r=a.data,200==r.code?(e.alertMessage.alertStyle.visibility="visible",e.alertMessage.alertType="success",e.alertMessage.text="注册成功"):(e.alertMessage.alertStyle.visibility="visible",e.alertMessage.alertType="error",e.alertMessage.text=r.msg);case 6:case"end":return t.stop()}}),t)})))()},login:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,e.$http.post("login",{username:e.formData.username,password:e.formData.password});case 2:a=t.sent,r=a.data,200==r.code?(e.alertMessage.alertStyle.visibility="visible",e.alertMessage.alertType="success",e.alertMessage.text="登录成功",window.localStorage.setItem("token",r.token),e.$router.push("/space")):(e.alertMessage.alertStyle.visibility="visible",e.alertMessage.alertType="error",e.alertMessage.text=r.msg);case 5:case"end":return t.stop()}}),t)})))()},setStatus:function(e){this.status=e}}}),w=k,y=(a("1a52"),a("6544")),C=a.n(y),x=a("0798"),_=a("7496"),S=a("8336"),D=a("b0af"),B=a("62ad"),M=a("4bd4"),R=a("f6c4"),V=a("0fd9"),O=a("8654"),j=Object(i["a"])(w,g,h,!1,null,"640b1748",null),$=j.exports;C()(j,{VAlert:x["a"],VApp:_["a"],VBtn:S["a"],VCard:D["a"],VCol:B["a"],VForm:M["a"],VMain:R["a"],VRow:V["a"],VTextField:O["a"]});var P=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-app",[a("v-main",{staticClass:"brown lighten-5"},[a("div",{staticClass:"grey--text ml-12 mt-5 mb-0",staticStyle:{"font-size":"20px",display:"inline-block"}},[e._v("分类")]),a("v-dialog",{attrs:{persistent:"","max-width":"600px"},scopedSlots:e._u([{key:"activator",fn:function(t){var r=t.on,s=t.attrs;return[a("v-btn",e._g(e._b({staticClass:"pa-0 ml-6",attrs:{color:"primary",fab:"",dark:""}},"v-btn",s,!1),r),[a("v-icon",{staticClass:"ma-0",staticStyle:{cursor:"pointer"},attrs:{size:"40"}},[e._v(e._s("mdi-plus"))])],1)]}}]),model:{value:e.categoryDialog,callback:function(t){e.categoryDialog=t},expression:"categoryDialog"}},[a("v-card",{staticClass:"py-2 px-5"},[a("v-card-title",[a("span",{staticClass:"text-h5"},[e._v("新建标签")])]),a("v-text-field",{attrs:{label:"标签名称",required:""},model:{value:e.newCategory,callback:function(t){e.newCategory=t},expression:"newCategory"}}),a("v-card-actions",[a("v-spacer"),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(t){e.categoryDialog=!1}}},[e._v("关闭")]),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(t){return e.addCategory()}}},[e._v("提交")])],1)],1)],1),0!=e.currentCategory?a("div",{staticClass:"mr-12 mt-5 grey--text",staticStyle:{float:"right","font-size":"20px",cursor:"pointer"},on:{click:function(t){return e.backToParent()}}},[e._v("回到上级")]):e._e(),a("v-sheet",{staticClass:"sheet brown lighten-5"},e._l(e.categories,(function(t){return a("v-card",{key:t.ID,staticClass:"category",attrs:{width:"200",height:"240"},on:{click:function(a){return e.changeCategory(t.ID)}}},[a("v-img",{attrs:{width:"200",height:"130",src:t.pic}}),a("v-card-text",{staticClass:"pa-0"},[a("v-row",{staticClass:"ma-0 pa-0"},[a("v-col",{attrs:{cols:"8"}},[a("div",[e._v(e._s(t.name))])]),a("v-col",{staticClass:"px-0",attrs:{cols:"2"}},[a("v-dialog",{attrs:{persistent:"","max-width":"600px"},scopedSlots:e._u([{key:"activator",fn:function(r){var s=r.on,n=r.attrs;return[a("v-icon",e._g(e._b({staticClass:"ma-0",on:{click:function(a){return e.selectCid(t.ID)}}},"v-icon",n,!1),s),[e._v(e._s("mdi-plus"))])]}}],null,!0),model:{value:e.bookmarkDialog2,callback:function(t){e.bookmarkDialog2=t},expression:"bookmarkDialog2"}},[a("v-card",[a("v-card-title",[a("span",{staticClass:"text-h5"},[e._v("新建书签")])]),a("v-text-field",{attrs:{label:"书签URL",required:""},model:{value:e.newBookmarkUrl,callback:function(t){e.newBookmarkUrl=t},expression:"newBookmarkUrl"}}),a("v-card-actions",[a("v-spacer"),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(t){e.bookmarkDialog2=!1}}},[e._v("关闭")]),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(t){return e.addBookmark()}}},[e._v("提交")])],1)],1)],1)],1),a("v-col",{staticClass:"pa-0",attrs:{cols:"2"}},[a("v-menu",{attrs:{"offset-y":""},scopedSlots:e._u([{key:"activator",fn:function(t){var r=t.on,s=t.attrs;return[a("v-icon",e._g(e._b({staticClass:"pa-0 mt-3"},"v-icon",s,!1),r),[e._v(e._s("mdi-trash-can-outline"))])]}}],null,!0)},[a("v-list",{staticClass:"py-0"},[a("v-list-item",{staticClass:"pa-1 green--text",staticStyle:{cursor:"pointer"},on:{click:function(a){return e.deleteCategory(t.ID,!0)}}},[e._v("仅删除标签")]),a("v-list-item",{staticClass:"pa-1 red--text",staticStyle:{cursor:"pointer"},on:{click:function(a){return e.deleteCategory(t.ID,!1)}}},[e._v("删除标签及内容")])],1)],1)],1)],1),a("p",{staticClass:"grey--text small mt-1 mb-0 ml-3"},[e._v("创建时间")]),a("p",{staticClass:"grey--text small mt-1 ml-3"},[e._v("书签总数")])],1)],1)})),1),a("v-divider",{staticClass:"mt-5 mb-0"}),a("div",{staticClass:"grey--text ml-12 mt-5 mb-0",staticStyle:{"font-size":"20px",display:"inline-block"}},[e._v("书签")]),a("v-dialog",{attrs:{persistent:"","max-width":"600px"},scopedSlots:e._u([{key:"activator",fn:function(t){var r=t.on,s=t.attrs;return[a("v-btn",e._g(e._b({staticClass:"pa-0 ml-6",attrs:{color:"primary",fab:"",dark:""}},"v-btn",s,!1),r),[a("v-icon",{staticClass:"ma-0",staticStyle:{cursor:"pointer"},attrs:{size:"40"}},[e._v(e._s("mdi-plus"))])],1)]}}]),model:{value:e.bookmarkDialog,callback:function(t){e.bookmarkDialog=t},expression:"bookmarkDialog"}},[a("v-card",[a("v-card-title",[a("span",{staticClass:"text-h5"},[e._v("新建书签")])]),a("v-text-field",{attrs:{label:"书签URL",required:""},model:{value:e.newBookmarkUrl,callback:function(t){e.newBookmarkUrl=t},expression:"newBookmarkUrl"}}),a("v-card-actions",[a("v-spacer"),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(t){e.bookmarkDialog=!1}}},[e._v("关闭")]),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(t){return e.addBookmark()}}},[e._v("提交")])],1)],1)],1),a("v-btn",{staticStyle:{"margin-left":"786px"},on:{click:function(t){return e.batchOperation()}}},[e._v("批量操作")]),a("v-btn",{staticClass:"ml-6",attrs:{disabled:!e.enableCut},on:{click:function(t){return e.cut()}}},[e._v("剪切")]),a("v-btn",{staticClass:"ml-6",attrs:{disabled:!e.enablePaste},on:{click:function(t){return e.paste()}}},[e._v("粘贴")]),a("v-sheet",{staticClass:"brown lighten-5 px-5"},e._l(e.bookmarks,(function(t,r){return a("v-card",{key:t.ID,staticClass:"mt-5 ml-7",staticStyle:{cursor:"pointer"},attrs:{width:"1000",height:"55"}},[a("v-row",{staticClass:"pa-0 ma-0 grey--text"},[a("v-col",{staticClass:"pa-0 ma-0",attrs:{cols:"1"},on:{click:function(a){return e.navigate(t.url)}}},[a("img",{attrs:{width:"55",height:"55",src:"data:image/png;base64,"+t.icon}})]),a("v-col",{staticClass:"pa-0 ma-0 hidden",attrs:{cols:"5","align-self":"center"},on:{click:function(a){return e.navigate(t.url)}}},[a("div",[e._v(e._s(t.title))])]),a("v-col",{staticClass:"pa-0 ma-0 hidden",attrs:{cols:"4","align-self":"center"},on:{click:function(a){return e.navigate(t.url)}}},[a("div",[e._v(e._s(t.url))])]),a("v-col",{staticClass:"pl-8 pt-4",attrs:{cols:"1"}},[a("v-icon",{on:{click:function(a){return e.deleteBookmark(t.ID)}}},[e._v(e._s("mdi-trash-can-outline"))])],1),a("v-col",{staticClass:"pa-0 ma-0",attrs:{cols:"1"}},[a("v-checkbox",{staticClass:"mt-4 ml-6 pa-0",attrs:{disabled:!e.enableBatchSelect},on:{click:function(a){return e.select(t.ID)}},model:{value:e.currentBatchSelect[r],callback:function(t){e.$set(e.currentBatchSelect,r,t)},expression:"currentBatchSelect[index]"}})],1)],1)],1)})),1)],1)],1)},E=[],T=(a("4ec9"),a("d3b7"),a("3ca3"),a("ddb0"),a("159b"),{data:function(){return{bookmarkDialog2:!1,bookmarkDialog:!1,categoryDialog:!1,parentCategory:[],currentCategory:0,selected:new Map,clipBoard:[],batchSelects:new Map,currentBatchSelect:[],categories:[],bookmarks:[],newCategory:"",newBookmarkUrl:"",selectedCid:0,enableBatchSelect:!1,enableCut:!1,enablePaste:!1,selectedNum:0}},created:function(){this.getCategories(),this.getBookmarks()},methods:{select:function(e){this.selected.has(e)?(this.selected.delete(e),this.selectedNum--):(this.selected.set(e,0),this.selectedNum++),this.selectedNum>0?this.enableCut=!0:this.enableCut=!1},cut:function(){var e=this;this.selected.forEach((function(t,a){e.clipBoard.push(a)})),this.enablePaste=!0,this.enableCut=!1},paste:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,e.$http.put("bookmark/",{ids:e.clipBoard,cid:e.currentCategory});case 2:a=t.sent,a.data,e.selected=new Map,e.batchSelects=new Map,e.currentBatchSelect=[],e.clipBoard=[],e.getBookmarks(),e.enableBatchSelect=!1,e.enablePaste=!1;case 11:case"end":return t.stop()}}),t)})))()},batchOperation:function(){this.enableBatchSelect=!this.enableBatchSelect,this.selected=new Map,this.batchSelects=new Map,this.currentBatchSelect=[],this.clipBoard=[],this.currentBatchSelect=[];for(var e=0;e<this.bookmarks.length;e++)this.currentBatchSelect.push(!1)},getCategories:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,e.$http.get("categories/".concat(e.currentCategory));case 2:a=t.sent,r=a.data,200==r.code&&(e.categories=r.data);case 5:case"end":return t.stop()}}),t)})))()},getBookmarks:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a,r,s;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return e.done=!1,t.next=3,e.$http.get("bookmark/".concat(e.currentCategory));case 3:if(a=t.sent,r=a.data,200==r.code)for(e.bookmarks=r.data,e.currentBatchSelect=[],s=0;s<e.bookmarks.length;s++)e.currentBatchSelect.push(!1);case 6:case"end":return t.stop()}}),t)})))()},navigate:function(e){window.open(e)},changeCategory:function(e){var t=this;return Object(b["a"])(regeneratorRuntime.mark((function a(){var r,s;return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:return t.parentCategory.push(t.currentCategory),r=t.currentCategory,s=t.currentBatchSelect,t.currentCategory=e,t.getCategories(),a.next=7,t.getBookmarks();case 7:t.batchSelects.set(r,s),void 0!=t.batchSelects.get(e)&&(t.currentBatchSelect=t.batchSelects.get(e));case 9:case"end":return a.stop()}}),a)})))()},backToParent:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return a=e.currentCategory,r=e.currentBatchSelect,e.currentCategory=e.parentCategory.pop(),e.getCategories(),t.next=6,e.getBookmarks();case 6:e.batchSelects.set(a,r),void 0!=e.batchSelects.get(e.currentCategory)&&(e.currentBatchSelect=e.batchSelects.get(e.currentCategory));case 8:case"end":return t.stop()}}),t)})))()},addCategory:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,e.$http.post("category",{pid:e.currentCategory,name:e.newCategory});case 2:e.getCategories(),e.categoryDialog=!1,e.newCategory="";case 5:case"end":return t.stop()}}),t)})))()},deleteCategory:function(e,t){var a=this;return Object(b["a"])(regeneratorRuntime.mark((function r(){return regeneratorRuntime.wrap((function(r){while(1)switch(r.prev=r.next){case 0:return r.next=2,a.$http.delete("category/".concat(e),{data:{ids:[],retain_content:t}});case 2:a.getCategories(),a.getBookmarks();case 4:case"end":return r.stop()}}),r)})))()},deleteBookmark:function(e){var t=this;return Object(b["a"])(regeneratorRuntime.mark((function a(){return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:return a.next=2,t.$http.delete("bookmark",{data:{ids:[e]}});case 2:t.getBookmarks();case 3:case"end":return a.stop()}}),a)})))()},addBookmark:function(){var e=this;return Object(b["a"])(regeneratorRuntime.mark((function t(){var a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return a=e.currentCategory,0!=e.selectedCid&&(a=e.selectedCid),t.next=4,e.$http.post("bookmark",{url:e.newBookmarkUrl,cid:a});case 4:e.newBookmarkUrl="",e.bookmarkDialog=!1,e.bookmarkDialog2=!1,e.selectedCid=0,e.getBookmarks();case 9:case"end":return t.stop()}}),t)})))()},selectCid:function(e){this.selectedCid=e}}}),I=T,q=(a("00a9"),a("99d9")),U=a("ac7c"),z=a("169a"),A=a("ce7e"),L=a("132d"),N=a("adda"),F=a("8860"),J=a("da13"),G=a("e449"),H=a("8dd9"),K=a("2fa4"),Q=Object(i["a"])(I,P,E,!1,null,"1b4ceb1a",null),W=Q.exports;C()(Q,{VApp:_["a"],VBtn:S["a"],VCard:D["a"],VCardActions:q["a"],VCardText:q["b"],VCardTitle:q["c"],VCheckbox:U["a"],VCol:B["a"],VDialog:z["a"],VDivider:A["a"],VIcon:L["a"],VImg:N["a"],VList:F["a"],VListItem:J["a"],VMain:R["a"],VMenu:G["a"],VRow:V["a"],VSheet:H["a"],VSpacer:K["a"],VTextField:O["a"]}),r["a"].use(f["a"]);var X=[{path:"/login",name:"login",component:$},{path:"/space",name:"space",component:W}],Y=new f["a"]({routes:X});Y.beforeEach((function(e,t,a){"/login"==e.path&&a();var r=window.localStorage.getItem("token");"/space"!=e.path||r?a():a("login")}));var Z=Y;r["a"].config.productionTip=!1,v.a.defaults.baseURL="http://asilentboy.cn:9999/bookmark/api/v1",v.a.interceptors.request.use((function(e){return e.headers.Authorization="Bearer ".concat(window.localStorage.getItem("token")),e})),r["a"].prototype.$http=v.a,new r["a"]({vuetify:m,router:Z,render:function(e){return e(u)}}).$mount("#app")},9456:function(e,t,a){}});
//# sourceMappingURL=app.0a4d0463.js.map