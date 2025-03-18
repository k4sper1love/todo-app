<script setup lang="ts">
import {onMounted, ref} from "vue";
import {Greet} from "../../wailsjs/go/main/App";

// Define props for username and current time
const props = defineProps<{
  name: string,
  currentTime: string,
}>();

const emit = defineEmits(["logout"])

// Reactive variable to store the greeting message
const message = ref('');

// Fetch greeting message based on username
const getMessage = async () => {
  message.value = await Greet(props.name);
};

// Fetch the message when component is mounted
onMounted(() => {
  getMessage();
});
</script>

<template>
  <div class="d-flex flex-column justify-content-center align-items-center">
    <!-- Display the greeting message -->
    <p class="fw-lighter fs-2 mt-4">{{ message }}</p>
    <!-- Display the current time -->
    <p class="fw-lighter fs-5">🕓 {{ currentTime }}</p>
  </div>
  <div class="accordion-button" @click="emit('logout')">
    exit
  </div>
</template>

<style scoped>

</style>