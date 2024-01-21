<script setup>
import router from '@/router'
import {useAuthStore} from "@/stores/store";

const items = [
  { text: 'Home', icon: 'mdi-home', link: '/home' },
  { text: 'School', icon: 'mdi-school', link: '/school' },
  { text: 'Module', icon: 'mdi-view-module', link: '/module' },
  { text: 'Class', icon: 'mdi-google-classroom', link: '/class' },
  { text: 'Result', icon: 'mdi-format-list-numbered', link: '/result' },
]

function navigateTo(route) {
  router.push(route)
}

async function logout() {
  const authStore = useAuthStore();
  await authStore.logout();
  navigateTo('/')
}
</script>


<template>
  <v-card
      class="mx-auto"
      width="256"
  >
    <v-layout>
      <v-navigation-drawer permanent absolute>
        <v-list class="d-flex flex-column pa-1">
          <div class="d-flex justify-space-between">
            <v-list-item
                title="Bram Terlouw"
                subtitle="admin@admin.com"
                height="50"
            >
            </v-list-item>
            <v-btn icon="mdi-logout" size="small" class="mt-1 mr-3" @click="logout"></v-btn>
          </div>
        </v-list>
        <v-divider></v-divider>

        <v-list
            :lines="false"
            density="default"
            nav
        >
          <v-list-item
              v-for="(item, i) in items"
              :key="i"
              :value="item"
              color="primary"
              @click="navigateTo(item.link)"
          >
            <template v-slot:prepend>
              <v-icon :icon="item.icon"></v-icon>
            </template>
            <v-list-item-title v-text="item.text"></v-list-item-title>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main style="height: 100vh;"></v-main>
    </v-layout>
  </v-card>
</template>

<style scoped>
</style>