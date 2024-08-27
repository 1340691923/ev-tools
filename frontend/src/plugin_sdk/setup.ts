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
import {createI18n, useI18n} from "vue-i18n";

const messages = {
    "zh-cn": {
        ...zhCnLocale,
    },
    en: {
        ...enLocale,
    },
};


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

    const i18n = createI18n({
        legacy: false,
        locale: sdk.getLanguage(),
        messages: messages,
        globalInjection: true,
    });

    app.use(i18n)

    app = registerPlugin(app)

    app
        .use(router)
        .mount(container ? container.querySelector('#app') : '#app')

}

export const setupEvPlugin = (registerPlugin) => {
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
