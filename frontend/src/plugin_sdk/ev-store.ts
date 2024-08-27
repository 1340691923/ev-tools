import {ref} from "vue";

import bus from 'vue3-eventbus'
import {BusEnum} from "./enum";

let EvSettingsStore:any=ref(null)

let EvAppConigStore:any=ref(null)

// 获取主应用的数据,并存储在一个变量中，使用provide方便全部页面拿到
export function handleEvSettingsStore(parentStore: any) {
    bus.emit(BusEnum.changeEvSettings, parentStore)
    EvSettingsStore.value= {...parentStore}
}

export function handleEvAppConfig(parentStore: any) {
    bus.emit(BusEnum.changeEvAppConfig, parentStore)
    EvAppConigStore.value= {...parentStore}
}

export {EvSettingsStore,EvAppConigStore}
