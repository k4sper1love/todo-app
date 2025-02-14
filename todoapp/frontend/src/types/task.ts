export const TaskStatusEnum = {
    TODO: "todo",
    IN_PROGRESS: "in_progress",
    DONE: "done",
} as const;

export type TaskStatus = (typeof TaskStatusEnum)[keyof typeof TaskStatusEnum]

export interface Task {
    id: number;
    text: string;
    status: TaskStatus;
    hasPriority: boolean;
    createdAt?: string,
    dueAt?: string,
    completedAt?: string
}
