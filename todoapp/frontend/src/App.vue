<script lang="ts" setup>
import CreateTask from "./components/CreateTask.vue";
import {onMounted, onUnmounted, ref} from "vue";
import {Task} from "./types/task";
import TaskList from "./components/TaskList.vue";
import {fetchActiveTasks, fetchCompletedTasks} from "./services/task";
import {fetchUsername} from "./services/user";
import Header from "./components/Header.vue";
import Greet from "./components/Greet.vue";

// State variables
const username = ref<string | null>(null); // Store the username
const showCompletedTasks = ref(true); // Toggle visibility of completed tasks
const currentTime = ref(''); // Store current time
const lastMinute = ref(new Date().getMinutes()); // Track last minute for task refresh
const activeTasks = ref<Task[]>([]); // Store active tasks
const completedTasks = ref<Task[]>([]); // Store completed tasks
let timeInterval: ReturnType<typeof setInterval>; // Interval variable for updating time

// Function to load the username
const loadUsername = async () => {
  username.value = await fetchUsername();
};

// Function to load tasks
const loadTasks = async () => {
  activeTasks.value = await fetchActiveTasks();
  completedTasks.value = await fetchCompletedTasks();
};

// Function to update the current time and refresh tasks every minute
const updateTime = () => {
  const now = new Date();
  currentTime.value = now.toLocaleTimeString();

  // Refresh tasks if the minute changes
  if (now.getMinutes() != lastMinute.value) {
    lastMinute.value = now.getMinutes();
    loadTasks();
  }
};

onMounted(() => {
  loadUsername(); // Load username when component mounts
  loadTasks(); // Load tasks when component mounts
  updateTime(); // Set initial time
  timeInterval = setInterval(updateTime, 1000); // Update time every second
});

onUnmounted(() => {
  clearInterval(timeInterval); // Clear the interval when component unmounts
});
</script>

<template>
  <Greet v-if="!username" @usernameSet="loadUsername" /> <!-- Show Greet until username is set -->
  <template v-else>
    <Header :username="username" :currentTime="currentTime" /> <!-- Show header with username and current time -->
    <CreateTask @loadTasks="loadTasks" /> <!-- Show task creation component -->

    <!-- Show active tasks if available -->
    <div v-if="activeTasks.length > 0">
      <TaskList :tasks="activeTasks" @loadTasks="loadTasks" /> <!-- Show active tasks list -->
    </div>
    <p class="fw-lighter fs-2 mt-2 ms-2" v-else>Активных задач нет 😌</p> <!-- Message when no active tasks -->

    <!-- Show completed tasks if available -->
    <div v-if="completedTasks.length > 0">
      <h2 class="fw-lighter fs-2 mt-4 ms-2" @click="showCompletedTasks = !showCompletedTasks" style="cursor: pointer;">
        Завершенные задачи {{ showCompletedTasks ? '▼' : '▶' }}
      </h2>
      <TaskList v-if="showCompletedTasks" :tasks="completedTasks" @loadTasks="loadTasks" /> <!-- Show completed tasks list -->
    </div>
  </template>
</template>

<style>
body .star-icon {
  cursor: pointer;
  font-size: 1.5rem;
  color: var(--primary-color) !important;
  transition: color 0.3s ease;
}

body .star-icon:hover {
  color: var(--secondary-color);
}
</style>
