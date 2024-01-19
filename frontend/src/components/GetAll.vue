<script setup>
import { reactive } from 'vue';
import { router } from '../router';
import { PrintAllRecords } from '../../wailsjs/go/main/App';

const data = reactive({
  text: "",
  result: null,
  id: "",
  plat: "",
  user: "",
  pass: "",

});

function printAllRecords() {
  PrintAllRecords().then(result => {
    data.result = result
  })
}
</script>

<template>
  <h1>Get all passwords</h1>
  <v-btn class="mx-auto bg-green-accent-1 font-weight-bold ma-6" size="x-large" rounded="xl"
    @click="router.push('/Home')">Go back</v-btn>

  <v-btn class="mx-auto bg-brown-lighten-1 accent-1 font-weight-bold ma-6" size="x-large" rounded="xl" @click="printAllRecords">Get all</v-btn>

  <!-- Use v-expansion-panels to loop through the outer layer of the array -->
  <v-expansion-panels v-if="data.result">
    <v-expansion-panel v-for="(record, index) in data.result" :key="index" class="bg-teal-lighten-2">
      <!-- Customize header appearance -->
      <v-expansion-panel-title class="bg-teal-lighten-2">
        {{ record[0] }}
      </v-expansion-panel-title>
      <!-- Customize content appearance -->
      <v-expansion-panel-text>
        <v-row v-if="record[0] !== ''" class="bg-teal-lighten-2">
          <v-col>
            <v-text-field v-model="record[1]" label="Username" outlined readonly></v-text-field>
          </v-col>
          <v-col>
            <v-text-field v-model="record[2]" label="Password" outlined readonly></v-text-field>
          </v-col>
        </v-row>
      </v-expansion-panel-text>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<style scoped>
.custom-header {
  margin-bottom: 10px; /* Add margin between expansion panels */
}

.custom-content {
  background-color: #f0f0f0; /* Set a different background color for the content */
  padding: 10px; /* Add padding for better appearance */
  margin-bottom: 10px; /* Specify the unit (e.g., px) for margin-bottom */
}

.custom-panel {
  background-color: #2196F3; /* Set your desired background color for the expansion panel */
}
</style>