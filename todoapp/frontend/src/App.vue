<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import CreateTask from "./components/CreateTask.vue";</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo.png"/>

  <HelloWorld/>
  <CreateTask/>
</template>

<style>
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
