<script>
import {defineComponent} from 'vue'

export default defineComponent({
  data() {
    return {
      tab: null,
    }
  },
})
</script>

<template>
  <v-card class="tabs">
    <v-tabs
        height="65"
        v-model="tab"
        bg-color="primary"
    >
      <v-tab value="overview">Overview</v-tab>
      <v-tab value="management">Management</v-tab>
      <v-tab value="other">Other</v-tab>
    </v-tabs>

    <v-card-text>
      <v-window v-model="tab">
        <v-window-item value="overview">
          Overview
        </v-window-item>

        <v-window-item value="management">
          Management
        </v-window-item>

        <v-window-item value="other">
          Other
        </v-window-item>
      </v-window>
    </v-card-text>
  </v-card>
</template>

<style scoped>
  .tabs {
    width: 85%;
    border-radius: 0;
  }
</style>