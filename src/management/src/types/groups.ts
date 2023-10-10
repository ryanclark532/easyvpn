

export type Group = {
    id: string;
    name: string;
    member_count: number;
    is_admin: boolean;
    enabled: boolean;
}


export type GetGroupsResponse = {
    groups : Group[];
}