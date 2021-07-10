import {
  Button,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  HStack,
  Input, Stack, useToast
} from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { AxiosError } from 'axios'
import { useForm } from 'react-hook-form'
import * as yup from 'yup'

import * as S from './Login.styles'
import { makeToastData } from '../../helpers'
import ApiService from '../../services/api-service'


const Api = new ApiService()

type LoginFormInputs = {
  email: string
  password: string
}

const schema = yup.object().shape({
  email: yup.string().email().required(),
  password: yup.string().min(8).required(),
})


const Login = (): JSX.Element => {
  const toast = useToast()
  const { register, handleSubmit, formState: { errors } } = useForm<LoginFormInputs>({
    mode: 'onBlur',
    resolver: yupResolver(schema),
  })


  const handleSuccess = (title: string, description = '') => toast(makeToastData({
    title, description,
  }))

  const handleError = (error: AxiosError, fallbackTitle: string) => toast(makeToastData({
    title: `${error.response?.data?.error || fallbackTitle}`,
    status: 'error'
  }))


  const onRegister = async (values: LoginFormInputs) => {
    Api.createUser(values.email, values.password)
      .then(() => handleSuccess('Account created', 'We\'ve created your account for you'))
      .catch(err => handleError(err, 'Failed to register'))
  }

  const onLogin = async (values: LoginFormInputs) => {
    Api.login(values.email, values.password)
      .then(() => handleSuccess('Logged in'))
      .catch(err => handleError(err, 'Failed to login'))
  }

  return (
    <S.Container maxW="xl" centerContent>
      <S.Section>

        <header>
          <S.Title>{'Let\'s sign you in'}</S.Title>
          <S.Paragraph>Welcome back!</S.Paragraph>
        </header>

        <main>
          <FormControl isInvalid={!!errors.email || !!errors.password}>
            <form>

              <Stack spacing={2}>

                <FormLabel display="none">Email address</FormLabel>
                <Input
                  id="email"
                  placeholder="email"
                  isInvalid={!!errors.email}
                  {...register('email')}
                />
                {errors.email && <FormErrorMessage>{errors.email.message}</FormErrorMessage>}
                <FormHelperText display="none">{'We\'ll never share your email.'}</FormHelperText>

                <FormLabel display="none">Password</FormLabel>
                <Input
                  id="password"
                  placeholder="password"
                  type="password"
                  isInvalid={!!errors.password}
                  {...register('password')}
                />
                {errors.password && <FormErrorMessage>{errors.password.message}</FormErrorMessage>}

                <HStack>
                  <Button onClick={handleSubmit(onRegister)} size="sm">Register</Button>
                  <Button onClick={handleSubmit(onLogin)} size="sm">Login</Button>
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
