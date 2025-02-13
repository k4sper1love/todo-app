<script setup lang="ts">
import {Task, TaskStatusEnum} from "../types/task";
import {deleteTask, updateTaskPriority, updateTaskStatus} from "../repositories/task";
import {ref} from "vue";

const props = defineProps<{ tasks: Task[] }>();
const emit = defineEmits(["loadTasks"])

const showModal = ref(false);
const taskToDelete = ref<number | null>(null)


const toggleTaskStatus = async (task: Task) => {
  const newStatus = task.status === TaskStatusEnum.TODO ? TaskStatusEnum.DONE : TaskStatusEnum.TODO;
  await updateTaskStatus(task.id, newStatus);
  emit("loadTasks")
};

const toggleTaskPriority = async (task: Task) => {
  await updateTaskPriority(task.id, !task.hasPriority)
  emit("loadTasks")
}

const removeTask = async (id: number) => {
  await deleteTask(id);
  emit("loadTasks")
};

const confirmDelete = async () => {
  if (taskToDelete.value !== null) {
    await removeTask(taskToDelete.value);
  }
  showModal.value = false;
}

const openModal = (id: number) => {
  taskToDelete.value = id
  showModal.value = true
}

</script>

<template>
  <div>
    <ul>
      <li v-for="task in tasks" :key="task.id" :class="{ priority: task.hasPriority }">
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

.trash-icon {
  cursor: pointer;
  margin-left: 10px;
}

</style>