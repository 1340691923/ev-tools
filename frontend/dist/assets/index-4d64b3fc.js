import{j as De,C as ke,f as xe}from"./es-382a2972.js";import{G as Se,D as Ce}from"./es-index-3eb427b6.js";import{n as Ve,r as Me,i as Ye,c as je,d as Le,L as fe}from"./es-map-fa49449a.js";import{U as Te,I as Ne,D as Oe,K as Ue}from"./style-c1f85648.js";import{b as S,N as K,f as ye,r as c,o as i,k as y,R as he,c as h,j as u,a as r,w as d,F as Y,q as G,t as j,u as B,h as O,x as U,S as pe,y as me,p as Ie,E as x,D as Ae,G as Ee,B as ve,L as ie,g as ge,s as z,O as be,z as we,A as ne,e as le,T as ze}from"./index-15b0a2bf.js";import{J as se}from"./index-5fbead7a.js";import{_ as Q}from"./_plugin-vue_export-helper-c27b6911.js";import{c as Be}from"./index-0aae8c1b.js";const oe={shortcuts:[{text:"今天",value:()=>[dayjs().format("YYYY-MM-DD 00:00:00"),dayjs().format("YYYY-MM-DD 23:59:59")]},{text:"昨天",value:()=>[dayjs().subtract(-1,"day").format("YYYY-MM-DD 00:00:00"),dayjs().subtract(-1,"day").format("YYYY-MM-DD 23:59:59")]},{text:"一周",value:()=>[dayjs().subtract(-7,"day").format("YYYY-MM-DD 00:00:00"),dayjs().format("YYYY-MM-DD 23:59:59")]},{text:"两周",value:()=>[dayjs().subtract(-14,"day").format("YYYY-MM-DD 00:00:00"),dayjs().format("YYYY-MM-DD 23:59:59")]},{text:"一月",value:()=>[dayjs().subtract(-30,"day").format("YYYY-MM-DD 00:00:00"),dayjs().format("YYYY-MM-DD 23:59:59")]},{text:"三个月",value:()=>[dayjs().subtract(-90,"day").format("YYYY-MM-DD 00:00:00"),dayjs().format("YYYY-MM-DD 23:59:59")]},{text:"半年",value:()=>[dayjs().subtract(-180,"day").format("YYYY-MM-DD 00:00:00"),dayjs().format("YYYY-MM-DD 23:59:59")]},{text:"一年",value:()=>[dayjs().subtract(-360,"day").format("YYYY-MM-DD 00:00:00"),dayjs().format("YYYY-MM-DD 23:59:59")]}]},Ge={name:"Values",props:{data:{type:Array,default:()=>[]},value:{type:[String,Number,Array],default:""}},setup(e,{emit:l}){const a=S([]),s=S(e.value),t=()=>{s.value=[],l("update:value",s.value)},o=()=>{l("update:value",s.value)},w=f=>{a.value=[]};return K(()=>e.data,f=>{w()}),K(s,f=>{l("update:value",f)}),ye(()=>{w(e.data)}),{options:a,selectVal:s,cleanValues:t,changeValue:o}}};function Re(e,l,a,s,t,o){const w=c("el-select");return i(),y(w,{"collapse-tags-tooltip":"",modelValue:s.selectVal,"onUpdate:modelValue":l[0]||(l[0]=f=>s.selectVal=f),filterable:"",options:s.options,placeholder:"请输入筛选值",style:{width:"240px"},multiple:"","default-first-option":"","allow-create":"","collapse-tags":"","reserve-keyword":!1,clearable:""},null,8,["modelValue","options"])}const $e=Q(Ge,[["render",Re]]),Je={class:"ta-filter-condition"},Pe={style:{float:"left"}},Fe={style:{float:"right",color:"#8492a6","font-size":"13px"}},_e={__name:"ActionRow",props:{value:{type:Object,default:()=>({})},datas:{type:Object,default:()=>({})},options:{type:Array,default:()=>[]},dataTypeMap:{type:Array,default:()=>[]},tableTyp:{type:Number,default:0}},emits:["update:value"],setup(e,{emit:l}){const a=e,s=l,t=he(a.value),o=S(oe);K(()=>t.columnName,()=>{s("update:value",t)}),K(()=>t.comparator,()=>{s("update:value",t)}),K(()=>t.ftv,()=>{s("update:value",t)}),ye(()=>{o.value=oe});const w=C=>{let p=0;if(a.dataTypeMap.hasOwnProperty(a.tableTyp.toString()))for(const L of a.dataTypeMap[a.tableTyp.toString()])C===L.attribute_name&&(p=L.data_type);return Le[p]},f=()=>{const C=w(t.columnName);for(const p in C){t.comparator=p;break}s("update:value",t)};return(C,p)=>{const L=c("el-option"),V=c("el-option-group"),I=c("el-select"),N=c("el-input"),v=c("el-date-picker");return i(),h("div",null,[u("div",Je,[r(I,{modelValue:t.columnName,"onUpdate:modelValue":p[0]||(p[0]=m=>t.columnName=m),filterable:"",style:{width:"300px"},placeholder:"请选择",onChange:f},{default:d(()=>[(i(!0),h(Y,null,G(e.options,m=>(i(),y(V,{key:m.label,label:m.label},{default:d(()=>[(i(!0),h(Y,null,G(m.options,g=>(i(),y(L,{key:g.value,label:g.label,value:g.value},{default:d(()=>[u("span",Pe,j(g.label),1),u("span",Fe,j(g.value),1)]),_:2},1032,["label","value"]))),128))]),_:2},1032,["label"]))),128))]),_:1},8,["modelValue"]),r(I,{modelValue:t.comparator,"onUpdate:modelValue":p[1]||(p[1]=m=>t.comparator=m),style:{width:"80px"}},{default:d(()=>[(i(!0),h(Y,null,G(w(t.columnName),(m,g,A)=>(i(),y(L,{key:A,label:m,value:g},null,8,["label","value"]))),128))]),_:1},8,["modelValue"]),B(Ve).includes(t.comparator)?(i(),h(Y,{key:0},[],64)):B(Me).includes(t.comparator)?(i(),h(Y,{key:1},[r(N,{modelValue:t.ftv[0],"onUpdate:modelValue":p[2]||(p[2]=m=>t.ftv[0]=m),clearable:"",type:"number",style:{width:"150px"}},null,8,["modelValue"]),O(" ~ "),r(N,{modelValue:t.ftv[1],"onUpdate:modelValue":p[3]||(p[3]=m=>t.ftv[1]=m),clearable:"",type:"number",style:{width:"150px"}},null,8,["modelValue"])],64)):B(Ye).includes(t.comparator)?(i(),y(N,{key:2,modelValue:t.ftv,"onUpdate:modelValue":p[4]||(p[4]=m=>t.ftv=m),clearable:"",style:{width:"300px"}},null,8,["modelValue"])):B(je).includes(t.comparator)?(i(),y(v,{key:3,modelValue:t.ftv,"onUpdate:modelValue":p[5]||(p[5]=m=>t.ftv=m),type:"datetimerange",align:"right","default-time":["00:00:00","23:59:59"],"unlink-panels":"","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期",format:"yyyy-MM-dd HH:mm:ss","value-format":"yyyy-MM-dd HH:mm:ss","picker-options":B(oe)},null,8,["modelValue","picker-options"])):t.comparator=="match"?(i(),y(N,{key:4,modelValue:t.ftv,"onUpdate:modelValue":p[6]||(p[6]=m=>t.ftv=m),clearable:"",style:{width:"300px"}},null,8,["modelValue"])):(i(),y($e,{key:5,ref:"values",modelValue:t.ftv,"onUpdate:modelValue":p[7]||(p[7]=m=>t.ftv=m),"table-typ":e.tableTyp,style:{width:"300px"},data:t.columnName},null,8,["modelValue","table-typ","data"]))])])}}};const re=e=>(Ae("data-v-5e3c0ee4"),e=e(),Ee(),e),He={class:"xwl_main",style:{"margin-top":"20px"}},qe={class:"relation-editor globalFilters_xwl"},Ke={class:"relation-relation"},We=re(()=>u("em",{class:"relation-relation-line"},null,-1)),Qe={class:"relation-main"},Xe={class:"ta-multa-filter-condition"},Ze={key:0,class:"action-row row___xwl"},et={class:"action-right"},tt={key:1,class:"relation-editor"},at={class:"relation-relation"},lt=re(()=>u("div",{class:"relation-relation-sub"},"◆",-1)),ot=re(()=>u("em",{class:"relation-relation-line-sub"},null,-1)),nt=["onClick"],st={class:"relation-main"},it={class:"relation-row"},rt={class:"action-left"},dt={class:"action-right"},ut={style:{padding:"0 12px"}},ct={__name:"index",props:{value:{type:Object,default:()=>({relationLine:"且",relationArr:[],ftv:[]})},fontColor:{type:String,default:"#3d90ff"},tableTyp:{type:Number,default:0},dataTypeMap:{type:Array,default:()=>[]},options:{type:Array,default:()=>[]}},emits:["update:value"],setup(e,{emit:l}){const a=e,s=l,t=S(!0),o=he({...a.value}),w=()=>{t.value=!1,ve(()=>{t.value=!0})},f=v=>{const m=o.filts[v],g={filterType:"COMPOUND",filts:[],relation:"且"};g.filts.push(m),g.filts.push({columnName:a.options[0].options[0].value,comparator:"=",filterType:"SIMPLE",ftv:""}),o.filts[v]=g,s("update:value",o)},C=v=>{o.filts[v].filts.push({columnName:a.options[0].options[0].value,comparator:"=",filterType:"SIMPLE",ftv:""}),s("update:value",o)},p=(v,m)=>{if(o.filts[v].filts.length>2)o.filts[v].filts.splice(m,1);else{const g=[...o.filts[v].filts];g.splice(m,1);const A=g[0];o.filts[v]=A}s("update:value",o)},L=()=>{if(a.options.length==0){s("update:value",o),x({offset:60,message:"该索引没有字段，无法筛选",type:"error"});return}if(a.options[0].options.length===0){s("update:value",o),x({offset:60,message:"该索引没有字段，无法筛选",type:"error"});return}o.filts.push({columnName:a.options[0].options[0].value,comparator:"=",filterType:"SIMPLE",ftv:""}),s("update:value",o)},V=v=>{o.filts.splice(v,1),s("update:value",o),w()},I=()=>{o.relation=o.relation==="或"?"且":"或",s("update:value",o)},N=v=>{o.filts[v].relation=o.filts[v].relation==="或"?"且":"或",s("update:value",o)};return(v,m)=>{const g=c("el-button"),A=c("el-button-group");return i(),h("div",null,[u("div",He,[u("div",qe,[u("div",Ke,[We,o.filts.length>1?(i(),h("div",{key:0,class:"relation-relation-value",onClick:I},j(o.relation),1)):U("",!0)]),u("div",Qe,[(i(!0),h(Y,null,G(o.filts,(R,T)=>(i(),h("div",{key:T,class:"relation-row"},[u("div",Xe,[R.filterType==="SIMPLE"?(i(),h("div",Ze,[t.value?(i(),y(_e,{key:0,value:o.filts[T],"onUpdate:value":F=>o.filts[T]=F,options:a.options,"table-typ":a.tableTyp,"data-type-map":a.dataTypeMap,class:"action-left"},null,8,["value","onUpdate:value","options","table-typ","data-type-map"])):U("",!0),u("div",et,[r(A,null,{default:d(()=>[r(g,{type:"link",class:"actions_xwl_btn",onClick:()=>f(T),icon:B(pe)},null,8,["onClick","icon"]),r(g,{type:"link",class:"actions_xwl_btn",icon:B(me),onClick:()=>V(T)},null,8,["icon","onClick"])]),_:2},1024)])])):(i(),h("div",tt,[u("div",at,[lt,ot,u("div",{class:"relation-relation-value",onClick:()=>N(T)},j(R.relation),9,nt)]),u("div",st,[u("div",it,[(i(!0),h(Y,null,G(R.filts,(F,E)=>(i(),h("div",{key:E,class:"action-row row___xwl"},[u("div",rt,[(i(),y(_e,{key:(T+"_"+E).toString(),value:R.filts[E],"onUpdate:value":H=>R.filts[E]=H,options:a.options,"table-typ":a.tableTyp,"data-type-map":a.dataTypeMap,class:"action-left"},null,8,["value","onUpdate:value","options","table-typ","data-type-map"]))]),u("div",dt,[r(A,null,{default:d(()=>[r(g,{type:"link",class:"actions_xwl_btn",onClick:()=>C(T),icon:B(pe)},null,8,["onClick","icon"]),r(g,{type:"link",class:"actions_xwl_btn",icon:B(me),onClick:()=>p(T,E)},null,8,["icon","onClick"])]),_:2},1024)])]))),128))])])]))])]))),128))])])]),u("div",ut,[u("span",{style:Ie({color:a.fontColor}),class:"footadd___2D4YB",onClick:L},j(v.$t("增加条件")),5)])])}}},pt=Q(ct,[["__scopeId","data-v-5e3c0ee4"]]),de="/api/es_crud/";function mt(e){return ie({url:de+"GetList",method:"post",transformResponse:[l=>De.parse(l)],data:e})}function _t(e){return ie({url:de+"GetDSL",method:"post",data:e})}function ft(e){return ie({responseType:"arraybuffer",url:de+"Download",method:"post",data:e})}const yt={name:"Crud",setup(){return{ctx:ge().appContext.config.globalProperties}},components:{JsonEditor:se,FilterWhere:pt},computed:{jsonDataString(){return JSON.stringify(this.jsonData,null,"	")},queryDslString(){return JSON.stringify(this.queryDsl,null,"	")}},props:{indexName:{type:String,default:"test2"},attrMapProp:{type:Array,default:[]},eventAttrOptionsProp:{type:Array,default:[]}},data(){return{ctx:null,showJsonEditor1:!1,showJsonEditor2:!1,properties:{},openAddDialog:!1,tableLoading:!1,tableHeader:[],tableData:[],crudTab:"filter",eventAttrOptions:this.eventAttrOptionsProp,attrMap:this.attrMapProp,whereFilter:{filterType:"COMPOUND",filts:[],relation:"且"},sortList:[],input:{page:1,limit:20},count:0,typName:"",newDoc:{},jsonData:{},tableDataClone:[],drawerShow:!1,queryDslShow:!1,queryDsl:{},downloadLoading:!1}},mounted(){this.indexName!=""&&this.search(1)},methods:{download(){const e={index_name:this.indexName,relation:this.whereFilter,sort_list:this.sortList,es_connect:z.GetSelectEsConnID()},l=document.createElement("a");this.downloadLoading=!0,ft(e).then(a=>{if(a){const s=new Blob([a],{type:"application/vnd.ms-excel"}),t=URL.createObjectURL(s);l.href=t,l.download=this.indexName+".csv",l.click(),URL.revokeObjectURL(t)}this.downloadLoading=!1}).catch(a=>{this.downloadLoading=!1,console.error(a)})},drawerShowFn(){this.drawerShow=!1,this.$nextTick(()=>{this.drawerShow=!0,this.$nextTick(()=>{this.showJsonEditor1=!0})})},look(e){this.index=e,console.log("this.tableDataClone",this.tableDataClone),this.jsonData=this.tableDataClone[e],this.drawerShowFn()},async updateByID(){console.log("this.jsonData",this.jsonData);const e=this.jsonData,l=e._source,a={};a.es_connect=z.GetSelectEsConnID(),a.index=e._index,a.type_name=e._type,a.id=e._id,a.json=l;const s=await Te(a);s.code==0?(this.search(1),x({type:"success",message:s.msg})):x({type:"error",message:s.msg})},getEditDoc(e){try{const l=JSON.parse(e);this.jsonData=l}catch{}},getNewDoc(e){try{this.newDoc=JSON.parse(e)}catch{}},async add(){const e=this.newDoc,l={};l.es_connect=z.GetSelectEsConnID(),l.index=this.indexName,l.type_name=this.typName,l.json=e;const a=await Ne(l);a.code==0?x({type:"success",message:a.msg+"(_id为:"+a.data._id+")",duration:2e4}):x({type:"error",message:a.msg})},drawerHandleClose(e){e()},async deleteByID(e,l){be.confirm("确定删除ID为【"+e._id+"】的这条文档吗?","警告",{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(async()=>{const a={};a.es_connect=z.GetSelectEsConnID(),a.index_name=this.indexName,a.type=this.typName,a.id=e._id;const s=await Oe(a);s.code==0?setTimeout(async()=>{await this.search(1),x({type:"success",message:s.msg})},1e3):x({type:"error",message:s.msg})}).catch(a=>{console.error(a)})},async initTableData(e){if(e.hasOwnProperty("hits")&&e.hits.hits.length>0){this.typName=e.hits.hits[0]._type;const l=[];this.tableDataClone=Be(e.hits.hits);for(const a in e.hits.hits){const s=e.hits.hits[a]._source;s._id=e.hits.hits[a]._id,s._score=e.hits.hits[a]._score,s.xwl_index=a,l.push(s)}this.tableData=l;for(const a in this.tableData)for(const s in this.tableData[a])(typeof this.tableData[a][s]=="object"||typeof this.tableData[a][s]=="array")&&(this.tableData[a][s]=JSON.stringify(this.tableData[a][s]))}this.$nextTick(()=>{this.$refs.multipleTable.doLayout()})},async getMapping(){const e={};e.es_connect=z.GetSelectEsConnID(),e.index_name=this.indexName;const{data:l,code:a,msg:s}=await fe(e);if(a==0){if(l.ver==6){const t=Object.keys(l.list[e.index_name].mappings);for(const f in t)if(t[f]=="_default_"){t.splice(f,1);break}const o=Object.keys(l.list[e.index_name].mappings[t[0]].properties);this.properties=l.list[e.index_name].mappings[t[0]].properties;const w=[];w.push("_id");for(const f in o)o[f]!="_id"&&w.push(o[f]);this.tableHeader=w}else if(l.ver==7||l.ver==8){if(JSON.stringify(l.list[e.index_name].mappings)=="{}"){x({type:"error",message:"该索引没有设置映射结构"});return}const t=Object.keys(l.list[e.index_name].mappings.properties);this.properties=l.list[e.index_name].mappings.properties;const o=[];o.push("_id");for(const w in t)t[w]!="_id"&&o.push(t[w]);this.tableHeader=o}}else x({type:"error",message:s})},pushSort(){this.sortList.push({col:"",sortRule:"desc"})},deleteSort(e){this.sortList.splice(e,1)},handleSizeChange(e){this.input.limit=e,this.search(1)},async getdsl(){const e={index_name:this.indexName,relation:this.whereFilter,sort_list:this.sortList,es_connect:z.GetSelectEsConnID(),page:this.input.page,limit:this.input.limit},l=await _t(e);if(l.code!=0){x({type:"error",message:l.msg});return}this.queryDsl=l.data.list,this.queryDslShow=!0,this.$nextTick(()=>{this.showJsonEditor2=!0})},async search(e){this.tableLoading=!0,e?this.input.page=e:this.input.page=1;const l={index_name:this.indexName,relation:this.whereFilter,sort_list:this.sortList,es_connect:z.GetSelectEsConnID(),page:this.input.page,limit:this.input.limit},a=await mt(l);if(a.code!=0){x({type:"error",message:a.msg}),this.tableData=[],this.tableLoading=!1;return}this.count=a.data.count,await this.getMapping(),this.count>0?await this.initTableData(a.data.list):this.tableData=[],this.tableLoading=!1}}},ht={class:"app-container",style:{height:"95%"}},vt={class:"search-container"},gt={class:"search-container"},bt={class:"relation-row"},wt={class:"action-left"},Dt={style:{float:"left"}},kt={style:{float:"right",color:"#8492a6","font-size":"13px"}},xt={style:{padding:"0 12px"}},St={class:"search-container"},Ct={key:1,style:{padding:"40px",width:"300px",height:"800px","text-align":"center",margin:"0px auto",position:"relative",top:"30%"}};function Vt(e,l,a,s,t,o){const w=c("filter-where"),f=c("el-tab-pane"),C=c("el-option"),p=c("el-option-group"),L=c("el-select"),V=c("el-button"),I=c("el-tabs"),N=c("el-form-item"),v=c("el-form"),m=c("el-table-column"),g=c("el-button-group"),A=c("el-table"),R=c("el-pagination"),T=c("el-card"),F=c("el-tag"),E=c("json-editor"),H=c("el-drawer"),X=c("el-empty"),D=we("loading");return i(),h("div",null,[u("div",ht,[a.indexName!=""?(i(),h(Y,{key:0},[u("div",vt,[r(I,{modelValue:t.crudTab,"onUpdate:modelValue":l[2]||(l[2]=n=>t.crudTab=n)},{default:d(()=>[r(f,{lazy:!0,label:e.$t("筛选"),name:"filter"},{default:d(()=>[r(w,{value:t.whereFilter,"onUpdate:value":l[0]||(l[0]=n=>t.whereFilter=n),"table-typ":"2","data-type-map":t.attrMap,options:t.eventAttrOptions},null,8,["value","data-type-map","options"])]),_:1},8,["label"]),r(f,{lazy:"true",label:e.$t("排序"),name:"sort"},{default:d(()=>[u("div",gt,[u("div",bt,[(i(!0),h(Y,null,G(t.sortList,(n,_)=>(i(),h("div",{key:_,class:"action-row row___xwl"},[u("div",wt,[r(L,{modelValue:t.sortList[_].col,"onUpdate:modelValue":b=>t.sortList[_].col=b,filterable:"",style:{width:"300px","margin-left":"30px"},placeholder:e.$t("请选择")},{default:d(()=>[(i(!0),h(Y,null,G(a.eventAttrOptionsProp,b=>(i(),y(p,{key:b.label,label:b.label},{default:d(()=>[(i(!0),h(Y,null,G(b.options,M=>(i(),y(C,{key:M.value,label:M.label,value:M.value},{default:d(()=>[u("span",Dt,j(M.label),1),u("span",kt,j(M.value),1)]),_:2},1032,["label","value"]))),128))]),_:2},1032,["label"]))),128))]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder"]),r(L,{modelValue:t.sortList[_].sortRule,"onUpdate:modelValue":b=>t.sortList[_].sortRule=b,style:{width:"100px","margin-left":"30px"},placeholder:e.$t("请选择排序规则"),filterable:""},{default:d(()=>[r(C,{label:e.$t("正序排"),value:"asc"},null,8,["label"]),r(C,{label:e.$t("倒序排"),value:"desc"},null,8,["label"])]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder"]),r(V,{style:{"margin-left":"30px"},class:"actions_xwl_btn",onClick:b=>o.deleteSort(_)},{default:d(()=>[O("删除")]),_:2},1032,["onClick"])])]))),128))])]),u("div",xt,[u("span",{style:{color:"#3d90ff"},class:"footadd___2D4YB",onClick:l[1]||(l[1]=(...n)=>o.pushSort&&o.pushSort(...n))},j(e.$t("增加排序")),1)])]),_:1},8,["label"])]),_:1},8,["modelValue"])]),r(T,{class:"box-card"},{default:d(()=>[u("div",St,[r(v,null,{default:d(()=>[r(N,null,{default:d(()=>[r(V,{type:"warning",onClick:l[3]||(l[3]=n=>o.getdsl(1))},{default:d(()=>[O("查看查询语句 ")]),_:1}),r(V,{type:"success",onClick:l[4]||(l[4]=n=>o.search(1))},{default:d(()=>[O("查询 ")]),_:1}),r(V,{disabled:t.downloadLoading,loading:t.downloadLoading,type:"primary",onClick:l[5]||(l[5]=n=>o.download())},{default:d(()=>[O("下载 ")]),_:1},8,["disabled","loading"])]),_:1})]),_:1})]),ne((i(),y(A,{height:"400",data:t.tableData,"use-virtual":"",ref:"multipleTable","row-height":30,"show-overflow-tooltip":""},{default:d(()=>[r(m,{label:"ID",align:"center",fixed:"","min-width":"100",prop:"_id"}),(i(!0),h(Y,null,G(t.tableHeader,(n,_)=>(i(),h(Y,null,[n!="xwl_index"&&n!="_id"?(i(),y(m,{key:0,prop:n,sortable:!0,align:"center",label:n},null,8,["prop","label"])):U("",!0)],64))),256)),r(m,{align:"center",label:e.$t("操作"),fixed:"right",width:"200"},{default:d(n=>[r(g,null,{default:d(()=>[r(V,{type:"primary",onClick:_=>o.look(n.$index)},{default:d(()=>[O(" 编辑 ")]),_:2},1032,["onClick"]),r(V,{type:"danger",onClick:_=>o.deleteByID(n.row,n.$index)},{default:d(()=>[O(j(e.$t("删除")),1)]),_:2},1032,["onClick"])]),_:2},1024)]),_:1},8,["label"])]),_:1},8,["data"])),[[D,t.tableLoading]]),r(R,{"current-page":t.input.page,"page-sizes":[10,20,30,50,100,150,200],"page-size":t.input.limit,layout:"total, sizes, prev, pager, next, jumper",total:t.count,onSizeChange:o.handleSizeChange,onCurrentChange:o.search},null,8,["current-page","page-size","total","onSizeChange","onCurrentChange"])]),_:1}),r(H,{ref:"drawer",title:"详细数据","before-close":o.drawerHandleClose,modelValue:t.drawerShow,"onUpdate:modelValue":l[7]||(l[7]=n=>t.drawerShow=n),direction:"rtl","close-on-press-escape":"","destroy-on-close":"",size:"80%"},{default:d(()=>[r(F,null,{default:d(()=>[O("操作")]),_:1}),r(V,{type:"primary",onClick:o.updateByID},{default:d(()=>[O("修改 ")]),_:1},8,["onClick"]),t.showJsonEditor1?(i(),y(E,{key:0,value:o.jsonDataString,"onUpdate:value":l[6]||(l[6]=n=>o.jsonDataString=n),height:"900",class:"res-body",styles:"width: 100%",title:"详细数据",onGetValue:o.getEditDoc},null,8,["value","onGetValue"])):U("",!0)]),_:1},8,["before-close","modelValue"]),r(H,{ref:"drawer",title:"查询DSL","before-close":o.drawerHandleClose,modelValue:t.queryDslShow,"onUpdate:modelValue":l[9]||(l[9]=n=>t.queryDslShow=n),direction:"rtl","close-on-press-escape":"","destroy-on-close":"",size:"80%"},{default:d(()=>[t.showJsonEditor2?(i(),y(E,{key:0,value:o.queryDslString,"onUpdate:value":l[8]||(l[8]=n=>o.queryDslString=n),height:"900",class:"res-body",styles:"width: 100%",title:"DSL",read:!0},null,8,["value"])):U("",!0)]),_:1},8,["before-close","modelValue"])],64)):(i(),h("div",Ct,[r(X,{content:"请先点击左侧索引名称"})]))])])}const Mt=Q(yt,[["render",Vt],["__scopeId","data-v-94b0c470"]]);const Yt={style:{display:"flex","justify-content":"space-between"}},jt={class:"content_xwl"},Lt={id:"scollL",style:{height:"100%",width:"100%",display:"inline-block","vertical-align":"top"}},Tt={style:{width:"100%",height:"calc(100% - 80px)","overflow-x":"hidden","overflow-y":"auto",padding:"10px"}},Nt={class:"el-dropdown-link"},Ot={slot:"title"},Ut={style:{width:"100%",height:"calc(100% - 80px)","overflow-x":"hidden","overflow-y":"auto",padding:"10px"}},It={style:{height:"100%",width:"100%",display:"inline-block","vertical-align":"top"}},At={__name:"index",setup(e){ge().appContext.config.globalProperties;const l=S([]),a=S({}),s=S(!1),t=S(!0),o=S("settings"),w=S(!1),f=S([]),C=S(""),p=S(""),L=S({}),V=S([]),I=S(!1),N=S([]),v=le(()=>C.value===""?f.value:xe(f.value,C.value)),m=le(()=>JSON.stringify(L.value,null,"	")),g=le(()=>JSON.stringify(a.value,null,"	")),A=()=>{t.value=!1,ve(()=>{t.value=!0})},R=async D=>{p.value=D;const n={es_connect:z.GetSelectEsConnID(),index_name:p.value};s.value=!0,console.log("123");try{const _=await fe(n);if(_.code!==0){s.value=!1,x.error(_.msg);return}const b=await Se(n);if(b.code!==0){s.value=!1,x.error(b.msg);return}a.value=_.data.list;const M=[{label:"筛选字段",options:[]}],q={2:[]};let J={};switch(_.data.ver){case 6:J=a.value[p.value].mappings[Object.keys(a.value[p.value].mappings)[0]].properties;break;case 7:case 8:J=a.value[p.value].mappings.properties;break}const Z=1,ee=2,te=3,W=5;for(const $ in J)if(J[$].type){M[0].options.push({value:$,label:$});const P={attribute_name:$,show_name:$};switch(J[$].type){case"text":P.data_type=W;break;case"keyword":P.data_type=te;break;case"byte":case"short":case"integer":case"long":P.data_type=Z;break;case"float":case"half_float":case"scaled_float":case"double":P.data_type=ee;break}q[2].push(P)}N.value=q,V.value=M,L.value=b.data,s.value=!1,A()}catch(_){s.value=!1,x.error("An error occurred while fetching data.",_)}},T=D=>{const n=f.value.findIndex(_=>_.k.toString()===D.toString());n!==-1&&f.value.splice(n,1)},F=async(D,n)=>{try{await be.confirm("删除该索引不可恢复，确认删除吗？",$t("警告"),{confirmButtonText:$t("确定"),cancelButtonText:$t("取消"),type:"warning"})}catch{return}const _=await Ce({es_connect:z.GetSelectEsConnID(),index_name:n});_.code!==0?x({type:"error",message:_.msg}):(T(D),x({type:"success",message:_.msg}))};ze(async()=>{E()});const E=async()=>{const D={cat:"CatIndices",es_connect:z.GetSelectEsConnID()};f.value=[],w.value=!0,l.value=[];const n=await ke(D);if(n.code!=0){x({type:"error",message:n.msg}),w.value=!1;return}for(const _ in n.data){const b=n.data[_],M={health:b.health,index:b.index,k:_,storeSize:b["store.size"],docsCount:b["docs.count"]};f.value.push(M),l.value.push({value:b.index,data:b.index})}w.value=!1},H=(D,n)=>{const _=D?l.value.filter(X(D)):l.value;n(_)},X=D=>n=>n.value.toLowerCase().includes(D.toLowerCase());return(D,n)=>{const _=c("el-autocomplete"),b=c("Grid"),M=c("el-icon"),q=c("el-dropdown-item"),J=c("el-dropdown-menu"),Z=c("el-dropdown"),ee=c("el-menu-item"),te=c("el-menu"),W=c("el-tab-pane"),$=c("el-tabs"),P=c("el-drawer"),ue=we("loading");return i(),h("div",Yt,[u("div",jt,[r(B(Ue),{style:{"margin-top":"3px"},distribute:.14,lineThickness:6,isVertical:!0},{first:d(()=>[u("div",Lt,[u("div",Tt,[r(_,{modelValue:C.value,"onUpdate:modelValue":n[0]||(n[0]=k=>C.value=k),style:{width:"90%"},clearable:"","fetch-suggestions":H,placeholder:D.$t("请输入索引名")},{default:d(({item:k})=>[u("span",null,j(k.value),1)]),_:1},8,["modelValue","placeholder"]),ne((i(),y(te,{style:{"margin-top":"10px"}},{default:d(()=>[(i(!0),h(Y,null,G(v.value,(k,ae)=>(i(),y(ee,{index:ae.toString(),onClick:ce=>R(k.index)},{default:d(()=>[r(Z,null,{dropdown:d(()=>[r(J,null,{default:d(()=>[r(q,{onClick:ce=>F(v.value[ae].k,v.value[ae].index)},{default:d(()=>[O(j(D.$t("删除")),1)]),_:2},1032,["onClick"]),r(q,{onClick:n[1]||(n[1]=ce=>I.value=!0)},{default:d(()=>[O(j(D.$t("结构")),1)]),_:1})]),_:2},1024)]),default:d(()=>[u("span",Nt,[k.health=="red"?(i(),y(M,{key:0,style:{color:"red"}},{default:d(()=>[r(b)]),_:1})):U("",!0),k.health=="green"?(i(),y(M,{key:1,style:{color:"#13ce66"}},{default:d(()=>[r(b)]),_:1})):U("",!0),k.health=="yellow"?(i(),y(M,{key:2,style:{color:"#ffba00"}},{default:d(()=>[r(b)]),_:1})):U("",!0)])]),_:2},1024),u("span",Ot,j(k.index)+"【"+j(k.storeSize)+"】",1)]),_:2},1032,["index","onClick"]))),256))]),_:1})),[[ue,w.value]])])])]),second:d(()=>[u("div",Ut,[t.value?(i(),y(Mt,{key:0,"attr-map-prop":N.value,"event-attr-options-prop":V.value,"index-name":p.value},null,8,["attr-map-prop","event-attr-options-prop","index-name"])):U("",!0)]),r(P,{ref:"drawer",title:D.$t("索引结构"),modelValue:I.value,"onUpdate:modelValue":n[5]||(n[5]=k=>I.value=k),direction:"rtl","close-on-press-escape":"","destroy-on-close":"",size:"80%"},{default:d(()=>[u("div",It,[ne((i(),y($,{modelValue:o.value,"onUpdate:modelValue":n[4]||(n[4]=k=>o.value=k),type:"border-card"},{default:d(()=>[r(W,{lazy:!0,label:D.$t("索引设置"),name:"settings"},{default:d(()=>[t.value?(i(),y(se,{key:0,value:m.value,"onUpdate:value":n[2]||(n[2]=k=>m.value=k),height:"720",styles:"width: 100%",read:!0,title:p.value},null,8,["value","title"])):U("",!0)]),_:1},8,["label"]),r(W,{lazy:!0,label:D.$t("映射结构"),name:"mappings"},{default:d(()=>[t.value?(i(),y(se,{key:0,value:g.value,"onUpdate:value":n[3]||(n[3]=k=>g.value=k),height:"720",styles:"width: 100%",read:!0,title:p.value},null,8,["value","title"])):U("",!0)]),_:1},8,["label"])]),_:1},8,["modelValue"])),[[ue,s.value]])])]),_:1},8,["title","modelValue"])]),_:1})])])}}},Ht=Q(At,[["__scopeId","data-v-b6e0f9a8"]]);export{Ht as default};