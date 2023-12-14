<script setup>
import { RouterView } from 'vue-router'
import SideNav from "@/components/nav/side-nav.vue";
</script>

<template>
  <header>
    <side-nav></side-nav>
  </header>

  <RouterView />
</template>

<style scoped>
 header {
   height: 100vh;
 }

 nav {
   display: flex;
   flex-direction: column;
   align-items: center;
   gap: 15px;
 }
@media (min-width: 1024px) {

}
</style>
