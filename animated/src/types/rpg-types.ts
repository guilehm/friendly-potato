export type Win = {
  defeated: Player
}

export type Player = {
  type: string
  username: string
  positionX: number
  positionY: number
  lastKey: string
  lastDirection: string
  steps: number
  wins: Array<Win>
  hpTotal: number,
  hp: number,
}

export type BroadcastMessage = {
  type: string
  players: Array<Player>
}
