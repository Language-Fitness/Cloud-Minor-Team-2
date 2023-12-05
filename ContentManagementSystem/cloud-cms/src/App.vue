<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/school">School</RouterLink>
        <RouterLink to="/module">Module</RouterLink>
        <RouterLink to="/class">Class</RouterLink>
        <RouterLink to="/exercise">Exercise</RouterLink>
        <RouterLink to="/result">Result</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
 header {
   width: 15%;
   height: 100vh;
   background-color: #f8f8f8;
   border: 1px solid black;
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
