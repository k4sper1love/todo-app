<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {Greet, SetUsername} from '../../wailsjs/go/main/App'

const username = ref('');
const emit = defineEmits(['usernameSet'])

const saveUsername = async () => {
  if (username.value.trim() !== '') {
    await SetUsername(username.value)
    emit('usernameSet')
  }
}

</script>

<template>
  <main>
    <img id="logo" alt="ToDo logo" src="../assets/images/logo.png"/>
    <div id="result" class="result">Давайте знакомиться! Введите свое имя 👇</div>
    <div id="input" class="input-box">
      <input id="name" v-model="username" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="saveUsername">Сохранить</button>
    </div>
  </main>
</template>

<style scoped>
#logo {
  display: block;
  max-width: 50vh;
  max-height: 50vh;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}

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
