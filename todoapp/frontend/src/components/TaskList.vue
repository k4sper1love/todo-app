<script setup lang="ts">
import {Task, TaskStatusEnum} from "../types/task";
import {deleteTask, updateTaskPriority, updateTaskStatus} from "../repositories/task";
import {onMounted, onUnmounted, ref} from "vue";

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

const isDeadlineExpired = (deadline: string) => {
  if (!deadline) return false;
  const now = new Date();
  const deadlineDate = new Date(deadline);
  return deadlineDate < now;
};

const formatDeadline = (deadline: string): string => {
  const date = new Date(deadline);
  return date.toLocaleString("ru-RU", {
    day: "numeric",
    month: "short",
    year: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  });
};

</script>

<template>
  <ul class="list-unstyled ">
    <li v-for="task in tasks" :key="task.id" :class="{ priority: task.hasPriority }">
      <!-- Карточка для каждой задачи -->
      <div class="card mb-3 w-50 justify-content-center mx-auto custom-card">
        <div class="card-body d-flex justify-content-between align-items-center">
          <!-- Звездочка и текст задачи -->
          <div class="d-flex align-items-center">
            <span class="star-icon me-3" role="button" @click="toggleTaskPriority(task)">
              {{ task.hasPriority ? '★' : '☆' }}
            </span>
            <span :class="['pointer', { done: task.status === 'done' }]" @click="toggleTaskStatus(task)">
              {{ task.text }}
            </span>
          </div>

          <div class="d-flex align-items-center">
            <span v-if="task.deadline" :class="['deadline', { expired: isDeadlineExpired(task.deadline) }]">
              {{ formatDeadline(task.deadline) }}
            </span>
            <span class="trash-icon ms-3" role="button" @click="openModal(task.id)">
              ✖️
            </span>
          </div>
        </div>
      </div>

      <div v-if="showModal && taskToDelete === task.id" class="modal">
        <div class="modal-content">
          <p>Вы уверены, что хотите удалить задачу "{{ task.text }}"?</p>
          <button class="btn btn-outline-danger mb-2" @click="confirmDelete">Удалить</button>
          <button class="btn btn-outline-primary" @click="showModal = false">Отмена</button>
        </div>
      </div>
    </li>
  </ul>
</template>


<style scoped>
.deadline {
  font-size: 0.9em;
  color: #888;
  margin-left: 5px;
}

.deadline.expired {
  color: #dc3545;
  font-weight: bold;
}

.custom-card {
  background-color: #1e1e1e;
  color: var(--text-color);
  border: 1px solid #333333;
  border-radius: 5px;
  transition: all 0.3s ease;
}

.custom-card:hover {
  box-shadow: 0 0 5px rgba(187, 134, 252, 0.5);
  outline: none;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: #1e1e1e;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  width: 300px;
  text-align: center;
}

.modal-content p {
  margin-bottom: 20px;
  font-size: 16px;
  color: white;
}


.pointer {
  cursor: pointer;
}

.done {
  text-decoration: line-through;
  color: gray;
  font-weight: normal;
}

.trash-icon {
  cursor: pointer;
  margin-left: 10px;
}

</style>