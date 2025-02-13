<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import CreateTask from "./components/CreateTask.vue";
import {onMounted, onUnmounted, ref} from "vue";
import {Task} from "./types/task";
import TaskList from "./components/TaskList.vue";
import {fetchTasks} from "./repositories/task";
import {fetchUsername} from "./repositories/user";
import Header from "./components/Header.vue";

const username = ref<string | null>(null)
const tasks = ref<Task[]>([]);
const showCompletedTasks = ref(true);
const currentTime = ref('');
const lastMinute = ref(new Date().getMinutes());


const loadUsername = async () => {
  username.value = await fetchUsername();
}

const loadTasks = async () => {
  tasks.value = await fetchTasks()
};

const updateTime = () => {
  const now = new Date();
  currentTime.value = now.toLocaleTimeString()

  if (now.getMinutes() != lastMinute.value) {
    lastMinute.value = now.getMinutes()
    loadTasks()
  }
}

let timeInterval: ReturnType<typeof setInterval>;

onMounted(() => {
  loadUsername()
  loadTasks();
  updateTime();
  timeInterval = setInterval(updateTime, 1000);
});

onUnmounted(() => {
  clearInterval(timeInterval);
})
</script>


<template>
  <HelloWorld v-if="!username" @usernameSet="loadUsername"/>
  <template v-else>
    <Header :username="username" :currentTime="currentTime"/>
    <CreateTask @loadTasks="loadTasks"/>
    <div v-if="tasks.filter(task => task.status === 'todo').length > 0">
      <TaskList :tasks="tasks.filter(task => task.status === 'todo')" @loadTasks="loadTasks"/>
    </div>
    <p class="fw-lighter fs-2 mt-2 ms-2" v-else>Активных задач нет 😌</p>
    <div v-if="tasks.filter(task => task.status === 'done').length > 0">
      <h2 class="fw-lighter fs-2 mt-4 ms-2" @click="showCompletedTasks = !showCompletedTasks" style="cursor: pointer;">
        Завершенные задачи {{ showCompletedTasks ? '▼' : '▶' }}
      </h2>
      <TaskList v-if="showCompletedTasks" :tasks="tasks.filter(task => task.status === 'done')" @loadTasks="loadTasks"/>
    </div>
  </template>

</template>

<style>
body .star-icon {
  cursor: pointer;
  font-size: 1.5rem;
  color: var(--primary-color);
  transition: color 0.3s ease;
}

body .star-icon:hover {
  color: var(--secondary-color);
}

</style>
