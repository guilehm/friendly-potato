export type WSMessageType = "login" | "update"
export interface WSMessage {
  type: WSMessageType
  data: unknown
}
