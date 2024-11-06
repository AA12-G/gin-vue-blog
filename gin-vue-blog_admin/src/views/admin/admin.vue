<script setup>
import GVBTheme from '../../components/gvb_admin.vue'
import GVBAside from '../../components/admin/gvb_aside.vue'
import GVBFullScreen from '../../components/gvb_full_screen.vue'
import GVBUserInfo from '../../components/gvb_user_info.vue'

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
                    <GVBFullScreen></GVBFullScreen>
                </div>
                    <GVBUserInfo></GVBUserInfo>
           </div>    
        </header>
        <div class="tabs"></div>
        <main>
            <router-view></router-view>
        </main>
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