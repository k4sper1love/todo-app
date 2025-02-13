export interface Task{
    id: number;
    text: string;
    status: "todo" | "in_progress" | "done";
    deadline?: string;
    hasPriority: boolean,
}

export const createEmptyTask = (): Task => ({
    id: -1,
    text: "",
    status: "todo",
    deadline: undefined,
    hasPriority: false,
});