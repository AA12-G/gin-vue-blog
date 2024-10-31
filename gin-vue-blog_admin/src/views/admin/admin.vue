<script setup>
import GVBTheme from '../../components/gvb_admin.vue'
import GVBAside from '../../components/admin/gvb_aside.vue'
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
<div class="gvb_admin">
    <GVBAside></GVBAside>
    <div class="main">
        <header>
           <div class="left">
            <a-breadcrumb>
                <a-breadcrumb-item class="menu"><a href="">首页</a></a-breadcrumb-item>
                <a-breadcrumb-item class="menu"><a href="">个人中心</a></a-breadcrumb-item>
                <a-breadcrumb-item class="menu"><a href="">用户列表</a></a-breadcrumb-item>
              </a-breadcrumb>
           </div>
           <div class="right">
                <div class="icon_actions">
                    <i class="fa fa-home"></i>
                    <GVBTheme></GVBTheme>
                    <i class="fa fa-arrows-alt"></i>
                </div>
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
        </header>
        <div class="tabs"></div>
        <main></main>
    </div>
</div>
</template>

<style lang="scss">
.gvb_admin{
    width: 100%;
    display: flex;

    aside{
        width: 240px;
        height: 100vh;
        background-color: var(--slide);
    }

    .main{
        width: calc(100% - 240px);
        header{
            height: 60px;
            padding: 0 40px;
            display: flex;
            justify-content: space-between;
            align-items: center;
        

        .right{
            display: flex;
            align-items: center;
        }

        .icon_actions{
            margin-right: 20px;

            i{
                margin-left: 10px;
                cursor: pointer;
                font-size: 16px;
                color: var(--text);
            }
            i:hover{
                color: var(--active);
            }
        }

        .avatar{
            img{
                width: 40px;
                height: 40px;
                border-radius: 50%;
            }
        }
        .drop_menu{
            margin-left: 10px;
        
        }
    }
        .tabs{
            height: 30px;
            border: 1px solid var(--order);
        }
        main{
            background-color: var(--bg);
            height: calc(100vh - 90px);
        }
    }
}

.dark .menu-item-text{
    color: gray !important;
}
.dark .menu{
    color: gray !important;
}
</style>