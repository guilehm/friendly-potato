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

