<script lang="ts" setup>
import CreateTask from "./components/CreateTask.vue";
import {onMounted, onUnmounted, ref} from "vue";
import {Task} from "./types/task";
import TaskList from "./components/TaskList.vue";
import Profiles from "./components/Profiles.vue";
import {fetchActiveTasks, fetchCompletedTasks} from "./services/task";
import {fetchProfile, fetchProfiles} from "./services/profile";
import Header from "./components/Header.vue";
import {Profile} from "./types/profile";

// State variables
const profile = ref<Profile | null> (null)
const showCompletedTasks = ref(true); // Toggle visibility of completed tasks
const currentTime = ref(''); // Store current time
const lastMinute = ref(new Date().getMinutes()); // Track last minute for task refresh
const profiles = ref<Profile[]>([])
const activeTasks = ref<Task[]>([]); // Store active tasks
const completedTasks = ref<Task[]>([]); // Store completed tasks
let timeInterval: ReturnType<typeof setInterval>; // Interval variable for updating time

const loadProfiles = async() => {
  profiles.value = await fetchProfiles()
}

const loadProfile = async(id: number) => {
  profile.value = await fetchProfile(id)
  if (profile.value !== null) {
    await loadTasks(profile.value.id)
  }
}

// Function to load tasks
const loadTasks = async (id: number) => {
  activeTasks.value = await fetchActiveTasks(id);
  completedTasks.value = await fetchCompletedTasks(id);
};

// Function to update the current time and refresh tasks every minute
const updateTime = () => {
  const now = new Date();
  currentTime.value = now.toLocaleTimeString();

  // Refresh tasks if the minute changes
  if (now.getMinutes() != lastMinute.value) {
    lastMinute.value = now.getMinutes();
    if (profile.value !== null) {
      loadTasks(profile.value.id);
    }
  }
};

const changeProfile = async() => {
  profile.value = null
}

onMounted(() => {
  loadProfiles();
  updateTime(); // Set initial time
  timeInterval = setInterval(updateTime, 1000); // Update time every second
});

onUnmounted(() => {
  clearInterval(timeInterval); // Clear the interval when component unmounts
});
</script>

<template>
  <Profiles v-if="profile === null" @profileSelected="loadProfile" /> <!-- Show Greet until username is set -->
  <template v-else>
    <Header :name="profile.name" :currentTime="currentTime" @logout="changeProfile" /> <!-- Show header with username and current time -->
    <CreateTask :profileId="profile.id" @loadTasks="loadTasks(profile.id)" /> <!-- Show task creation component -->

    <!-- Show active tasks if available -->
    <div v-if="activeTasks.length > 0">
      <TaskList :tasks="activeTasks"  @loadTasks="loadTasks(profile.id)" /> <!-- Show active tasks list -->
    </div>
    <p class="fw-lighter fs-2 mt-2 ms-2" v-else>Активных задач нет 😌</p> <!-- Message when no active tasks -->

    <!-- Show completed tasks if available -->
    <div v-if="completedTasks.length > 0">
      <h2 class="fw-lighter fs-2 mt-4 ms-2" @click="showCompletedTasks = !showCompletedTasks" style="cursor: pointer;">
        Завершенные задачи {{ showCompletedTasks ? '▼' : '▶' }}
      </h2>
      <TaskList v-if="showCompletedTasks" :tasks="completedTasks"  @loadTasks="loadTasks(profile.id)" /> <!-- Show completed tasks list -->
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
