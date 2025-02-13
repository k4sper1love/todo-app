<script setup lang="ts">
import { onMounted, ref } from 'vue';
import {GetTasks, AddTask, DeleteTask, Greet, UpdateTaskStatus, UpdateTaskPriority} from '../../wailsjs/go/main/App';
import {createEmptyTask, Task} from "../types/task";
import {createTask} from "../repositories/task";

const props = defineProps<{ username: string }>();
const emit = defineEmits(["loadTasks"])
const message = ref('');

const newTask = ref<Task>(createEmptyTask())

const getMessage = async () => {
  message.value = await Greet(props.username);
};

const addTask = async () => {
  if (newTask.value.text.trim() !== '') {
    await createTask(newTask.value)
    newTask.value = createEmptyTask()
    emit("loadTasks")
  }
};

onMounted(() => {
  getMessage();
});

</script>

<template>
  <div>
    <h1>{{ message }}</h1>
    <h1>Список задач</h1>
    <div>
      <input v-model="newTask.text" placeholder="Введите новую задачу" />
      <input type="datetime-local" v-model="newTask.deadline" />
      <label>
        <input type="checkbox" v-model="newTask.hasPriority" />
        Важная задача
      </label>
      <span class="star-icon" role="button" @click="newTask.hasPriority = !newTask.hasPriority">
        {{ newTask.hasPriority ? '★' : '☆' }}
      </span>
      <button @click="addTask">Добавить задачу</button>
    </div>
  </div>
</template>

<style scoped>

</style>
