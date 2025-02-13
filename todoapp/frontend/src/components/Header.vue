<script setup lang="ts">
import {onMounted, ref} from "vue";
import {Greet} from "../../wailsjs/go/main/App";

const message = ref('');
const props = defineProps<{
  username: string,
  currentTime: string,
}>();

const getMessage = async () => {
  message.value = await Greet(props.username);
};

onMounted(() => {
  getMessage();
});
</script>

<template>
  <div class="d-flex flex-column justify-content-center align-items-center">
    <p class="fw-lighter fs-2 mt-4">{{ message }}</p>
    <p class="fw-lighter fs-5">🕓 {{ currentTime }}</p>
  </div>
</template>

<style scoped>

</style>