import {Task} from "../types/task";

export const createEmptyTask = (): Task => ({
    id: -1,
    text: "",
    status: "todo",
    deadline: undefined,
    hasPriority: false,
});

