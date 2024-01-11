import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from "axios";

export const useAuthStore = defineStore('auth', () => {
  const isAuthenticated = ref(false)
  const username = ref('')
  const accessToken = ref(null)
  const refreshToken = ref(null)

  const login = async (username, password) => {
    const formData = new URLSearchParams();
    formData.append('grant_type', 'password');
    formData.append('client_id', 'login-client');
    formData.append('client_secret', '');
    formData.append('username', username);
    formData.append('password', password);
    try {
      const response = await axios.post('https://example-keycloak-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/realms/cloud-project/protocol/openid-connect/token', formData, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      });

      accessToken.value = response.data.access_token;
      refreshToken.value = response.data.refresh_token;
      console.log(accessToken.value)
      console.log(refreshToken.value)

      isAuthenticated.value = true;
    } catch (error) {
      console.error('Authentication failed', error);
      isAuthenticated.value = false;
    }
  }

  const logout = () => {
    isAuthenticated.value = false
    accessToken.value = null;
    refreshToken.value = null;
  }

  const isUserAdmin = computed(() => {
    // return username.value === 'admin';
  })

  return {
    isAuthenticated,
    username,
    login,
    logout,
    isUserAdmin
  }
})
