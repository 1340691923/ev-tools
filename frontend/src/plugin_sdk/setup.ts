import pluginJson from '../../../plugin.json'
import {renderWithQiankun} from 'vite-plugin-qiankun/dist/helper'
import {QiankunProps} from "vite-plugin-qiankun/es/helper";
import router from "../router"
import {EvAppConigStore, EvSettingsStore, handleEvAppConfig, handleEvSettingsStore} from "./ev-store";

import eventBus from 'vue3-eventbus'
import {sdk} from './sdk';
import {createApp} from "vue";
import App from "../App.vue";

// 本地语言包
import enLocale from "../lang/en";
import zhCnLocale from "../lang/zh-cn";
import {createI18n} from "vue-i18n";

//elementui

// 引入依赖
import ElementPlus from 'element-plus'
// 引入全局 CSS 样式
import 'element-plus/dist/index.css'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'


let app

// 引入依赖结束
const render = (props:QiankunProps,registerPlugin) => {

    const { container } = props

    app = createApp(App)

    app.use(eventBus)

    handleEvSettingsStore(props.store.useSettingsStore)

    props.store.onChangeSettingsStore?.((parentStore: any) => {
        handleEvSettingsStore(parentStore)
    })

    handleEvAppConfig(props.store.useAppStore)

    props.store.onChangeAppStore?.((parentStore: any) => {
        handleEvAppConfig(parentStore)
    })

    app.provide('useSettingsStoreHook',props.store.useSettingsStore)
    app.provide('useSettingsStore',EvSettingsStore)
    app.provide('useEvAppConigStoreHook',props.store.useAppStore)
    app.provide('useEvAppConigStore',EvAppConigStore)

    sdk.pluginAlias = pluginJson.plugin_alias
    sdk.callPluginCallBack = props.CallPluginApi
    sdk.selectEsConnId = props.GetSelectEsConnID()
    sdk.linkOptCallBack = props.LinkOptAction
    sdk.evRouter = props.store.router

    sdk.subToChannel = props.SubToChannel
    sdk.callToChannel = props.CallToChannel
    sdk.unSubscribeToChannel = props.UnSubscribeToChannel

    let i18nMessage = props.GetI18nMessage()

    let zhCnLocaleCfg = zhCnLocale
    let enLocaleCfg = enLocale

    if(i18nMessage.hasOwnProperty('zh-cn')){
        for(let k in i18nMessage['zh-cn']){
            zhCnLocaleCfg[k] = i18nMessage['zh-cn'][k]
        }
    }

    if(i18nMessage.hasOwnProperty('en')){
        for(let k in i18nMessage['en']){
            enLocaleCfg[k] = i18nMessage['en'][k]
        }
    }

    const messages = {
        "zh-cn": {
            ...zhCnLocaleCfg,
        },
        'en': {
            ...enLocaleCfg,
        },
    };

    const i18n = createI18n({
        legacy: false,
        locale: sdk.getLanguage(),
        messages: messages,
        globalInjection: true,
    });

    app.use(i18n)

    for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
        app.component(key, component)
    }

    app.use(ElementPlus)

    registerPlugin(app)

    app
        .use(router)
        .mount(container ? container.querySelector('#app') : '#app')

}

export const setupEvPlugin = (registerPlugin:Function) => {
    renderWithQiankun({
        update(props: QiankunProps): void | Promise<void> {
            return
        },
        mount(props) {
            render(props,registerPlugin)
        },
        bootstrap() {},
        unmount() {
            app.unmount()
        }
    })
}
