var H=Object.defineProperty;var I=(e,l,s)=>l in e?H(e,l,{enumerable:!0,configurable:!0,writable:!0,value:s}):e[l]=s;var m=(e,l,s)=>(I(e,typeof l!="symbol"?l+"":l,s),s);import{s as v,E as w,r,o as f,c as D,a as t,w as a,h as N,i as y,t as _,H as k,g as G,z as T,k as V,x as B,A as J,F as M,q as P}from"./index-e47c9f2e.js";import{S as q,C as K,a as Q,b as F}from"./es-backup-2aa89d8a.js";import{_ as U}from"./_plugin-vue_export-helper-c27b6911.js";import{t as W,V as X,C as Y}from"./index-bffeb1e5.js";import{J as Z}from"./index-efe8a954.js";const x={name:"Add",components:{},props:{open:{type:Boolean,default:!1},snapshotData:{type:Object,default:null}},data(){return{props_data:{open:this.open},isOpen:!1,form:{name:"",type:"fs",location:"",compress:"false",max_restore_bytes_per_sec:"40mb",max_snapshot_bytes_per_sec:"40mb",chunk_size:"",readonly:"false"}}},computed:{},created(){this.snapshotData!=null&&(this.form=Object.assign({},this.snapshotData))},methods:{closeDialog(){this.$emit("close",!1)},async confirm(){const e=this.form;e.es_connect=v.GetSelectEsConnID();const{code:l,data:s,msg:i}=await q(e);if(l==0){this.$emit("close",!0),w({type:"success",message:i});return}else{w({type:"error",message:i});return}}}},ee={style:{"text-align":"right"}};function le(e,l,s,i,o,p){const c=r("el-input"),b=r("el-form-item"),h=r("el-option"),g=r("el-select"),$=r("el-form"),d=r("el-button"),E=r("el-card"),S=r("el-dialog");return f(),D("div",null,[t(S,{"close-on-click-modal":!1,modelValue:o.props_data.open,"onUpdate:modelValue":l[8]||(l[8]=u=>o.props_data.open=u),title:e.$t("新增/修改快照存储库"),onClose:p.closeDialog},{default:a(()=>[t(E,{class:"box-card"},{default:a(()=>[t($,{"label-width":"300px","label-position":"left"},{default:a(()=>[t(b,{label:e.$t("存储库名")},{default:a(()=>[t(c,{modelValue:o.form.name,"onUpdate:modelValue":l[0]||(l[0]=u=>o.form.name=u),placeholder:e.$t("存储库名")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),t(b,{label:e.$t("类型（type）")},{default:a(()=>[t(g,{modelValue:o.form.type,"onUpdate:modelValue":l[1]||(l[1]=u=>o.form.type=u),filterable:""},{default:a(()=>[t(h,{label:"fs",value:"fs"}),t(h,{label:"url",value:"url"})]),_:1},8,["modelValue"])]),_:1},8,["label"]),t(b,{label:e.$t("存储位置（location/url）")},{default:a(()=>[t(c,{modelValue:o.form.location,"onUpdate:modelValue":l[2]||(l[2]=u=>o.form.location=u),type:"textarea",autosize:{minRows:2,maxRows:4},placeholder:e.$t("存储位置（location/url）")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),t(b,{label:e.$t("是否压缩 (compress)")},{default:a(()=>[t(g,{modelValue:o.form.compress,"onUpdate:modelValue":l[3]||(l[3]=u=>o.form.compress=u),clearable:"",filterable:""},{default:a(()=>[t(h,{label:e.$t("是"),value:"true"},null,8,["label"]),t(h,{label:e.$t("否"),value:"false"},null,8,["label"])]),_:1},8,["modelValue"])]),_:1},8,["label"]),t(b,{label:e.$t("大文件分解块大小")},{default:a(()=>[t(c,{modelValue:o.form.chunk_size,"onUpdate:modelValue":l[4]||(l[4]=u=>o.form.chunk_size=u),placeholder:e.$t("1GB，10MB，5KB，500B。默认为null（无限制的块大小）")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),t(b,{label:e.$t("是否只读")},{default:a(()=>[t(g,{modelValue:o.form.readonly,"onUpdate:modelValue":l[5]||(l[5]=u=>o.form.readonly=u),filterable:""},{default:a(()=>[t(h,{label:e.$t("是"),value:"true"},null,8,["label"]),t(h,{label:e.$t("否"),value:"false"},null,8,["label"])]),_:1},8,["modelValue"])]),_:1},8,["label"]),t(b,{label:e.$t("节点恢复速率")},{default:a(()=>[t(c,{modelValue:o.form.max_restore_bytes_per_sec,"onUpdate:modelValue":l[6]||(l[6]=u=>o.form.max_restore_bytes_per_sec=u),placeholder:e.$t("节点恢复速率")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),t(b,{label:e.$t("每个节点快照速率")},{default:a(()=>[t(c,{modelValue:o.form.max_snapshot_bytes_per_sec,"onUpdate:modelValue":l[7]||(l[7]=u=>o.form.max_snapshot_bytes_per_sec=u),placeholder:e.$t("每个节点快照速率")},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),N("div",ee,[t(d,{type:"danger",onClick:p.closeDialog},{default:a(()=>[y(_(e.$t("取消")),1)]),_:1},8,["onClick"]),t(d,{type:"primary",onClick:p.confirm},{default:a(()=>[y(_(e.$t("确认")),1)]),_:1},8,["onClick"])])]),_:1})]),_:1},8,["modelValue","title","onClose"])])}const te=U(x,[["render",le]]);var oe=Object.defineProperty,ae=Object.getOwnPropertyDescriptor,se=(e,l,s,i)=>{for(var o=i>1?void 0:i?ae(l,s):l,p=e.length-1,c;p>=0;p--)(c=e[p])&&(o=(i?c(l,s,o):c(o))||o);return i&&o&&oe(l,s,o),o};let A=class extends X{constructor(){super(...arguments);m(this,"ctx",null);m(this,"snapshotData",null);m(this,"openAddDialog",!1);m(this,"loading",!1);m(this,"snapshotNameList",[]);m(this,"resData",{});m(this,"jsonData","{}");m(this,"name","");m(this,"drawerShow",!1);m(this,"tableData",[]);m(this,"index",0);m(this,"tableHeader",[])}mounted(){this.initSnapshotName(),this.search()}async CleanupeRepository(l){const s={};s.es_connect=v.GetSelectEsConnID(),s.name=l;const{data:i,code:o,msg:p}=await K(s);if(o!=0){w.error({type:"error",message:p});return}else{this.initSnapshotName(),this.search(),w.success({type:"success",message:p});return}}async SnapshotDeleteRepositoryAction(l){const s={};s.es_connect=v.GetSelectEsConnID(),s.name=l;const{data:i,code:o,msg:p}=await Q(s);if(o!=0){w.error({type:"error",message:p});return}else{this.initSnapshotName(),this.search(),w.success({type:"success",message:p});return}}updateSnapshotData(l){this.snapshotData=l,this.openAddDialog=!0}closeAddDialog(l){this.openAddDialog=!1,this.snapshotData=null,l&&(this.initSnapshotName(),this.search())}async initSnapshotName(){this.loading=!0;const l={};l.es_connect=v.GetSelectEsConnID();const{data:s,code:i,msg:o}=await F(l);i==0?(this.resData=s.res,k({title:"成功",dangerouslyUseHTMLString:!0,message:`
<strong> <b style="color: orange">您设置的path.repo为${s.pathRepo.join(",")}</b><br></strong>
          `,type:"success"})):i==199999?k({title:"Error",dangerouslyUseHTMLString:!0,message:`
<strong>
            <i style="color: orange">path.repo没有设置</i><br>
<i>在elasticsearch.yml 配置文件中配置仓库base目录</i><br>
<i>添加path.repo: /tmp/tmp (/tmp/tmp 为快照备份所在文件夹, <br><i style="color: orange">注意</i>首先要先创建这个文件夹)</i>

          `,type:"error"}):w.error({type:"error",message:o}),this.loading=!1}async search(){const l={};l.es_connect=v.GetSelectEsConnID(),l.snapshot_info_list=this.snapshotNameList,l.repository="test3";const{data:s,code:i,msg:o}=await F(l);if(i!=0){if(i!=199999){w.error({type:"error",message:o});return}}else i===199999?k({title:"Error",dangerouslyUseHTMLString:!0,message:`
        <strong>
          <i style="color: orange">path.repo没有设置</i><br>
          <i>在elasticsearch.yml 配置文件中配置仓库base目录</i><br>
          <i>添加path.repo: /tmp/tmp (/tmp/tmp 为快照备份所在文件夹, <br><i style="color: orange">注意</i>首先要先创建这个文件夹)</i>
        </strong>
      `,type:"error"}):this.tableData=s.list}openDetail(l,s,i){this.name=l.name,this.drawerShow=!0}look(l){this.name=l,this.jsonData=JSON.stringify(this.resData[l],null,"	"),this.drawerShow=!0}drawerHandleClose(l){l()}};A=se([Y({name:"Task",components:{Add:te,JsonEditor:Z},setup(){return{ctx:G().appContext.config.globalProperties}}})],A);const ne=W(A);const re={class:"app-container"},ie={class:"search-container"},ue={key:0},pe={key:1},de={key:0},me={key:1};function ce(e,l,s,i,o,p){const c=r("el-option"),b=r("el-select"),h=r("el-form-item"),g=r("el-button"),$=r("el-form"),d=r("el-table-column"),E=r("el-button-group"),S=r("el-table"),u=r("el-card"),j=r("json-editor"),R=r("el-drawer"),O=r("add"),L=T("loading");return f(),D("div",re,[N("div",ie,[t($,{inline:!0},{default:a(()=>[t(h,{label:"存储库"},{default:a(()=>[J((f(),V(b,{style:{width:"10rem"},modelValue:e.snapshotNameList,"onUpdate:modelValue":l[0]||(l[0]=n=>e.snapshotNameList=n),clearable:"",filterable:"",multiple:"",placeholder:e.$t("请选择存储库"),onChange:e.search},{default:a(()=>[(f(!0),D(M,null,P(e.resData,(n,C,z)=>(f(),V(c,{key:z,label:C,value:C},null,8,["label","value"]))),128))]),_:1},8,["modelValue","placeholder","onChange"])),[[L,e.loading]])]),_:1}),t(h,null,{default:a(()=>[t(g,{type:"warning",onClick:l[1]||(l[1]=n=>e.openAddDialog=!0)},{default:a(()=>[y(_(e.$t("新建存储库")),1)]),_:1})]),_:1}),t(h,null,{default:a(()=>[t(g,{type:"success",onClick:e.search},{default:a(()=>[y(_(e.$t("刷新")),1)]),_:1},8,["onClick"])]),_:1})]),_:1})]),t(u,{shadow:"never",class:"table-container"},{default:a(()=>[t(S,{data:e.tableData},{default:a(()=>[t(d,{label:e.$t("序号"),align:"center",fixed:"",width:"100"},{default:a(n=>[y(_(n.$index+1),1)]),_:1},8,["label"]),t(d,{align:"center",label:e.$t("存储库名"),prop:"name",width:"200"},null,8,["label"]),t(d,{align:"center",label:e.$t("存储库地址"),prop:"location",width:"300"},null,8,["label"]),t(d,{align:"center",label:e.$t("类型"),prop:"type",width:"200"},null,8,["label"]),t(d,{align:"center",label:e.$t("是否压缩"),prop:"compress",width:"100"},null,8,["label"]),t(d,{align:"center",label:e.$t("分解块大小"),prop:"chunk_size",width:"100"},null,8,["label"]),t(d,{align:"center",label:e.$t("是否只读(默认false)"),prop:"readonly",width:"100"},null,8,["label"]),t(d,{align:"center",label:e.$t("制作快照的速度"),width:"100"},{default:a(n=>[n.row.max_snapshot_bytes_per_sec!=""?(f(),D("div",ue,_(n.row.max_snapshot_bytes_per_sec)+"/s",1)):(f(),D("div",pe,"20mb/s"))]),_:1},8,["label"]),t(d,{align:"center",label:e.$t("快照恢复的速度"),width:"100"},{default:a(n=>[n.row.max_restore_bytes_per_sec!=""?(f(),D("div",de,_(n.row.max_restore_bytes_per_sec)+"/s",1)):(f(),D("div",me,"20mb/s"))]),_:1},8,["label"]),t(d,{align:"center",label:e.$t("操作"),fixed:"right",width:"350"},{default:a(n=>[t(E,null,{default:a(()=>[t(g,{type:"success",onClick:C=>e.look(n.row.name)},{default:a(()=>[y(_(e.$t("查看")),1)]),_:2},1032,["onClick"]),t(g,{type:"warning",onClick:C=>e.updateSnapshotData(n.row)},{default:a(()=>[y(_(e.$t("修改")),1)]),_:2},1032,["onClick"]),t(g,{type:"danger",onClick:C=>e.SnapshotDeleteRepositoryAction(n.row.name)},{default:a(()=>[y(_(e.$t("删除")),1)]),_:2},1032,["onClick"]),t(g,{type:"warning",onClick:C=>e.CleanupeRepository(n.row.name)},{default:a(()=>[y(_(e.$t("清理")),1)]),_:2},1032,["onClick"])]),_:2},1024)]),_:1},8,["label"])]),_:1},8,["data"])]),_:1}),t(R,{ref:"drawer",title:e.$t("JSON数据"),"before-close":e.drawerHandleClose,modelValue:e.drawerShow,"onUpdate:modelValue":l[3]||(l[3]=n=>e.drawerShow=n),direction:"rtl","close-on-press-escape":"","destroy-on-close":"",size:"50%"},{default:a(()=>[e.drawerShow?(f(),V(j,{key:0,value:e.jsonData,"onUpdate:value":l[2]||(l[2]=n=>e.jsonData=n),styles:"width: 100%",read:!0,title:e.$t("JSON数据")},null,8,["value","title"])):B("",!0)]),_:1},8,["title","before-close","modelValue"]),e.openAddDialog?(f(),V(O,{key:0,"snapshot-data":e.snapshotData,open:e.openAddDialog,onClose:e.closeAddDialog},null,8,["snapshot-data","open","onClose"])):B("",!0)])}const we=U(ne,[["render",ce],["__scopeId","data-v-4861b8cc"]]);export{we as default};