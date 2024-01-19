<script setup>
import { router } from '../router';
</script>

<template>
  <div class="mb-16 mt-5 gb-5">
    <div class="text-h3">Password manager</div>
    <div class="text-h6">Welcome to the password manager</div>
  </div>
  <div class="d-flex flex-column justify-center mb-6 gy-3">

    <v-btn class="mx-auto bg-teal-lighten-1 font-weight-bold ma-6" min-width="300" size="x-large" rounded="pill"
      @click="router.push('/AddNew')">Add</v-btn>
    <v-btn class="mx-auto bg-teal-darken-1 font-weight-bold ma-6" min-width="300" size="x-large" rounded="pill"
      @click="router.push('/AddOwn')">Add with
      password</v-btn>
    <v-btn class="mx-auto bg-teal-darken-2 font-weight-bold ma-6" min-width="300" size="x-large" rounded="pill"
      @click="router.push('/Get')">Get</v-btn>
    <v-btn class="mx-auto bg-teal-darken-3 font-weight-bold ma-6" min-width="300" size="x-large" rounded="pill"
      @click="router.push('/GetAll')">Get
      all</v-btn>
    <v-btn class="mx-auto bg-teal-darken-4 font-weight-bold ma-6" min-width="300" size="x-large" rounded="pill"
      @click="router.push('/Delete')">Delete</v-btn>
  </div>
</template>
