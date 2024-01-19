<script setup>
import { reactive } from 'vue'
import { CheckKey } from '../../wailsjs/go/main/App'
import { router } from '../router';


const data = reactive({
  name: "",
  resultText: "Please enter the password:",
})


function checkKey() {
  CheckKey(data.name).then(result => {
    console.log(result)
    if (result) {
      data.resultText = "Correct, accessing page..."
      router.push("/Home")
    } else { data.resultText = "Wrong password" }
  })
}

</script>

<template>
  <main>
    <img id="logo" alt="Wails logo" src="../assets/images/logo-universal.png" />
    <div id="result" class="result">{{ data.resultText }}</div>
    <v-row align="center" justify="center">
      <v-col class="mx-auto" cols="12" md="4">
        <v-text-field label="Password" v-model="data.name" autocomplete="off" type="password"></v-text-field>
        <div class="d-flex align-center">
          <v-btn class="mx-auto mt-6 bg-green-accent-1 font-weight-bold" size="x-large" rounded="xl" 
          @click="checkKey">Login</v-btn>
        </div>
      </v-col>
    </v-row>

  </main>
</template>


<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  font-size: 24px;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
