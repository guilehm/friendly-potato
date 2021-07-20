import { Alert, AlertDescription, AlertIcon, AlertProps, AlertTitle } from "@chakra-ui/react"


type AlertCompositionType = AlertProps & {
  description: string
}


const AlertComposition = ({ title, description, status }: AlertCompositionType): JSX.Element => {
  return <Alert
    status={status}
    variant="subtle"
    flexDirection="column"
    alignItems="center"
    justifyContent="center"
    textAlign="center"
    height="200px"
  >
    <AlertIcon boxSize="40px" mr={0} />
    <AlertTitle mt={4} mb={1} fontSize="lg">
      {title}
    </AlertTitle>
    <AlertDescription maxWidth="sm">
      {description}
    </AlertDescription>
  </Alert>
}


export default AlertComposition
