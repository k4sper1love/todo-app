<script setup lang="ts">
import {onMounted, ref} from "vue";
import {Profile} from "../types/profile";
import {createEmptyProfile} from "../utils/profile";
import {createProfile, fetchProfiles} from "../services/profile";

const emit = defineEmits<{
  (event: "profileSelected", id: number): void;
}>();

const newProfile = ref<Profile>(createEmptyProfile())
const profiles = ref<Profile[]>([]);

const addProfile = async() => {
  if(!newProfile.value.name.trim()) return

  const newId = await createProfile(newProfile.value)
  if (newId !== -1) {
    emit("profileSelected", newId)
  }

  newProfile.value = createEmptyProfile()
}

const loadProfiles = async() => {
  profiles.value = await fetchProfiles()
}

const selectProfile = async(id: number) => {
  if (id == -1) return
  emit("profileSelected", id)
}

onMounted(() => {
  loadProfiles();
})
</script>

<template>
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
  <div id="result" class="result">Создайте новый профиль 😊</div>

  <!-- Input field for the username -->
  <div id="input" class="input-group justify-content-center">
    <span class="input-group-text"><i class="fas fa-user">⌨️</i></span>
    <input id="name" v-model="newProfile.name" autocomplete="off" class="input form-control-sm" type="text" placeholder="Введите имя"/>
    <!-- Button to save the username -->
    <button class="btn  btn-outline-primary"  type="button"  @click="addProfile">Это я</button>
  </div>

  <ul>
    <li v-for="profile in profiles" :key="profile.id" @click="selectProfile(profile.id)">{{ profile.name }}</li>
  </ul>
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