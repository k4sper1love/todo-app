import {Profile} from "../types/profile";

export const createEmptyProfile = (): Profile => ({
    id: -1,
    name: "",
});

export const mapProfile = (profile: any): Profile => ({
    ...profile,
    created_at: profile.created_at?. Valid ? profile.created_at.String : undefined,
});