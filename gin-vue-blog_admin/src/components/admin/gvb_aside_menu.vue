<script setup>
import {reactive, ref} from "vue"
import {useRouter} from "vue-router"

const data = reactive({
    menuList:[
        {
            id:1,
            icon:"fa-home", //icon的图片 统一用 fa
            title:"首页", //菜单名称
            name:"home", //路由名称
            children:[]
        },
        {
            id:2,
            icon:"fa-user-circle-o", //icon的图片 统一用 fa
            title:"用户管理", //菜单名称
            name:"", //路由名称
            children:[
                {
                  id:3,
                  icon:"fa-user-circle", //icon的图片 统一用 fa
                  title:"用户列表", //菜单名称
                  name:"user_list", //路由名称
                },
            ]
        },
        {
            id:5,
            icon:"fa-cog", //icon的图片 统一用 fa
            title:"系统管理", //菜单名称
            name:"", //路由名称
            children:[
                {
                    id:6,
                    icon:"fa-cog", //icon的图片 统一用 fa
                    title:"系统配置", //菜单名称
                    name:"system_list", //路由名称
                }
            ]    
        }
    ]
})
const selectedKeys=ref(["1"])
const router = useRouter()

function goto(event){
    router.push({
        name:event.key
    })
}
</script>

<template>
    <a-menu
          v-model:selectedKeys="selectedKeys"
          mode="inline"
          theme="dark"
          :inline-collapsed="false"
          @click="goto"
        >
          <template v-for="menu in data.menuList" :key="menu.name">
          <a-menu-item :key="menu.name" v-if="menu.children.length === 0">
            <template #icon>
              <i :class="'fa ' +menu.icon "></i>
            </template>
            <span>{{ menu.title }}</span>
          </a-menu-item>
          <a-sub-menu :key="menu.id" v-else>
            <template #icon>
              <i :class="'fa ' +menu.icon "></i>
            </template>
            <template #title>{{ menu.title }}</template>
            <a-menu-item v-for="sub_menu in menu.children" :key="sub_menu.name">
                <template #icon>
                    <i :class="'fa ' +sub_menu.icon "></i>
                </template>
                <span>{{ sub_menu.title }}</span>
            </a-menu-item>
          </a-sub-menu>
        </template>
        </a-menu>
</template>

<style>

</style>