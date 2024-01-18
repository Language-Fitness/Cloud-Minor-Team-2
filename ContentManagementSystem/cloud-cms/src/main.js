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
    // await authStore.refresh()
    console.log('Refresh has en error, has to do with the env files i think.')
};
const startRefreshTokenInterval = () => {
    refreshTokenInterval = setInterval(refreshTokenFunc, 30000);
};
const stopRefreshTokenInterval = () => {
    clearInterval(refreshTokenInterval);
};

app.use(router)
app.use(vuetify)
app.mount('#app')
