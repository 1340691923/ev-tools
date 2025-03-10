<template>
  <div  ref="maxDiv" style="width: 100%;" >
    <div class="tools-box">
      <div>
        <span v-if="title != ''" >{{ title }}</span>
      </div>
      <div class="tools-list">

        <span class="aides" style="margin-right: 1rem" @click="format">
          <i class="el-icon-edit"></i>
          美化</span>
        <span class="aides" @click="copy">
          <i class="el-icon-delete"></i>
          复制</span>
      </div>
    </div>
    <div id="container" class="json-monaco-editor" ref="container" :style="{'height':`${height}px`,width:containerWidth}"></div>
  </div>
</template>
<script setup lang="ts">

import {monaco} from '../../utils/monaco';


import {computed, onMounted, ref, watch} from "vue";
import handleClipboard from "../../utils/clipboard";

import {sdk} from "@elasticview/plugin-sdk";


const props = defineProps({

  read: {
    type: Boolean,
    default: ()=>{
      return false
    }
  },
  title: {
    type: String,
    default: ''
  },
  value: {
    type: String,
    default: ()=>{
      return '{}'
    }
  },
  pointOut: {
    type: Array,
    default: ()=>{
      return []
    }
  },
  height: {
    type: String,
    default: ()=>{
      return '400'
    }
  },
  fontSize: {
    type: String,
    default: ()=>{
      return '18'
    }
  },
  language:{
    type:String,
    default:()=>{
      return 'json'
    }
  },
})

const copy = ()=>{
  handleClipboard(editor.getValue())
}

const getTheme = computed(()=>{
  return sdk.isDarkTheme() ?"vs-dark":"vs"
})

watch(() => sdk.isDarkTheme(), (newVal) => {

  if(newVal){
    document.documentElement.style.setProperty(
        "--line-number-bg-color",
        "#080808"
    );
  }else{
    document.documentElement.style.setProperty(
        "--line-number-bg-color",
        "#ebebeb"
    );
  }
})

const codesCopy = ref(null)

const SetText = (msg)=>{
  const currentPosition = editor.getPosition();

  editor.setValue(msg)

  if (currentPosition) {
    editor.setPosition(currentPosition);
    editor.focus();
  }
}

const emits = defineEmits(['update:value','getValue'])

onMounted(()=>{

  if(sdk.isDarkTheme()){
    document.documentElement.style.setProperty(
        "--line-number-bg-color",
        "#080808"
    );
  }else{
    document.documentElement.style.setProperty(
        "--line-number-bg-color",
        "#ebebeb"
    );
  }

  initEditor()

  sdk.getEventBus().on(sdk.getBusEnum().changeEvSettings,(settings)=>{

    if(editor){
      monaco.editor.setTheme(settings.theme == 'dark'?"vs-dark":"vs")
    }

  })

})

const container = ref()

let editor = null;

const format = () => {
  try {
    const value = editor.getValue()
    const tmp = JSON.parse(value)
    editor.setValue(JSON.stringify(tmp, null, '\t'))
  } catch (e) {
    console.log(e)
  }
}

const initEditor = () => {
  if(editor != null) return

  monaco.languages.registerCompletionItemProvider('json', {
    provideCompletionItems: function(model, position) {

      let arr = []

      for(let v of props.pointOut){

        let obj = {
          label: v.caption,
          insertText: v.value,
          kind: monaco.languages.CompletionItemKind.Keyword
        }

        if(v.hasOwnProperty('score')){
          obj["sortText"] = `${v.score}`
        }

        arr.push(
            obj
        )
      }

      return {
        suggestions: arr
      };
    }
  });

  container.value.innerHTML = '';
  editor = monaco.editor.create(container.value,{
    value:codesCopy.value || props.value,
    language: props.language,
    theme: getTheme.value,
    selectOnLineNumbers: true,
    roundedSelection: false,
    readOnly: props.read,		// 只读
    cursorStyle: 'line',		//光标样式
    automaticLayout: true,	//自动布局
    glyphMargin: true,  //字形边缘
    useTabStops: false,
    fontSize: props.fontSize,		//字体大小
    autoIndent:true,//自动布局
    folding: true,  // 启用折叠功能
  });


  editor.onDidChangeModelContent(function(event){//编辑器内容changge事件
    codesCopy.value = editor.getValue();
    emits('update:value',editor.getValue())
    emits('getValue',editor.getValue())
  });

  //updateEditorWidth()
  //编辑器随窗口自适应
  window.addEventListener('resize',()=>{
    updateEditorWidth()
  })

}

const containerWidth = ref('100%');
const maxDiv = ref()

const updateEditorWidth = () => {
  if (container.value) {
    containerWidth.value = `${maxDiv.value.clientWidth}px`;
  }
};

defineExpose({updateEditorWidth,SetText})

</script>
<style scoped>


.tools-box {
  display: flex;
  line-height: 1;
  font-size: 12px;
  padding: 8px 15px;
  align-items: center;
  justify-content: space-between;
  border-top: 1px solid #e7e8ee;
}

.aides:hover {
  color: #409eff;
  cursor: pointer;
}

.tools-list {
  display: flex;
}



</style>

<style>
.json-monaco-editor .margin {
  background-color: var(--line-number-bg-color);;
}



</style>
