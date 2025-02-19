import {EvAppConigStore, EvSettingsStore} from "./ev-store";
import {Router} from "vue-router";
import bus from 'vue3-eventbus'
import {SizeEnum} from "./enum";

interface PluginApiInterface {
    CallPluginApi(req:Req): Promise<any>;
    GetSelectEsConnID() :number;
    LinkOptAction() :Promise<any>;
    IsMobile():boolean;
    isDarkTheme():boolean
    getRouter():Router
    getLayoutSize():SizeEnum
    getLanguage():string
    getEventBus():any
}

// 实现接口的方法
class SDK implements PluginApiInterface {

    pluginAlias:string;

    callPluginCallBack;

    linkOptCallBack;

    selectEsConnId:number;

    evRouter:Router;

    subToChannel;

    callToChannel;

    //channel 频道
    //msgCb 消息回调
    async SubToChannel(channel:string,msgCb:any) {
        this.subToChannel(channel,msgCb)
    }

    //channel 频道
    //msgCb 消息
    async CallToChannel(channel:string,msg:any) {
        this.callToChannel(channel,msg)
    }

    async CallPluginApi(req:Req): Promise<any> {
       return this.callPluginCallBack(this.pluginAlias,req.url,req.method,req.data)
    }

    async LinkOptAction(){
        return this.linkOptCallBack
    }

    GetSelectEsConnID():number{
        return this.selectEsConnId
    }

    IsMobile():boolean{
        return EvAppConigStore.value.device == 'mobile'
    }

    getRouter():Router{
        return this.evRouter
    }

    isDarkTheme():boolean{
        return EvSettingsStore.value.theme == 'dark'
    }

    getLayoutSize():SizeEnum{
        return EvAppConigStore.value.size
    }

    getLanguage():string{
        return EvAppConigStore.value.language
    }

    getEventBus(){
        return bus
    }

}

let sdk = new SDK()

declare interface Req{
    url:string;
    method:string;
    data:any;
}

export function request(req:Req){
    return sdk.CallPluginApi(req)
}

export {sdk}
