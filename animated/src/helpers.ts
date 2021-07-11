type ToastData = {
  title: string
  description?: string
  status?: "info" | "warning" | "success" | "error"
  duration?: number
  isClosable?: boolean
}

const makeToastData = ({
  title,
  description = "",
  status = "success",
  duration = 5000,
  isClosable = true,
}: ToastData): ToastData => ({
  title: title,
  description,
  status,
  duration,
  isClosable,
})

export { makeToastData }
