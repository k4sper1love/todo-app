<script setup lang="ts">
import { ref } from "vue";
import { Task, TaskStatusEnum } from "../types/task";
import {deleteTask, updateTaskPriority, updateTaskStatus, updateTask,} from "../services/task";

// Define props for task list
const props = defineProps<{ tasks: Task[] }>();

// Define emitted events
const emit = defineEmits(["loadTasks"]);

// State variables
const showRemoveModal = ref(false); // Toggle delete confirmation modal
const showEditModal = ref(false); // Toggle edit modal
const taskToRemove = ref<number | null>(null); // Store task ID for removal
const taskToEdit = ref<Task | null>(null); // Store task for editing

// Toggle task completion status
const toggleTaskStatus = async (task: Task, e: Event) => {
  e.stopPropagation(); // Prevent opening edit modal
  const newStatus =
      task.status === TaskStatusEnum.TODO ? TaskStatusEnum.DONE : TaskStatusEnum.TODO;
  await updateTaskStatus(task.id, newStatus);
  emit("loadTasks");
};

// Toggle task priority
const toggleTaskPriority = async (task: Task, e: Event) => {
  e.stopPropagation();
  await updateTaskPriority(task.id, !task.hasPriority);
  emit("loadTasks");
};

// Delete a task
const removeTask = async (id: number) => {
  await deleteTask(id);
  emit("loadTasks");
};

// Open delete confirmation modal
const openRemoveModal = (id: number, e: Event) => {
  e.stopPropagation();
  taskToRemove.value = id;
  showRemoveModal.value = true;
};

// Confirm task removal
const confirmRemove = async () => {
  if (taskToRemove.value !== null) {
    await removeTask(taskToRemove.value);
  }
  showRemoveModal.value = false;
};

// Open task edit modal
const openEditModal = (task: Task) => {
  taskToEdit.value = { ...task };
  showEditModal.value = true;
};

// Save edited task
const saveEditedTask = async () => {
  if (taskToEdit.value) {
    await updateTask(taskToEdit.value);
    emit("loadTasks");
    showEditModal.value = false;
  }
};

// Check if the deadline has expired
const isDeadlineExpired = (deadline: string) => {
  if (!deadline) return false;
  return new Date(deadline) < new Date();
};

// Format deadline for display
const formatDeadline = (deadline: string): string => {
  return new Date(deadline).toLocaleString("ru-RU", {
    day: "numeric",
    month: "short",
    year: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  });
};
</script>

<template>
  <ul class="list-unstyled">
    <li v-for="task in tasks" :key="task.id" :class="{ priority: task.hasPriority }">
      <!-- Task card -->
      <div class="card mb-3 w-50 mx-auto custom-card" @click="openEditModal(task)">
        <div class="card-body d-flex justify-content-between align-items-center">
          <!-- Task priority star and text -->
          <div class="d-flex align-items-center">
            <span class="star-icon me-3" role="button" @click="(e) => toggleTaskPriority(task, e)">
              {{ task.hasPriority ? "★" : "☆" }}
            </span>
            <span
                :class="['pointer', { done: task.status === TaskStatusEnum.DONE }]"
                @click="(e) => toggleTaskStatus(task, e)"
            >
              {{ task.text }}
            </span>
          </div>

          <!-- Task due date / completion time -->
          <div class="d-flex align-items-center">
            <span v-if="task.completedAt" class="time completed">
              {{ formatDeadline(task.completedAt) }}
            </span>
            <span v-else-if="task.dueAt" :class="['time', { expired: isDeadlineExpired(task.dueAt) }]">
              {{ formatDeadline(task.dueAt) }}
            </span>
            <span class="trash-icon ms-3" role="button" @click="(e) => openRemoveModal(task.id, e)">
              ✖️
            </span>
          </div>
        </div>
      </div>

      <!-- Delete confirmation modal -->
      <div v-if="showRemoveModal" class="modal">
        <div class="modal-content">
          <p>Вы уверены, что хотите удалить задачу?</p>
          <button class="btn btn-outline-danger mb-2" @click="confirmRemove">Удалить</button>
          <button class="btn btn-outline-primary" @click="showRemoveModal = false">Отмена</button>
        </div>
      </div>

      <!-- Task edit modal -->
      <div v-if="showEditModal" class="modal">
        <div class="modal-content">
          <h3>Редактирование</h3>
          <input v-model="taskToEdit!.text" class="pt-2 pb-2 mb-2" type="text" placeholder="Введите текст задачи" />
          <input v-model="taskToEdit!.dueAt" class="pt-2 pb-2 mb-2" type="datetime-local" />
          <button class="btn btn-outline-success mb-2" @click="saveEditedTask">Сохранить</button>
          <button class="btn btn-outline-secondary" @click="showEditModal = false">Отмена</button>
        </div>
      </div>
    </li>
  </ul>
</template>



<style scoped>
/* Time styling */
.time {
  font-size: 0.8em;
  color: #888;
  margin-left: 5px;
}

.time.expired {
  color: #dc3545;
  font-weight: bold;
}

.time.completed {
  color: var(--success-color);
}

/* Task card styles */
.custom-card {
  background-color: #1e1e1e;
  color: var(--text-color);
  border: 1px solid #333;
  border-radius: 5px;
  transition: all 0.3s ease;
}

.custom-card:hover {
  box-shadow: 0 0 5px rgba(187, 134, 252, 0.5);
}

/* Modal styles */
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

.modal-content p,
.modal-content h3 {
  color: white;
}

/* Miscellaneous */
.pointer {
  cursor: pointer;
}

.done {
  text-decoration: line-through;
  color: gray;
}

.trash-icon {
  cursor: pointer;
  margin-left: 10px;
}
</style>