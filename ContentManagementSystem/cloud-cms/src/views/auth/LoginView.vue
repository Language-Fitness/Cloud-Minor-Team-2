<script>
import {defineComponent} from 'vue'
import {useAuthStore} from "@/stores/store";
import router from "@/router";

export default defineComponent({
  name: "LoginView",
  data: () => ({
    valid: false,
    username: '',
    password: '',
    nameRules: [
      value => {
        if (value) return true
        return 'Name is required.'
      },
      value => {
        if (value?.length <= 10) return true
        return 'Name must be less than 10 characters.'
      },
    ],
  }),
  methods: {
    async login() {
      // const authStore = useAuthStore(); <- should go somewhere else ???
      // await authStore.login(this.username, this.password);
      await router.push('/')
    }
  }
})
</script>

<template>
  <v-form v-model="valid" class="w-100">
    <v-container class="mx-auto my-5 w-25 d-flex flex-column border-solid">
      <h2>Login</h2>

      <v-row>
        <v-col
            cols="12"
            md="12"
        >
          <v-text-field
              class="w-100"
              v-model="username"
              :rules="nameRules"
              label="User name"
              required
          ></v-text-field>
        </v-col>

        <v-col
            cols="12"
            md="12"
        >
          <v-text-field
              v-model="password"
              label="Password"
              required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-btn @click="login" type="submit" block class="mt-5">Submit</v-btn>
    </v-container>
  </v-form>
</template>

<style scoped>

</style>