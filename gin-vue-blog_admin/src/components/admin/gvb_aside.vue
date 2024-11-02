<script setup>
import {ref} from "vue"
const selectedKeys=ref(["1"])
</script>

<template>
    <aside class="gvb_aside">
        <div class="gvb_aside_header flex">
            <div class="gvb_aside_logo">
                <img src="https://www.logoko.com.cn/uploadfile/icon_case/201808/5b86063b48933.png">
            </div>
            <div class="gvb_aside_title">博客后台系统</div>
        </div>
        <div class="gvb_aside_body">
          <a-menu
          v-model:selectedKeys="selectedKeys"
          mode="inline"
          theme="dark"
          :inline-collapsed="false"
        >
          <a-menu-item key="1">
            <template #icon>
              <i class="fa fa-home"></i>
            </template>
            <span>首页</span>
          </a-menu-item>
          <a-sub-menu key="sub1">
            <template #icon>
              <i class="fa fa-user-circle-o"></i>
            </template>
            <template #title>用户管理</template>
            <a-menu-item key="5">用户列表</a-menu-item>
            <a-menu-item key="6">添加用户</a-menu-item>
          </a-sub-menu>
        </a-menu>
        </div>
    </aside>
</template>

<style lang="scss">
.gvb_aside {
  width: 240px;
  height: 100vh;
  background-color: var(--slide);

  .gvb_aside_header{
    width: 100%;
    height: 260px;
    flex-direction: column;
  }

  .gvb_aside_logo{
    padding: 10px;
    img{
        width: 90%;
    }
  }

  .gvb_aside_title{
    margin-top: 20px;
    font-size: 20px;
  }
}
</style>