export type WSMessageType = "login" | "update" | "game-join"
export interface WSMessage {
  type: WSMessageType
  data: unknown
}
