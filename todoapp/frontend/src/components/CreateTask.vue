<script setup lang="ts">
import { onMounted, ref } from 'vue';
import {GetTasks, AddTask, RemoveTask, Greet, UpdateTaskStatus, UpdateTaskPriority} from '../../wailsjs/go/main/App';

type Task = {
  id: number;
  text: string;
  status: "todo" | "in_progress" | "done";
  deadline?: string;
  hasPriority: boolean,
};

const props = defineProps<{ username: string }>();

const message = ref('');
const tasks = ref<Task[]>([]);

const newTask = ref<Task>({
  id: -1,
  text: "",
  status: "todo",
  deadline: undefined,
  hasPriority: false
});

const getMessage = async () => {
  message.value = await Greet(props.username);
};

const addTask = async () => {
  if (newTask.value.text.trim() !== '') {
    await AddTask(newTask.value.text, newTask.value.deadline, newTask.value.hasPriority);
    newTask.value.text = "";
    newTask.value.deadline = undefined;
    newTask.value.hasPriority = false;
    await loadTasks();
  }
};

const removeTask = async (id: number) => {
  await RemoveTask(id);
  await loadTasks();
};

const loadTasks = async () => {
  const data = (await GetTasks()).map(task => ({
    ...task,
    deadline: task.deadline?.Valid ? task.deadline.String : null,
    hasPriority: task.has_priority,
  }));

  tasks.value = data as unknown as Task[]
};

const toggleTaskStatus = async (task: Task) => {
  const newStatus = task.status === "todo" ? "done" : "todo";
  await UpdateTaskStatus(task.id, newStatus);
  task.status = newStatus;
};

const toggleTaskPriority = async (task: Task) => {
  await UpdateTaskPriority(task.id, !task.hasPriority)
  task.hasPriority = !task.hasPriority
}

onMounted(() => {
  getMessage();
  loadTasks();
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
      <button @click="addTask">Добавить задачу</button>
    </div>
    <ul>
      <li v-for="(task, index) in tasks" :key="task.id" :class="{ priority: task.hasPriority }">
        <span :class="['pointer', { done: task.status === 'done' }]" @click="toggleTaskStatus(task)">
          {{ task.text }} <span v-if="task.deadline">({{ task.deadline }})</span>
        </span>
        <span class="star-icon" role="button" @click="toggleTaskPriority(task)"> {{ task.hasPriority ? '★' : '☆' }}</span>
        <span class="trash-icon" role="button" @click="removeTask(task.id)">❌</span>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.pointer {
  cursor: pointer;
}

.done {
  text-decoration: line-through;
  color: gray;
  font-weight: normal;
}

.priority {
  font-weight: bold;
  color: red;
}

.star-icon {
  cursor: pointer;
  font-size: 20px;
  margin-left: 10px;
  color: gold;
}

.trash-icon {
  cursor: pointer;
  margin-left: 10px;
}
</style>
