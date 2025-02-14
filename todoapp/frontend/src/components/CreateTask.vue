<script setup lang="ts">
import {ref} from 'vue';
import {Task} from "../types/task";
import {createTask} from "../services/task";
import {createEmptyTask} from "../utils/task";

// Define emitted events
const emit = defineEmits(["loadTasks"]);

// Reactive variable for new task
const newTask = ref<Task>(createEmptyTask());

// Function to add a new task
const addTask = async () => {
  if (!newTask.value.text.trim()) return; // Prevent empty tasks

  await createTask(newTask.value);
  newTask.value = createEmptyTask();
  emit("loadTasks");
};

// Function to trigger the datetime picker
const openDateTimePicker = () => {
  const input = document.getElementById('taskDeadline') as HTMLInputElement | null;
  if (input) {
    input.showPicker();
  }
};
</script>

<template>
  <div class="input-group justify-content-center mt-3 mb-5">
    <!-- Priority toggle -->
    <span class="star-icon input-group-text" role="button" @click="newTask.hasPriority = !newTask.hasPriority">
      {{ newTask.hasPriority ? "★" : "☆" }}
    </span>

    <!-- Task input -->
    <input class="w-25" v-model="newTask.text" placeholder="Введите новую задачу" />

    <!-- Date picker -->
    <div class="input-group-text">
      <span class="me-2 calendar-icon" @click="openDateTimePicker">📅</span>
      <input v-model="newTask.dueAt" id="taskDeadline" type="datetime-local" placeholder="Выберите дату и время" />
    </div>

    <!-- Add task button -->
    <button class="btn btn-outline-primary" type="button" @click="addTask">Добавить задачу</button>
  </div>
</template>

<style scoped>
.star-icon {
  cursor: pointer;
  font-size: 1.5rem;
  color: var(--primary-color);
  transition: color 0.3s ease;
}

.star-icon:hover {
  color: var(--secondary-color);
}

.calendar-icon {
  cursor: pointer;
}

input[type="datetime-local"]::-webkit-calendar-picker-indicator,
input[type="datetime-local"]::-ms-clear {
  display: none;
}
</style>