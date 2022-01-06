export type Win = {
  defeated: Player
}

export type Player = {
  type: string
  username: string
  position_x: number
  position_y: number
  last_key: string
  last_direction: string
  steps: number
  wins: Array<Win>
  hp_total: number,
  hp: number,
}

export type BroadcastMessage = {
  type: string
  players: Array<Player>
}
