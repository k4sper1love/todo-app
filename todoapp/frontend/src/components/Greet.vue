<script lang="ts" setup>
import {ref} from 'vue';
import {setUsername} from "../services/user";

const emit = defineEmits(['usernameSet']);  // Emit event when username is set
const newUsername = ref(''); // Reactive variable for storing username input

// Function to save the username
const saveUsername = async () => {
  // Check if the username is not empty and trim whitespace
  if (newUsername.value.trim() !== '') {
    await setUsername(newUsername.value); // Save username
    emit('usernameSet'); // Emit event to notify parent component
  }
};
</script>

<template>
  <main>
    <!-- Display the logo -->
    <div>
      <img id="logo" alt="Todo logo" src="../assets/images/logo.png"/>
      <p>
        Привет! Я
        <a href="https://github.com/k4sper1love/todo-app" target="_blank" rel="noopener noreferrer"><span>TodoApp</span></a>
        by @k4sper1love
      </p>
    </div>

    <!-- Display prompt asking for the user's name -->
    <div id="result" class="result">А как вас зовут? 😊</div>

    <!-- Input field for the username -->
    <div id="input" class="input-group justify-content-center">
      <span class="input-group-text"><i class="fas fa-user">⌨️</i></span>
      <input id="name" v-model="newUsername" autocomplete="off" class="input form-control-sm" type="text" placeholder="Введите свое имя"/>
      <!-- Button to save the username -->
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
