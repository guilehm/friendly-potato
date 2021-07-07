import * as S from './Login.styles'
import {
  FormControl,
  Input,
  FormLabel,
  FormHelperText,
  InputGroup,
  InputRightElement,
  Button,
  Stack,
  HStack,
} from '@chakra-ui/react'
import { useState } from 'react'

function PasswordInput() {
  const [show, setShow] = useState(false)
  const handleClick = () => setShow(!show)

  return (
    <InputGroup size="md">
      <Input
        pr="4.5rem"
        type={show ? 'text' : 'password'}
        placeholder="Enter password"
      />
      <InputRightElement width="4.5rem">
        <Button h="1.75rem" size="sm" onClick={handleClick}>
          {show ? 'Hide' : 'Show'}
        </Button>
      </InputRightElement>
    </InputGroup>
  )
}

const Login = (): JSX.Element => (
  <S.Container>
    <S.Section>
      <header>
        <S.Title>{'Let\'s sign you in'}</S.Title>
        <S.Paragraph>Welcome back!</S.Paragraph>
      </header>
      <main>

        <FormControl id="email">
          <Stack spacing={2}>
            <FormLabel>Email address</FormLabel>
            <Input type="email" />
            <FormHelperText>{'We\'ll never share your email.'}</FormHelperText>
            <FormLabel>Password</FormLabel>
            <PasswordInput />

            <HStack>
              <Button size="sm">Register</Button>
              <Button size="sm">Sign In</Button>
            </HStack>
          </Stack>

        </FormControl>

      </main>
      <footer>
      </footer>
    </S.Section>
  </S.Container>
)


export default Login
