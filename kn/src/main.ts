import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import axios from 'axios';

import * as ElementPlusIconsVue from '@element-plus/icons-vue'


const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
// 注册 Vue Router
app.use(router);
axios.defaults.baseURL = "http://localhost:8090";
app.config.globalProperties.$axios = axios;


// 注册 Element Plus
app.use(ElementPlus);

app.mount('#app');