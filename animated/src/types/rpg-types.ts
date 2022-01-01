export type Player = {
  type: string
  username: string
  position_x: number
  position_y: number
}

export type BroadcastMessage = {
  type: string
  players: Array<Player>
}
