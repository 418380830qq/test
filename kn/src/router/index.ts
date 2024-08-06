import { createRouter, createWebHistory,createWebHashHistory } from 'vue-router'
import index from '../components/index.vue'
import table from '../components/table.vue'


const routes = [
  {
    path: '/',
    name: 'index',
    component: index,
    children: [
      {
        path: '/table',
        name: 'table',
        component: table,
        
      }
    ]
  },
 
]

const router = createRouter({
  history: createWebHistory(),
  routes
})


export default router