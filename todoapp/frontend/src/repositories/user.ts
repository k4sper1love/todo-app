import {GetUsername, SetUsername} from "../../wailsjs/go/main/App";

export const fetchUsername = async ():Promise<string> => {
    return await GetUsername()
}

export const setUsername = async (username: string) => {
    await SetUsername(username)
}