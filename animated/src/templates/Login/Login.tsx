import { useState } from 'react'
import * as S from './Login.styles'
import {
  Button,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  HStack,
  Input,
  InputGroup,
  InputRightElement,
  Stack,
} from '@chakra-ui/react'
import * as yup from 'yup'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { formatDiagnostic, getEffectiveTypeParameterDeclarations } from 'typescript'


type LoginFormInputs = {
  email: string
  password: string
}

const schema = yup.object().shape({
  email: yup.string().email().required(),
  password: yup.string().min(8).required(),
})


const Login = (): JSX.Element => {
  const { register, handleSubmit, formState: { errors } } = useForm<LoginFormInputs>({
    mode: 'onBlur',
    resolver: yupResolver(schema),
  })

  const onSubmit = (values: LoginFormInputs) => console.log('****', values)

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
          isInvalid={!!errors.password}
          {...register('password')}
        />
        <InputRightElement width="4.5rem">
          <Button h="1.75rem" size="sm" onClick={handleClick}>
            {show ? 'Hide' : 'Show'}
          </Button>
        </InputRightElement>
      </InputGroup>
    )
  }

  return (
    <S.Container maxW="xl" centerContent>
      <S.Section>

        <header>
          <S.Title>{'Let\'s sign you in'}</S.Title>
          <S.Paragraph>Welcome back!</S.Paragraph>
        </header>

        <main>
          <FormControl id="login" isInvalid={!!errors.email || !!errors.password}>
            <form onSubmit={handleSubmit(onSubmit)}>

              <Stack spacing={2}>

                <FormLabel display="none">Email address</FormLabel>
                <Input placeholder="email" {...register('email')} isInvalid={!!errors.email} />
                {errors.email && <FormErrorMessage>{errors.email.message}</FormErrorMessage>}
                <FormHelperText display="none">{'We\'ll never share your email.'}</FormHelperText>

                <FormLabel display="none">Password</FormLabel>
                <PasswordInput />
                {errors.password && <FormErrorMessage>{errors.password.message}</FormErrorMessage>}

                <HStack>
                  {/* <Button onClick={handleSubmit(onSubmit)} size="sm">Register</Button> */}
                  <Button type="submit" size="sm">Sign In</Button>
                </HStack>

              </Stack>
            </form>
          </FormControl>
        </main>

      </S.Section>
    </S.Container >
  )
}


export default Login
