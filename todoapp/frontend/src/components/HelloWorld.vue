<script lang="ts" setup>
import {ref} from 'vue'
import {setUsername} from "../repositories/user";

const username = ref('');
const emit = defineEmits(['usernameSet'])

const saveUsername = async () => {
  if (username.value.trim() !== '') {
    await setUsername(username.value)
    emit('usernameSet')
  }
}
</script>

<template>
  <main>
    <div>
      <img id="logo" alt="ToDo logo" src="../assets/images/logo.png"/>
      <p>
        Привет! Я
        <a href="https://github.com/k4sper1love/todo-app" target="_blank" rel="noopener noreferrer"><span>TodoApp</span></a>
        by @k4sper1love
      </p>
    </div>
    <div id="result" class="result">А как вас зовут? 😊</div>
    <div id="input" class="input-group justify-content-center">
      <span class="input-group-text"><i class="fas fa-user">⌨️</i></span>
      <input id="name" v-model="username" autocomplete="off" class="input form-control-sm" type="text" placeholder="Введите свое имя"/>
      <button class="btn  btn-outline-primary"  type="button"  @click="saveUsername">Это я</button>
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
</style>
