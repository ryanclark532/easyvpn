export type serverStatus = "notRunning" | "running" | "starting"

export type ServerStatusResponse = {
    status: serverStatus
}

export const ServerStatusMapping: Record<serverStatus, string> = {
    running:"The VPN Server Is Running",
    notRunning:"The VPN Server Is Not Running",
    starting:"The VPN Server Is Starting"
}