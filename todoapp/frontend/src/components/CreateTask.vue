<script setup lang="ts">
import {onMounted, ref} from 'vue';
  import {reactive} from 'vue'
  import {GetTasks, AddTask, RemoveTask} from '../../wailsjs/go/main/App'

  type Task = {
    id: number,
    text: string,
    status: "todo" | "in_progress" | "done";
  }

  const newTask = ref('')
  const tasks = ref<Task[]>([])

    const addTask = async () => {
      if (newTask.value.trim() !== '') {
        await AddTask(newTask.value);
        newTask.value = '';
        await loadTasks();
      }
    }

  const removeTask = async(id: number) => {
    await RemoveTask(id);
    await loadTasks();
  }

  const loadTasks = async () => {
    const data = await GetTasks();
    tasks.value = data as Task[];
  }

  onMounted(() => {
    loadTasks();
  });

</script>

<template>
  <div>
    <h1>Список задач</h1>
    <div>
      <input v-model="newTask" placeholder="Введите новую задачу" />
      <button @click="addTask">Добавить задачу</button>
    </div>
    <ul>
      <li v-for="(task, index) in tasks" :key="index">
        <span>{{ task.text }} {{ task.status }}</span>
        <button @click="removeTask(task.id)">Удалить</button>
      </li>
    </ul>
  </div>
</template>

<style scoped>

</style>