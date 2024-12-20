import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:"/login",
      name:"login",
      component:()=>import("../views/login.vue")
    },
    {
      path:"/admin",
      name:"admin",
      component:()=>import("../views/admin/admin.vue"),
      children:[
        {
          path:"home",
          name:"home",
          component:()=>import("../views/admin/home/home.vue"),
        },
        {
          path:"system_list",
          name:"system_list",
          component:()=>import("../views/admin/system_mgr/system_list.vue"),
        },
        {
          path:"user_list",
          name:"user_list",
          component:()=>import("../views/admin/user_mgr/user_list.vue"),
        }
      ]
    },
    {
      path:"/",
      name:"",
      component:()=>import("../views/admin/admin.vue"),
    }
  ]
})

export default router
