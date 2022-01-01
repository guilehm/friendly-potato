export type Player = {
  type: string
  username: string
  position_x: number
  position_y: number
  last_key: string
  last_direction: string
}

export type BroadcastMessage = {
  type: string
  players: Array<Player>
}
