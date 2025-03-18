// Enum for task statuses
export const TaskStatusEnum  = {
    TODO: "todo",
    IN_PROGRESS: "in_progress",
    DONE: "done",
} as const;

// Type for Task status, that ensures only valid enum values are used
export type TaskStatus = (typeof TaskStatusEnum)[keyof typeof TaskStatusEnum];

// Interface for a Task object
export interface Task {
    id: number; // Unique identifier
    profileID: number;
    text: string; // Task description
    status: TaskStatus; // Current task status
    hasPriority: boolean; // Indicates if the task has high priority
    createdAt?: string; // Optional: Timestamp when task was created
    dueAt?: string; // Optional: Due date and time
    completedAt?: string; // Optional: Timestamp when task was completed
}
