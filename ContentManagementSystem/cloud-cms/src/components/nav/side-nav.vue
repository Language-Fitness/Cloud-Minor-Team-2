<script setup>
import router from '@/router'

const items = [
  { text: 'Home', icon: 'mdi-home', link: '/' },
  { text: 'School', icon: 'mdi-school', link: '/school' },
  { text: 'Module', icon: 'mdi-view-module', link: '/module' },
  { text: 'Class', icon: 'mdi-google-classroom', link: '/class' },
  { text: 'Result', icon: 'mdi-format-list-numbered', link: '/result' },
]

function navigateTo(route) {
  router.push(route)
}
</script>


<template>
  <v-card
      class="mx-auto"
      width="256"
  >
    <v-layout>
      <v-navigation-drawer permanent absolute>
        <v-list>
          <v-list-item
              prepend-avatar="https://cdn.vuetifyjs.com/images/john.png"
              title="Bram Terlouw"
              subtitle="bram@teacher.com"
              height="50"
          >
          </v-list-item>
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