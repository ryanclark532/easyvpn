export type serverStatus = "notRunning" | "running" | "starting" | "unknown"

export type ServerStatusResponse = {
    status: serverStatus
}

export const ServerStatusMapping: Record<serverStatus, string> = {
    running:"The VPN Server Is Running",
    notRunning:"The VPN Server Is Not Running",
    starting:"The VPN Server Is Starting",
    unknown: "We failed to get the Server Status, please try again"
}