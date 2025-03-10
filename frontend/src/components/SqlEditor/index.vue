<template>
  <div  >
    <div class="tools-box">
      <div>
        <span v-if="title != ''" >{{ title }}</span>
      </div>
      <div class="tools-list">

        <span class="aides" style="margin-right: 1rem" @click="formatSql">
          <i class="el-icon-edit"></i>
          美化</span>
        <span class="aides" @click="copy">
          <i class="el-icon-delete"></i>
          复制</span>
      </div>
    </div>

    <div id="container" ref="container" :style="{'min-height':`${height}px`,width:containerWidth,'max-height':'400px'}"></div>
  </div>
</template>
<script setup lang="ts">
import * as monaco from 'monaco-editor';

import { format } from 'sql-formatter'

import editorWorker from "monaco-editor/esm/vs/editor/editor.worker?worker";

import {computed, inject, onMounted, ref} from "vue";
import handleClipboard from "../../utils/clipboard";

// 1. 引入monaco-editor中的sql文件
import { language as sqlLanguage } from 'monaco-editor/esm/vs/basic-languages/sql/sql.js';

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
      return 'sql'
    }
  },
})

const copy = ()=>{
  handleClipboard(editor.getValue())
}

self.MonacoEnvironment = {
  getWorker(workerId, label) {
    return new editorWorker();
  },
};

const codesCopy = ref(null)

const emits = defineEmits(['update:value'])

onMounted(()=>{
  initEditor()
  sdk.getEventBus().on("changeEvSettings",(settings)=>{
    monaco.editor.setTheme(settings.theme == 'dark'?"vs-dark":"vs")
  })
})

const container = ref()

let editor = null;

const formatSql = () => {
  try {
    const value = editor.getValue()
    editor.setValue(format(value))
  } catch (e) {
    console.log(e)
  }
}

const getTheme = computed(()=>{
  return sdk.isDarkTheme() ?"vs-dark":"vs"
})

const initEditor = () => {
  if(editor != null) return

  monaco.languages.registerCompletionItemProvider('sql', {
    provideCompletionItems: function(model, position) {

      let suggestions = []

      for(let v of props.pointOut){
        suggestions.push(
            { label: v.caption, kind: monaco.languages.CompletionItemKind.Keyword, insertText: v.value,sortText:  `${v.score}` }
        )
      }

      sqlLanguage.keywords.forEach(item => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Keyword,
          insertText: item
        });
      })
      sqlLanguage.operators.forEach(item => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Operator,
          insertText: item
        });
      })
      sqlLanguage.builtinFunctions.forEach(item => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Function,
          insertText: item
        });
      })
      sqlLanguage.builtinVariables.forEach(item => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Variable,
          insertText: item
        });
      })

      return {
        suggestions: suggestions
      };
    }
  });

  container.value.innerHTML = '';
  editor = monaco.editor.create(container.value,{
    value:codesCopy.value || props.value,
    language: props.language,
    theme: getTheme.value,
    editorOptions:{
      selectOnLineNumbers: true,
      roundedSelection: false,
      readOnly: props.read,		// 只读
      cursorStyle: 'line',		//光标样式
      automaticLayout: false,	//自动布局
      glyphMargin: true,  //字形边缘
      useTabStops: false,
      fontSize: props.fontSize,		//字体大小
      autoIndent:true,//自动布局
    },
  });
  editor.onDidChangeModelContent(function(event){//编辑器内容changge事件
    codesCopy.value = editor.getValue();
    emits('update:value',editor.getValue())
  });
  //编辑器随窗口自适应
  window.addEventListener('resize',()=>{
    editor.layout()
  })
}

const containerWidth = ref('100%');
const maxDiv = ref()

const updateEditorWidth = () => {
  if (container.value) {
    containerWidth.value = `${maxDiv.value.clientWidth}px`;
  }
};

const SetText = (msg)=>{
  const currentPosition = editor.getPosition();

  editor.setValue(msg)

  if (currentPosition) {
    editor.setPosition(currentPosition);
    editor.focus();
  }
}
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

