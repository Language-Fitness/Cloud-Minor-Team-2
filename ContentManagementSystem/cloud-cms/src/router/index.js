import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/home/HomeView.vue'
import LoginView from "@/views/auth/LoginView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView
    },
    {
      path: '/school',
      name: 'school',
      component: () => import('../views/school/SchoolOverviewView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/module',
      name: 'module',
      component: () => import('../views/module/ModuleOverview.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/class',
      name: 'class',
      component: () => import('../views/class/ClassOverview.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/exercise',
      name: 'exercise',
      component: () => import('../views/exercise/ExerciseOverview.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/result',
      name: 'result',
      component: () => import('../views/result/ResultOverview.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/home/HomeView.vue'),
      meta: {
        requiresAuth: true
      }
    }
  ]
})

export default router
