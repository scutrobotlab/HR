import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/admin/evaluation-std/edit/:stdId?",
    name: "EditEvaStd",
    component: () => import("../views/Admin/EditEvaStd.vue"),
  },
  {
    path: "/admin/group/:group?/:state?/:page?",
    name: "Group",
    component: () => import("../views/Admin/Group.vue"),
  },
  {
    path: "/admin/person/:appId",
    component: () => import("../views/Admin/Applicant/Applicant.vue"),
    children: [
      {
        path: "info",
        component: () => import("../views/Admin/Applicant/ApplicantInfo.vue"),
      },
      {
        path: "evaluation",
        component: () => import("../views/Admin/Applicant/MyEvaluation.vue"),
      },
      {
        path: "stastics",
        component: () => import("../views/Admin/Applicant/RatingStastics.vue"),
      },
      {
        path: "exam",
        component: () => import("../views/Admin/Applicant/Exam.vue"),
      },
    ],
  },
  {
    path: "/admin/callback",
    name: "AdminCallback",
    component: () => import("../views/Admin/Callback.vue"),
  },
  {
    path: "/admin/search",
    name: "SearchView",
    component: () => import("../views/Admin/SearchView.vue"),
  },
  {
    path: "/admin/configs",
    component: () => import("../views/Admin/Configs/Frame.vue"),
    children: [
      {
        path: "main",
        component: () => import("../views/Admin/Configs/Main.vue"),
      },
      {
        path: "times",
        component: () => import("../views/Admin/Configs/Time/Time.vue"),
        meta: {
          title: "面试时间",
          icon: "mdi-account-clock",
          msg: "配置面试时间、面试地点，包括每个时间段的人数。",
          color: "accent",
        },
        children: [
          {
            path: "/",
            redirect: "export",
          },
          {
            path: "export",
            component: () => import("../views/Admin/Configs/Time/Export.vue"),
            meta: {
              title: "面试时间 - 导出数据",
            },
          },
          {
            path: "import",
            component: () => import("../views/Admin/Configs/Time/Import.vue"),
            meta: {
              title: "面试时间 - 导入数据",
            },
          },
        ],
      }, {
        path: "exam",
        component: () => import("../views/Admin/Configs/Exam/Exam.vue"),
        meta: {
          title: "题库",
          icon: "mdi-help-circle",
          msg: "设置面试的题库。",
          color: "red",
        },
      }, {
        path: "time-frame",
        component: () => import("../views/Admin/Configs/TimeFrame/TimeFrame.vue"),
        meta: {
          title: "时间节点",
          icon: "mdi-progress-clock",
          msg: "控制报名、选择面试时间、公布面试成绩等时间节点。",
          color: "green",
        },
      },
      {
        path: "/",
        redirect: "main",
      },
    ],
  },
  {
    path: "/admin/",
    redirect: "/admin/group",
  },
  {
    path: "/joinus",
    component: () => import("../views/Joinus/Joinus.vue"),
    children: [
      {
        path: "form",
        component: () => import("../views/Joinus/Form.vue"),
      },
      {
        path: "time",
        component: () => import("../views/Joinus/Time.vue"),
      },
      {
        path: "done",
        component: () => import("../views/Joinus/Done.vue"),
      },
      {
        path: "exam/main",
        component: () => import("../views/Joinus/Exam/Drawer.vue"),
        meta: {
          noTabs: true,
        }
      },
      {
        path: "exam/:group",
        component: () => import("../views/Joinus/Exam/Exam.vue"),
        meta: {
          noTabs: true,
        }
      },
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
