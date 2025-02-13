<script setup lang="ts">
import {onMounted, ref} from 'vue';
import {Task} from "../types/task";
import {createTask} from "../repositories/task";
import {createEmptyTask} from "../utils/task";

const emit = defineEmits(["loadTasks"])

const newTask = ref<Task>(createEmptyTask());

const addTask = async () => {
  if (newTask.value.text.trim() !== '') {
    await createTask(newTask.value)
    newTask.value = createEmptyTask()
    emit("loadTasks")
  }
};

const openDateTimePicker = () => {
  const input = document.getElementById('taskDeadline') as HTMLInputElement | null;
  if (input) {
    input.showPicker();
  }
};

</script>

<template>
    <div class="input-group justify-content-center mt-3 mb-5">
      <span class="star-icon input-group-text" role="button" @click="newTask.hasPriority = !newTask.hasPriority">
       {{ newTask.hasPriority ? '★' : '☆' }}
      </span>
      <input class="w-25" v-model="newTask.text" placeholder="Введите новую задачу"/>
      <div class="input-group-text">
        <span class="me-2 calendar-icon" @click="openDateTimePicker">📅</span>
        <input v-model="newTask.deadline" id="taskDeadline" type="datetime-local" placeholder="Выберите дату и время">
      </div>
      <button class="btn btn-outline-primary" type="button" @click="addTask">Добавить задачу</button>
    </div>
</template>

<style scoped>
.datetime {
  font-family: inherit;
  transition: all 0.3s ease;
}

input[type="datetime-local"]::-webkit-calendar-picker-indicator {
  display: none;
}

input[type="datetime-local"]::-ms-clear {
  display: none;
}

.calendar-icon {
  cursor: pointer;
}

</style>
