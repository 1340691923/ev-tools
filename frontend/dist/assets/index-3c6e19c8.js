var x=Object.defineProperty;var ee=(a,o,u)=>o in a?x(a,o,{enumerable:!0,configurable:!0,writable:!0,value:u}):a[o]=u;var w=(a,o,u)=>(ee(a,typeof o!="symbol"?o+"":o,u),u);import{C as V,t as j,V as T,d as U}from"./index-38404389.js";import{C as L,f as te}from"./es-382a2972.js";import{s as P,E as M,g as W,r as d,z as q,o as i,c as B,j as k,a as e,w as t,h as b,t as g,A as J,k as h,F as K,q as Q,x as C}from"./index-15b0a2bf.js";import{_ as G}from"./_plugin-vue_export-helper-c27b6911.js";var ae=Object.defineProperty,le=Object.getOwnPropertyDescriptor,H=(a,o,u,l)=>{for(var n=l>1?void 0:l?le(o,u):o,_=a.length-1,c;_>=0;_--)(c=a[_])&&(n=(l?c(o,u,n):c(n))||n);return l&&n&&ae(o,u,n),n};let E=class extends T{constructor(){super(...arguments);w(this,"catType","");w(this,"tableInfo",[]);w(this,"ctx",null);w(this,"total",0);w(this,"connectLoading",!1);w(this,"page",1);w(this,"limit",10);w(this,"pageshow",!0);w(this,"list",[]);w(this,"input","");w(this,"allList",[])}pageLimit(){this.list=this.allList.filter((u,l)=>l<this.page*this.limit&&l>=this.limit*(this.page-1))}search(){this.page=1,this.pageshow=!1,this.searchData(),this.$nextTick(()=>{this.pageshow=!0})}handleSizeChange(u){this.limit=u,this.pageLimit()}handleCurrentChange(u){this.page=u,this.pageLimit()}searchData(){this.connectLoading=!0;const u={cat:this.catType,es_connect:P.GetSelectEsConnID()};L(u).then(l=>{if(l.code==0){let n=l.data;if(n==null)return;for(const _ in n){const c=n[_];for(const s in c){let m=Number(c[s]);isNaN(m)&&(m=c[s]),n[_][s.split(".").join("->")]=m}}n=te(n,this.input.trim()),this.allList=n,this.total=n.length,this.pageLimit()}else M.error({type:"error",message:l.msg});this.connectLoading=!1}).catch(l=>{console.log(l),this.connectLoading=!1})}mounted(){this.searchData()}};H([U],E.prototype,"catType",2);H([U],E.prototype,"tableInfo",2);E=H([V({name:"CatIndex",setup(){return{ctx:W().appContext.config.globalProperties}}})],E);const ne=j(E),se={class:"app-container"},oe={class:"search-container"};function ue(a,o,u,l,n,_){const c=d("el-input"),s=d("el-form-item"),m=d("el-button"),p=d("el-form"),f=d("el-table-column"),A=d("el-table"),D=d("el-pagination"),N=d("el-card"),I=q("loading");return i(),B("div",null,[k("div",se,[k("div",oe,[e(p,{inline:!0},{default:t(()=>[e(s,{label:"关键词"},{default:t(()=>[e(c,{modelValue:a.input,"onUpdate:modelValue":o[0]||(o[0]=y=>a.input=y),style:{width:"7rem"},clearable:"",onInput:a.search},null,8,["modelValue","onInput"])]),_:1}),e(s,{label:""},{default:t(()=>[e(m,{type:"primary",onClick:a.search},{default:t(()=>[b(g(a.$t("搜索")),1)]),_:1},8,["onClick"])]),_:1})]),_:1})]),e(N,{shadow:"never",class:"table-container"},{default:t(()=>[J((i(),h(A,{data:a.list},{default:t(()=>[e(f,{label:a.$t("序号"),align:"center",fixed:"",width:"100"},{default:t(y=>[b(g(y.$index+1),1)]),_:1},8,["label"]),(i(!0),B(K,null,Q(a.tableInfo,(y,$)=>(i(),h(f,{"show-overflow-tooltip":"",key:$,align:"center",label:a.tableInfo[$].desc,width:y.width,prop:y.data.toString(),sortable:y.sort},{default:t(z=>[b(g(z.row[y.data.split("->").join(".")]),1)]),_:2},1032,["label","width","prop","sortable"]))),128))]),_:1},8,["data"])),[[I,a.connectLoading]]),a.pageshow?(i(),h(D,{key:0,class:"pagination-container","current-page":a.page,"page-sizes":[10,20,30,50,100,150,200,500,1e3],"page-size":a.limit,layout:"total, sizes, prev, pager, next, jumper",total:a.total,onSizeChange:a.handleSizeChange,onCurrentChange:a.handleCurrentChange},null,8,["current-page","page-size","total","onSizeChange","onCurrentChange"])):C("",!0)]),_:1})])])}const re=G(ne,[["render",ue]]);var de=Object.defineProperty,ie=Object.getOwnPropertyDescriptor,ce=(a,o,u,l)=>{for(var n=l>1?void 0:l?ie(o,u):o,_=a.length-1,c;_>=0;_--)(c=a[_])&&(n=(l?c(o,u,n):c(n))||n);return l&&n&&de(o,u,n),n};let O=class extends T{constructor(){super(...arguments);w(this,"list",[]);w(this,"connectLoading",!1);w(this,"value",90)}beforeMount(){this.searchData()}getOpt(u){const l=Number(u);return l<60?"#4ad47f":l>=60&&l<=80?"#ff9800":"red"}async searchData(){const u=W().appContext.config.globalProperties;console.log("ctx",u),this.connectLoading=!0;const l=await L({cat:"Node",es_connect:P.GetSelectEsConnID()}),n=await L({cat:"CatAllocation",es_connect:P.GetSelectEsConnID()});if(l.code!=0){M.error({type:"error",message:l.msg}),this.connectLoading=!1;return}for(const _ in l.data){const c=l.data[_];for(const s in n.data){const m=n.data[s];if(m.node==c.name){l.data[_]["disk.indices"]=m["disk.indices"],l.data[_].shards=m.shards,l.data[_]["disk.avail"]=m["disk.avail"];break}}}this.list=l.data,this.connectLoading=!1}};O=ce([V({name:"Node"})],O);const pe=j(O);const _e={class:"app-container"},fe={class:"search-container"};function me(a,o,u,l,n,_){const c=d("el-button"),s=d("el-form-item"),m=d("el-form"),p=d("el-tag"),f=d("el-col"),A=d("Star"),D=d("el-icon"),N=d("el-tooltip"),I=d("Tickets"),y=d("StarFilled"),$=d("CreditCard"),z=d("Postcard"),R=d("Wallet"),v=d("el-row"),F=d("el-progress"),X=d("el-card"),Y=q("loading");return J((i(),B("div",_e,[k("div",fe,[e(m,null,{default:t(()=>[e(s,null,{default:t(()=>[e(c,{type:"primary",onClick:a.searchData},{default:t(()=>[b("刷新 ")]),_:1},8,["onClick"])]),_:1})]),_:1})]),e(v,{gutter:40},{default:t(()=>[(i(!0),B(K,null,Q(a.list,(r,De,Z)=>(i(),h(f,{key:Z,xs:40,sm:11,lg:11,class:"card-panel-col"},{default:t(()=>[e(X,{shadow:"never",class:"table-container"},{default:t(()=>[e(m,{class:"responsive-form"},{default:t(()=>[e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"节点名:"},{default:t(()=>[e(p,null,{default:t(()=>[b("节点名:"+g(r.name),1)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:""},{default:t(()=>[k("div",null,[r.master==="*"?(i(),h(N,{key:0,content:"主节点",placement:"top",style:{cursor:"pointer"}},{default:t(()=>[e(p,{type:"primary"},{default:t(()=>[e(D,null,{default:t(()=>[e(A)]),_:1})]),_:1})]),_:1})):r["node.role"].includes("m")&&r["node.role"].includes("v")?(i(),h(N,{key:1,content:"仅选举节点",placement:"top",style:{cursor:"pointer"}},{default:t(()=>[e(p,{type:"primary"},{default:t(()=>[e(D,null,{default:t(()=>[e(I)]),_:1})]),_:1})]),_:1})):r["node.role"].includes("m")?(i(),h(N,{key:2,content:"主节点候选",placement:"top",style:{cursor:"pointer"}},{default:t(()=>[e(p,{type:"primary"},{default:t(()=>[e(D,null,{default:t(()=>[e(y)]),_:1})]),_:1})]),_:1})):C("",!0),r["node.role"].includes("d")?(i(),h(N,{key:3,content:"数据节点",placement:"top",style:{cursor:"pointer"}},{default:t(()=>[e(p,{type:"success"},{default:t(()=>[e(D,null,{default:t(()=>[e($)]),_:1})]),_:1})]),_:1})):C("",!0),r["node.role"].includes("i")?(i(),h(N,{key:4,content:"预处理节点",placement:"top",style:{cursor:"pointer"}},{default:t(()=>[e(p,{type:"warning"},{default:t(()=>[e(D,null,{default:t(()=>[e(z)]),_:1})]),_:1})]),_:1})):C("",!0),r["node.role"]==="-"?(i(),h(N,{key:5,content:"仅协调节点",placement:"top",style:{cursor:"pointer"}},{default:t(()=>[e(p,{type:"warning"},{default:t(()=>[e(D,null,{default:t(()=>[e(R)]),_:1})]),_:1})]),_:1})):C("",!0)])]),_:2},1024)]),_:2},1024)]),_:2},1024),e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"IP:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r.ip),1)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:"主节点:"},{default:t(()=>[k("div",null,[r.master?(i(),h(p,{key:0},{default:t(()=>[b("yes")]),_:1})):r.masterEligible?(i(),h(p,{key:1},{default:t(()=>[b("eligible")]),_:1})):(i(),h(p,{key:2},{default:t(()=>[b("no")]),_:1}))])]),_:2},1024)]),_:2},1024)]),_:2},1024),e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"节点角色:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r["node.role"]),1)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:"负载:"},{default:t(()=>[e(p,null,{default:t(()=>[b("1m:"+g(r.load_1m)+" / 5m:"+g(r.load_5m)+" / 15m:"+g(r.load_15m),1)]),_:2},1024)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"内存:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r["ram.current"])+"/"+g(r["ram.max"]),1)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:"堆内存:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r["heap.current"])+"/ "+g(r["heap.max"]),1)]),_:2},1024)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"磁盘:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r["disk.used"])+"/ "+g(r["disk.total"]),1)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:"磁盘可用:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r["disk.avail"]),1)]),_:2},1024)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"分片数:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r.shards),1)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:"索引所占空间大小:"},{default:t(()=>[e(p,null,{default:t(()=>[b(g(r["disk.indices"]),1)]),_:2},1024)]),_:2},1024)]),_:2},1024)]),_:2},1024),e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"CPU:"},{default:t(()=>[e(F,{"stroke-width":24,color:a.getOpt,"text-inside":!0,style:{"margin-right":"1px",width:"8rem"},percentage:Number(r.cpu)},null,8,["color","percentage"])]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:"内存:"},{default:t(()=>[e(F,{"stroke-width":24,color:a.getOpt,"text-inside":!0,style:{"margin-right":"1px",width:"8rem"},percentage:Number(r["ram.percent"])},null,8,["color","percentage"])]),_:2},1024)]),_:2},1024)]),_:2},1024),e(v,{gutter:20},{default:t(()=>[e(f,{span:12},{default:t(()=>[e(s,{label:"堆内存:"},{default:t(()=>[e(F,{"stroke-width":24,color:a.getOpt,"text-inside":!0,style:{"margin-right":"1px",width:"8rem"},percentage:Number(r["heap.percent"])},null,8,["color","percentage"])]),_:2},1024)]),_:2},1024),e(f,{span:12},{default:t(()=>[e(s,{label:"磁盘:"},{default:t(()=>[e(F,{"stroke-width":24,color:a.getOpt,"text-inside":!0,style:{width:"8rem"},percentage:Number(r["disk.used_percent"])},null,8,["color","percentage"])]),_:2},1024)]),_:2},1024)]),_:2},1024)]),_:2},1024)]),_:2},1024)]),_:2},1024))),128))]),_:1})])),[[Y,a.connectLoading]])}const he=G(pe,[["render",me],["__scopeId","data-v-8b6988b3"]]);var ge=Object.defineProperty,we=Object.getOwnPropertyDescriptor,be=(a,o,u,l)=>{for(var n=l>1?void 0:l?we(o,u):o,_=a.length-1,c;_>=0;_--)(c=a[_])&&(n=(l?c(o,u,n):c(n))||n);return l&&n&&ge(o,u,n),n};let S=class extends T{constructor(){super(...arguments);w(this,"activeName","Node");w(this,"catData",{CatIndices:[{data:"health",desc:"索引健康状态",sort:!0,width:100},{data:"status",desc:"索引的开启状态",sort:!0,width:100},{data:"index",desc:"索引名称",sort:!0,width:0},{data:"uuid",desc:"索引uuid",sort:!0,width:300},{data:"pri",desc:"索引主分片数",sort:!0,width:100},{data:"rep",desc:"索引副本分片数量",sort:!0,width:100},{data:"docs->count",desc:"索引中文档总数",sort:!0,width:100},{data:"docs->deleted",desc:"索引中删除状态的文档",sort:!0,width:100},{data:"store->size",desc:"主分片+副本分分片的大小",sort:!0,width:100},{data:"pri->store->size",desc:"主分片的大小",sort:!0,width:200}],CatAliases:[{data:"alias",desc:"别名",sort:!0,width:300},{data:"index",desc:"索引别名指向",sort:!0,width:0},{data:"filter",desc:"过滤器",sort:!0,width:220},{data:"index",desc:"索引路由",sort:!0,width:220},{data:"routing->search",desc:"搜索路由",sort:!0,width:220},{data:"is_write_index",desc:"写索引",sort:!0,width:220}],CatAllocation:[{data:"host",desc:"节点host",sort:!0,width:150},{data:"ip",desc:"节点ip",sort:!0,width:150},{data:"node",desc:"节点名称",sort:!0,width:0},{data:"shards",desc:"节点承载分片数量",sort:!0,width:150},{data:"disk->indices",desc:"索引占用空间大小",sort:!0,width:150},{data:"disk->used",desc:"节点所在机器已使用的磁盘空间大小",sort:!0,width:150},{data:"disk->avail",desc:"节点可用空间大小",sort:!0,width:170},{data:"disk->total",desc:"节点总空间大小",sort:!0,width:170},{data:"disk->percent",desc:"节点磁盘占用百分比",sort:!0,width:100}],CatCount:[{data:"epoch",desc:"自标准时间（1970-01-01 00:00:00）以来的秒数",sort:!0,width:0},{data:"timestamp",desc:"时分秒,utc时区",sort:!0,width:500},{data:"count",desc:"文档总数",sort:!0,width:500}],CatHealth:[{data:"cluster",desc:"集群名称",sort:!0,width:0},{data:"status",desc:"集群状态",sort:!0,width:100},{data:"node->total",desc:"节点总数",sort:!0,width:100},{data:"node->data",desc:"数据节点总数",sort:!0,width:100},{data:"shards",desc:"分片总数",sort:!0,width:100},{data:"pri",desc:"主分片总数",sort:!0,width:100},{data:"relo",desc:"复制节点总数",sort:!0,width:100},{data:"init",desc:"初始化节点总数",sort:!0,width:100},{data:"unassign",desc:"未分配分片总数",sort:!0,width:100},{data:"pending_tasks",desc:"待定任务总数",sort:!0,width:100},{data:"max_task_wait_time",desc:"等待最长任务的等待时间",sort:!0,width:100},{data:"active_shards_percent",desc:"活动分片百分比",sort:!0,width:100},{data:"epoch",desc:"自标准时间（1970-01-01 00:00:00）以来的秒数",sort:!0,width:0},{data:"timestamp",desc:"时分秒,utc时区",sort:!0,width:100}],CatShards:[{data:"index",desc:"索引名称",sort:!0,width:0},{data:"shard",desc:"分片序号",sort:!0,width:100},{data:"prirep",desc:"分片类型(p为主分片，r为复制分片)",sort:!0,width:220},{data:"state",desc:"分片状态",sort:!0,width:100},{data:"docs",desc:"该分片存放的文档数量",sort:!0,width:140},{data:"store",desc:"该分片占用的存储空间大小",sort:!0,width:200},{data:"ip",desc:"该分片所在的服务器ip",sort:!1,width:200},{data:"node",desc:"该分片所在的节点名称",sort:!0,width:200}]})}};S=be([V({name:"Cat",components:{CatIndex:re,Node:he}})],S);const Ce=j(S),ye={class:"app-container"};function ve(a,o,u,l,n,_){const c=d("node"),s=d("el-tab-pane"),m=d("cat-index"),p=d("el-tabs");return i(),B("div",ye,[e(p,{modelValue:a.activeName,"onUpdate:modelValue":o[0]||(o[0]=f=>a.activeName=f)},{default:t(()=>[e(s,{lazy:!0,label:"节点",name:"Node"},{default:t(()=>[e(c)]),_:1}),e(s,{lazy:!0,label:"查看索引信息",name:"CatIndices"},{default:t(()=>[a.activeName=="CatIndices"?(i(),h(m,{key:0,"cat-type":a.activeName,"table-info":a.catData[a.activeName]},null,8,["cat-type","table-info"])):C("",!0)]),_:1}),e(s,{lazy:!0,label:"显示别名,过滤器,路由信息",name:"CatAliases"},{default:t(()=>[a.activeName=="CatAliases"?(i(),h(m,{key:0,"cat-type":a.activeName,"table-info":a.catData[a.activeName]},null,8,["cat-type","table-info"])):C("",!0)]),_:1}),e(s,{lazy:!0,label:"显示每个节点分片数量、占用空间",name:"CatAllocation"},{default:t(()=>[a.activeName=="CatAllocation"?(i(),h(m,{key:0,"cat-type":a.activeName,"table-info":a.catData[a.activeName]},null,8,["cat-type","table-info"])):C("",!0)]),_:1}),e(s,{lazy:!0,label:"显示索引文档的数量",name:"CatCount"},{default:t(()=>[a.activeName=="CatCount"?(i(),h(m,{key:0,"cat-type":a.activeName,"table-info":a.catData[a.activeName]},null,8,["cat-type","table-info"])):C("",!0)]),_:1}),e(s,{lazy:!0,label:"查看集群健康状况",name:"CatHealth"},{default:t(()=>[a.activeName=="CatHealth"?(i(),h(m,{key:0,"cat-type":a.activeName,"table-info":a.catData[a.activeName]},null,8,["cat-type","table-info"])):C("",!0)]),_:1}),e(s,{lazy:!0,label:"显示索引分片信息",name:"CatShards"},{default:t(()=>[a.activeName=="CatShards"?(i(),h(m,{key:0,"cat-type":a.activeName,"table-info":a.catData[a.activeName]},null,8,["cat-type","table-info"])):C("",!0)]),_:1})]),_:1},8,["modelValue"])])}const Ie=G(Ce,[["render",ve]]);export{Ie as default};