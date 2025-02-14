  <script setup lang="ts">
  import {Task, TaskStatusEnum} from "../types/task";
  import {deleteTask, updateTask, updateTaskPriority, updateTaskStatus} from "../repositories/task";
  import {onMounted, onUnmounted, ref} from "vue";

  const props = defineProps<{ tasks: Task[] }>();
  const emit = defineEmits(["loadTasks"])

  const showRemoveModal = ref(false);
  const showEditModal = ref(false);

  const taskToRemove = ref<number | null>(null)
  const taskToEdit = ref<Task | null>(null);

  const toggleTaskStatus = async (task: Task, e: Event) => {
    e.stopPropagation(); // Чтобы клик не открыл редактирование
    const newStatus = task.status === TaskStatusEnum.TODO ? TaskStatusEnum.DONE : TaskStatusEnum.TODO;
    await updateTaskStatus(task.id, newStatus);
    emit("loadTasks")
  };

  const toggleTaskPriority = async (task: Task, e: Event) => {
    e.stopPropagation();
    await updateTaskPriority(task.id, !task.hasPriority)
    emit("loadTasks")
  }

  const removeTask = async (id: number) => {
    await deleteTask(id);
    emit("loadTasks")
  };

  const openRemoveModal = (id: number, e: Event) => {
    e.stopPropagation();
    taskToRemove.value = id
    showRemoveModal.value = true
  }

  const confirmRemove = async () => {
    if (taskToRemove.value !== null) {
      await removeTask(taskToRemove.value);
    }
    showRemoveModal.value = false;
  }

  const openEditModal = (task: Task) => {
    taskToEdit.value = {...task}
    showEditModal.value = true
  }

  const saveEditedTask = async() => {
    if(taskToEdit.value) {
      await updateTask(taskToEdit.value)
      emit("loadTasks")
      showEditModal.value = false
    }
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
        <div class="card mb-3 w-50 justify-content-center mx-auto custom-card" @click="openEditModal(task)">
          <div class="card-body d-flex justify-content-between align-items-center">
            <!-- Звездочка и текст задачи -->
            <div class="d-flex align-items-center">
              <span class="star-icon me-3" role="button"  @click="(e) => toggleTaskPriority(task, e)">
                {{ task.hasPriority ? '★' : '☆' }}
              </span>
              <span :class="['pointer', { done: task.status === 'done' }]" @click="(e) => toggleTaskStatus(task, e)">
                {{ task.text }}
              </span>
            </div>

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

        <div v-if="showRemoveModal" class="modal">
          <div class="modal-content">
            <p>Вы уверены, что хотите удалить задачу "{{ task.text }}"?</p>
            <button class="btn btn-outline-danger mb-2" @click="confirmRemove">Удалить</button>
            <button class="btn btn-outline-primary" @click="showRemoveModal = false">Отмена</button>
          </div>
        </div>

        <div v-if="showEditModal" class="modal">
          <div class="modal-content">
            <h3>Редактирование</h3>
            <input v-model="taskToEdit!.text" class="pt-2 pb-2 mb-2" type="text" placeholder="{{ taskToEdit!.text }}" />
            <input v-model="taskToEdit!.dueAt" class="pt-2 pb-2 mb-2" type="datetime-local" />
            <button class="btn btn-outline-success mb-2" @click="saveEditedTask">Сохранить</button>
            <button class="btn btn-outline-secondary" @click="showEditModal = false">Отмена</button>
          </div>
        </div>
      </li>
    </ul>
  </template>


  <style scoped>
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

  .modal-content h3 {
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