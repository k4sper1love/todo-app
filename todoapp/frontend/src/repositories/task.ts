import {AddTask, DeleteTask, GetTasks, UpdateTaskPriority, UpdateTaskStatus} from "../../wailsjs/go/main/App";
import {Task} from "../types/task";

export const fetchTasks = async(): Promise<Task[]> => {
    const data = (await GetTasks()).map(task => ({
        ...task,
        deadline: task.deadline?.Valid ? task.deadline.String : null,
        hasPriority: task.has_priority
    }));

    return data as unknown as Task[]
}

export const createTask = async (task: Task) => {
    await AddTask(task.text, task.deadline, task.hasPriority)
}

export const deleteTask = async (id: number) => {
    await DeleteTask(id)
}

export const updateTaskStatus = async (id: number, status: "todo" | "done") => {
    await UpdateTaskStatus(id, status)
}

export const updateTaskPriority = async (id: number, hasPriority: boolean) => {
    await UpdateTaskPriority(id, hasPriority)
}