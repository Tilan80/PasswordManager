<script setup>
import {reactive} from 'vue'
import {CheckKey} from '../../wailsjs/go/main/App'


const data = reactive({
  name: "",
  resultText: "Please enter the password:",
})

function checkKey() {
  console.log("Got to here")
  CheckKey(data.name).then(result => {
    console.log(result)
    data.resultText = result ? "You are logged in!" : "You are not logged in!"
  })
}


</script>

<template>
  <main>
    <img id="logo" alt="Wails logo" src="../assets/images/logo-universal.png" />
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text" />
      <button class="btn" @click="checkKey">Login</button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
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
