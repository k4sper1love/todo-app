import {AddTask, DeleteTask, GetTasks, UpdateTaskPriority, UpdateTaskStatus} from "../../wailsjs/go/main/App";
import {Task, TaskStatus, TaskStatusEnum} from "../types/task";

export const fetchTasks = async (): Promise<Task[]> => {
    try {
    const data = await GetTasks();
    return data.map((task): Task => ({
        ...task,
        deadline: task.deadline?.Valid ? task.deadline.String : undefined,
        hasPriority: task.has_priority,
        status: (Object.values(TaskStatusEnum) as string[]).includes(task.status)
            ? (task.status as TaskStatus)
            : TaskStatusEnum.TODO
    }));
    } catch (error) {
        console.error("Failed to fetch tasks:", error)
        return []
    }
};

export const createTask = async (task: Task) => {
    try {
        await AddTask(task.text, task.deadline, task.hasPriority)
    } catch (error) {
        console.error("Failed to create task:", error)
    }
}

export const deleteTask = async (id: number) => {
    try {
        await DeleteTask(id)
    } catch (error) {
        console.error("Failed to delete task:", error)
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