import {Task, TaskStatus, TaskStatusEnum} from "../types/task";

export const createEmptyTask = (): Task => ({
    id: -1,
    text: "",
    status: "todo",
    dueAt: undefined,
    hasPriority: false,
});

export const mapTask = (task: any): Task => ({
    ...task,
    status: (Object.values(TaskStatusEnum) as string[]).includes(task.status)
        ? (task.status as TaskStatus)
        : TaskStatusEnum.TODO,
    hasPriority: task.has_priority,
    createdAt: task.created_at?.Valid ? task.created_at.String : undefined,
    dueAt: task.due_at?.Valid ? task.due_at.String : undefined,
    completedAt: task.completed_at?.Valid ? task.completed_at.String : undefined,
});