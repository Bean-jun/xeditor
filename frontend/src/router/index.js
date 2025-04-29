import { createRouter, createWebHashHistory } from "vue-router";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      component: () => import("../layouts/Layout.vue"),
      redirect: (to) => {
        return "/index";
      },
      children: [
        {
          path: "/index",
          meta: { name: "首页" },
          component: () => import("../pages/index.vue"),
        },
      ],
    },
  ],
});

export default router;
