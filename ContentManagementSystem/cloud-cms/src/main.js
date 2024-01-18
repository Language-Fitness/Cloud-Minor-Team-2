import './assets/main.css'

import {createApp, watch} from 'vue'
import {createPinia} from 'pinia'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
    components,
    directives,
})

import App from './App.vue'
import router from './router'
import {useAuthStore} from "@/stores/store";

const app = createApp(App)

app.use(createPinia())

const authStore = useAuthStore();
watch(() => authStore.isAuthenticated, (newIsAuthenticated) => {
    if (newIsAuthenticated) {
        startRefreshTokenInterval();
    } else {
        stopRefreshTokenInterval();
    }
});


let refreshTokenInterval;
const refreshTokenFunc = async () => {
    await authStore.refresh()
};
const startRefreshTokenInterval = () => {
    refreshTokenInterval = setInterval(refreshTokenFunc, 300000);
};
const stopRefreshTokenInterval = () => {
    clearInterval(refreshTokenInterval);
};

router.beforeEach((to, from) => {
    if (to.meta.requiresAuth) {
        if (!authStore.isAuthenticated && to.name !== 'login') {
            console.log('yes')
            router.push('/')
        }
    }
})

app.use(router)
app.use(vuetify)
app.mount('#app')
