import { useState } from 'react'
import * as S from './Login.styles'
import {
  Button,
  FormControl,
  FormHelperText,
  FormLabel,
  HStack,
  Input,
  InputGroup,
  InputRightElement,
  Stack,
} from '@chakra-ui/react'


function PasswordInput() {
  const [show, setShow] = useState(false)
  const handleClick = () => setShow(!show)

  return (
    <InputGroup size="md">
      <Input
        pr="4.5rem"
        type={show ? 'text' : 'password'}
        placeholder="password"
        id="password"
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
  <S.Container maxW="xl" centerContent>
    <S.Section>

      <header>
        <S.Title>{'Let\'s sign you in'}</S.Title>
        <S.Paragraph>Welcome back!</S.Paragraph>
      </header>

      <main>
        <FormControl id="email">
          <Stack spacing={2}>
            <FormLabel display="none">Email address</FormLabel>
            <Input type="email" placeholder="email" />
            <FormHelperText display="none">{'We\'ll never share your email.'}</FormHelperText>
            <FormLabel display="none">Password</FormLabel>
            <PasswordInput />

            <HStack>
              <Button size="sm">Register</Button>
              <Button size="sm">Sign In</Button>
            </HStack>
          </Stack>
        </FormControl>
      </main>

    </S.Section>
  </S.Container>
)


export default Login
