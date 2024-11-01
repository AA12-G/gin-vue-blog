<script setup>
import {ref} from "vue";
const isFull=ref (true)
//页面全屏事件
function fullScreen(){
    let element=document.documentElement
    //判断是否处于全屏状态
    isFull.value=!isFull.value
    if(!isFullScreen()){
        //退出全屏
        exitFullScreen()
        return
    }
    let requestMethod = element.requestFullScreen || element.webkitRequestFullScreen || element.mozRequestFullScreen || element.msRequestFullScreen;
    if (requestMethod) {
        requestMethod.call(element);
    } else if (typeof window.ActiveXObject !== "undefined") {//for Internet Explorer
        let wscript = new ActiveXObject("WScript.Shell");
        if (wscript !== null) {
            wscript.SendKeys("{F11}");
        }
    }
}
//取消全屏事件
function exitFullScreen() {
    let exitMethod = document.cancelFullScreen || document.webkitCancelFullScreen || document.mozCancelFullScreen || document.exitFullScreen;
    if (exitMethod) {
        exitMethod.call(document);
    } else if (typeof window.ActiveXObject !== "undefined") {
        let wscript = new ActiveXObject("WScript.Shell");
        if (wscript != null) {
            wscript.SendKeys("{F11}");
        }
    }
}
//判断页面是否处于全屏
function isFullScreen() {
    return document.fullscreenElement == null
}
</script>

<template>
    <i v-if="isFull" @click="fullScreen" class="iconfont icon-quanping"></i>
    <i v-else @click="fullScreen" class="iconfont icon-quxiaoquanping"></i>
</template>