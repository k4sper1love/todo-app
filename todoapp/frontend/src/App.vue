<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import CreateTask from "./components/CreateTask.vue";
import {GetUsername, SetUsername} from '../wailsjs/go/main/App';
import {onMounted, ref} from "vue";

const username = ref<string | null>(null)

const loadUsername = async () => {
  username.value = await GetUsername();
}

onMounted(() => {
  loadUsername()
})

</script>


<template>
  <HelloWorld v-if="!username" @usernameSet="loadUsername"/>
  <CreateTask v-else :username="username"/>
</template>

<style>

</style>
