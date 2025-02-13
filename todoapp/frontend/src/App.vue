<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import CreateTask from "./components/CreateTask.vue";
import {onMounted, ref} from "vue";
import {Task} from "./types/task";
import TaskList from "./components/TaskList.vue";
import {fetchTasks} from "./repositories/task";
import {fetchUsername} from "./repositories/user";

const username = ref<string | null>(null)
const tasks = ref<Task[]>([]);

const loadUsername = async () => {
  username.value = await fetchUsername();
}

const loadTasks = async () => {
  tasks.value = await fetchTasks()
};

onMounted(() => {
  loadUsername()
  loadTasks();
});
</script>


<template>
  <HelloWorld v-if="!username" @usernameSet="loadUsername"/>
  <template v-else>
    <CreateTask :username="username" @loadTasks="loadTasks"/>
    <div>
      <h2>Ожидают выполнения</h2>
      <TaskList :tasks="tasks.filter(task => task.status === 'todo')" @loadTasks="loadTasks"/>
    </div>
    <div>
      <h2>Выполненные задания</h2>
      <TaskList :tasks="tasks.filter(task => task.status === 'done')" @loadTasks="loadTasks"/>
    </div>
  </template>

</template>

<style>
.star-icon {
  cursor: pointer;
  font-size: 20px;
  margin-left: 10px;
  color: gold;
}
</style>
