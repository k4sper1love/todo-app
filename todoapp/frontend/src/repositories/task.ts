import {
    AddTask,
    DeleteTask,
    GetActiveTasks,
    GetCompletedTasks,
    UpdateTaskPriority,
    UpdateTaskStatus,
    UpdateTask
} from "../../wailsjs/go/main/App";
import {Task, TaskStatus} from "../types/task";
import {mapTask} from "../utils/task";

const fetchTasks = async (fetchMethod: () => Promise<any[]>): Promise<Task[]> => {
    try {
        const data = await fetchMethod();
        return data.map(mapTask);
    } catch (error) {
        console.error("Failed to fetch tasks:", error);
        return [];
    }
};

export const fetchActiveTasks = () => fetchTasks(GetActiveTasks)

export const fetchCompletedTasks = () => fetchTasks(GetCompletedTasks)

export const createTask = async (task: Task) => {
    try {
        await AddTask(task.text, task.dueAt, task.hasPriority)
    } catch (error) {
        console.error("Failed to create task:", error)
    }
}

export const updateTask = async(task: Task) => {
    try {
        await UpdateTask(task.id, task.text, task.dueAt)
    } catch (error) {
        console.error("Failed to update task:", error)
    }
}

export const updateTaskStatus = async (id: number, status: TaskStatus) => {
    try {
        await UpdateTaskStatus(id, status)
    } catch (error) {
        console.error("Failed to update task status:", error)
    }
}

export const updateTaskPriority = async (id: number, hasPriority: boolean) => {
    try {
        await UpdateTaskPriority(id, hasPriority)
    } catch (error) {
        console.error("Failed to update task priority:", error)
    }
}


export const deleteTask = async (id: number) => {
    try {
        await DeleteTask(id)
    } catch (error) {
        console.error("Failed to delete task:", error)
    }
}