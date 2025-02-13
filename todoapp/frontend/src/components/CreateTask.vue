<script setup lang="ts">
import { onMounted, ref } from 'vue';
import {GetTasks, AddTask, DeleteTask, Greet, UpdateTaskStatus, UpdateTaskPriority} from '../../wailsjs/go/main/App';

type Task = {
  id: number;
  text: string;
  status: "todo" | "in_progress" | "done";
  deadline?: string;
  hasPriority: boolean,
};

const props = defineProps<{ username: string }>();

const showModal = ref(false);
const taskToDelete = ref<number | null>(null)
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

const deleteTask = async (id: number) => {
  await DeleteTask(id);
  await loadTasks();
};

const confirmDelete = async () => {
  if (taskToDelete.value !== null) {
    await deleteTask(taskToDelete.value);
  }
  showModal.value = false;
}

const openModal = (id: number) => {
  taskToDelete.value = id
  showModal.value = true
}

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
        <span class="trash-icon" role="button" @click="openModal(task.id)">❌</span>

        <div v-if="showModal" class="modal">
          <div class="modal-content">
            <p>Вы уверены, что хотите удалить эту задачу?</p>
            <button @click="confirmDelete">Удалить</button>
            <button @click="showModal = false">Отмена</button>
          </div>
        </div>
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

.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: rgba(0, 0, 0, 0.4);
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  text-align: center;
  width: 300px;
}

.modal-content p {
  margin-bottom: 15px;
  font-size: 16px;
  color: black;
}

.modal-content button {
  margin: 5px;
  padding: 8px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
}

.modal-content button:first-of-type {
  background: #e63946;
  color: white;
}

.modal-content button:last-of-type {
  background: #ccc;
  color: black;
}

.modal-content button:hover {
  opacity: 0.8;
}

</style>
