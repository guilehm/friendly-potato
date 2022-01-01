export type WSMessageType = "login" | "update" | "game-join" | "key-down"
export interface WSMessage {
  type: WSMessageType
  data: unknown
}
