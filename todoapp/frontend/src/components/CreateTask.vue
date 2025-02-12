<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetTasks, AddTask, RemoveTask, Greet, UpdateTaskStatus } from '../../wailsjs/go/main/App';

type Task = {
  id: number;
  text: string;
  status: "todo" | "in_progress" | "done";
  deadline?: string;
};

const props = defineProps<{ username: string }>();

const message = ref('');
const newTask = ref<Task>({id: -1, text: "", status: "todo", deadline: undefined });
const tasks = ref<Task[]>([]);

const getMessage = async () => {
  message.value = await Greet(props.username);
};

const addTask = async () => {
  if (newTask.value.text.trim() !== '') {
    await AddTask(newTask.value.text, newTask.value.deadline);
    newTask.value.text = "";
    newTask.value.deadline = undefined;
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
    deadline: task.deadline?.Valid ? task.deadline.String : null
  }));

  tasks.value = data as unknown as Task[]
};

const toggleTaskStatus = async (task: Task) => {
  const newStatus = task.status === "todo" ? "done" : "todo";
  await UpdateTaskStatus(task.id, newStatus);
  task.status = newStatus;
};

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
      <input type="date" v-model="newTask.deadline" />
      <button @click="addTask">Добавить задачу</button>
    </div>
    <ul>
      <li v-for="(task, index) in tasks" :key="task.id">
        <span :class="['pointer', { done: task.status === 'done' }]" @click="toggleTaskStatus(task)">
          {{ task.text }} <span v-if="task.deadline">({{ task.deadline }})</span>
        </span>
        <button @click="removeTask(task.id)">Удалить</button>
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
}
</style>
