import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/home/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/school',
      name: 'school',
      component: () => import('../views/school/SchoolOverviewView.vue')
    },
    {
      path: '/module',
      name: 'module',
      component: () => import('../views/module/ModuleOverview.vue')
    },
    {
      path: '/class',
      name: 'class',
      component: () => import('../views/class/ClassOverview.vue')
    },
    {
      path: '/exercise',
      name: 'exercise',
      component: () => import('../views/exercise/ExerciseOverview.vue')
    },
    {
      path: '/result',
      name: 'result',
      component: () => import('../views/result/ResultOverview.vue')
    }
  ]
})

export default router
