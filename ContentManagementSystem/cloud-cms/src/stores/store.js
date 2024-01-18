import {ref, computed} from 'vue'
import {defineStore} from 'pinia'
import axios from "axios";

export const useAuthStore = defineStore('auth', {
    state: () => ({
        isAuthenticated: false,
        username: '',
        accessToken: null,
        refreshToken: null
    }),

    actions: {
        async login(username, password) {
            const formData = new URLSearchParams();
            formData.append('grant_type', 'password');
            formData.append('client_id', 'login-client');
            formData.append('client_secret', 'xcfwgB7pZMyw9FshhZpcbyCwwmov10ux');
            formData.append('username', username);
            formData.append('password', password);

            try {
                const response = await axios
                    .post('https://example-keycloak-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/realms/cloud-project/protocol/openid-connect/token',
                        formData,
                        {
                            headers: {
                                'Content-Type': 'application/x-www-form-urlencoded',
                            }
                        });

                this.username = username
                this.accessToken = response.data.access_token;
                this.refreshToken = response.data.refresh_token;
                this.isAuthenticated = true;

            } catch (error) {
                console.error('Authentication failed', error);
                this.isAuthenticated.value = false;
                return false
            }
            this.isAuthenticated = true
            return true
        },

        async refresh() {
            const formData = new URLSearchParams();
            formData.append('grant_type', 'refresh_token');
            formData.append('client_id', import.meta.env.KEYCLOAK_CLIENT_ID);
            formData.append('client_secret', import.meta.env.KEYCLOAK_CLIENT_SECRET);
            formData.append('refresh_token', this.refreshToken)

            try {
                const response = await axios.post('https://example-keycloak-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/realms/cloud-project/protocol/openid-connect/token', formData, {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                    }
                });

                this.accessToken = response.data.access_token;
                this.refreshToken = response.data.refresh_token;

            } catch (error) {
                console.error('Refresh failed', error);
                this.isAuthenticated = false;
            }
        },

        logout() {
            this.isAuthenticated = false
            this.accessToken = null;
            this.refreshToken = null;
        }
    }
})
