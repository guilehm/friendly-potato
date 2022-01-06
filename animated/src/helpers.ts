type ToastData = {
  title: string
  description?: string
  status?: "info" | "warning" | "success" | "error"
  duration?: number
  isClosable?: boolean
}

export const makeToastData = ({
  title,
  description = "",
  status = "success",
  duration = 5000,
  isClosable = true,
}: ToastData): ToastData => ({
  title,
  description,
  status,
  duration,
  isClosable,
})

export const getRandomInt = (min: number, max: number): number =>
  Math.floor(Math.random() * (max - min + 1)) + min


export const getCharacterSprite = (character: string, direction: string, step: number) => {
  let suffix = ""
  if (step % 2) {
    suffix += "_w"
  }
  if (direction === "ArrowLeft") {
    suffix += "_b"
  }
  return `/img/assets/characters/${character}${suffix}.png`
}

export const getMovingCharacterSprite = (character: string, direction: string, move: boolean) => {
  let suffix = ""
  if (move) {
    suffix += "_w"
  }
  if (direction === "ArrowLeft") {
    suffix += "_b"
  }
  return `/img/assets/characters/${character}${suffix}.png`
}


export const getCursorPosition = (canvas: HTMLCanvasElement, event: MouseEvent) => {
  const rect = canvas.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top
  return [x, y]
}
