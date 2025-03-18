// Service layer for task-related API interactions with Go backend

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

// Fetch tasks with error handling
const fetchTasks = async (profileId: number, fetchMethod: (profileId: number) => Promise<any[]>): Promise<Task[]> => {
    try {
        const data = await fetchMethod(profileId);
        return data.map(mapTask);
    } catch (error) {
        console.error("Failed to fetch tasks:", error);
        return [];
    }
};

export const fetchActiveTasks = (profileId: number) => fetchTasks(profileId, GetActiveTasks)
export const fetchCompletedTasks = (profileId: number) => fetchTasks(profileId, GetCompletedTasks)

// Create a new task
export const createTask = async (profile_id: number, task: Task) => {
    try {
        await AddTask(profile_id, task.text, task.dueAt, task.hasPriority);
    } catch (error) {
        console.error("Failed to create task:", error);
    }
};

// Update task details
export const updateTask = async(task: Task) => {
    try {
        await UpdateTask(task.id, task.text, task.dueAt);
    } catch (error) {
        console.error("Failed to update task:", error);
    }
};

// Update task status
export const updateTaskStatus = async (id: number, status: TaskStatus) => {
    try {
        await UpdateTaskStatus(id, status);
    } catch (error) {
        console.error("Failed to update task status:", error);
    }
};

// Update task priority
export const updateTaskPriority = async (id: number, hasPriority: boolean) => {
    try {
        await UpdateTaskPriority(id, hasPriority);
    } catch (error) {
        console.error("Failed to update task priority:", error);
    }
};

// Delete a task
export const deleteTask = async (id: number) => {
    try {
        await DeleteTask(id);
    } catch (error) {
        console.error("Failed to delete task:", error);
    }
};