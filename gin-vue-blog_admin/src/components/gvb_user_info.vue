<script setup>
import {useRouter} from "vue-router"
const router = useRouter()
function menuClick({key}){
    if(key === "logout"){
        console.log("logout") 
        return
    }
    router.push({
        name:key
    })
}
</script>

<template>
  <div class="gvb_user_info">
    <div class="avatar">
        <img src="https://avatars.githubusercontent.com/u/141143150?v=4&size=64" alt="">
    </div>
    <div class="drop_menu">
        <a-dropdown placement="bottomRight">
            <a class="ant-dropdown-link" @click.prevent>
                小广
              <i class="fa fa-angle-down"></i>
            </a>
            <template #overlay>
              <a-menu @click="menuClick">
                <a-menu-item key="login">
                  <a href="javascript:;" class="menu-item-text">个人中心</a>
                </a-menu-item>
                <a-menu-item key="my_message">
                  <a href="javascript:;" class="menu-item-text">我的消息</a>
                </a-menu-item>
                <a-menu-item key="article_list">
                  <a href="javascript:;" class="menu-item-text">消息列表</a>
                </a-menu-item>
                <a-menu-item key="logout">
                  <a href="javascript:;" class="menu-item-text">注销退出</a>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
    </div>
  </div>
</template>

<style>
.gvb_user_info{
    display: flex;
    align-items: center;
}
</style>