<script setup>
import { reactive } from 'vue';
import { router } from '../router';
import { GetByPlat } from '../../wailsjs/go/main/App';

const data = reactive({
  text: "",
  result: null,
  id: "",
  plat: "",
  user: "",
  pass: "",

});

function getByPlat() {
  GetByPlat(data.text).then(result => {
    data.id = result[0]
    data.plat = result[1]
    data.user = result[2]
    data.pass = result[3]
  })
}
</script>

<template>
  <h1>Get the saved password from the platform</h1>
  <v-btn class="mx-auto bg-green-accent-1 font-weight-bold ma-6" size="x-large" rounded="xl"
    @click="router.push('/Home')">Go back</v-btn>

  <v-text-field v-model="data.text" label="Input the platform you want to check" outlined></v-text-field>
  <v-btn class="mx-auto bg-brown-lighten-1 accent-1 font-weight-bold ma-6" size="x-large" rounded="xl" @click="getByPlat">Get By
    Platform</v-btn>

  <v-text-field v-if="data.id !== ''" v-model="data.id" label="ID" outlined readonly></v-text-field>
  <v-text-field v-if="data.plat !== ''" v-model="data.plat" label="Platform" outlined readonly></v-text-field>
  <v-text-field v-if="data.user !== ''" v-model="data.user" label="Username" outlined readonly></v-text-field>
  <v-text-field v-if="data.pass !== ''" v-model="data.pass" label="Password" outlined readonly></v-text-field>
</template>
