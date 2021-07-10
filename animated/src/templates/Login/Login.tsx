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
  const { register, handleSubmit, formState: { errors } } = useForm<LoginFormInputs>({
    mode: 'onBlur',
    resolver: yupResolver(schema),
  })

  const toast = useToast()


  const onRegister = async (values: LoginFormInputs) => {
    const handleSuccess = () => toast(makeToastData({
      title: 'Account created.',
      description: 'We\'ve created your account for you.',
    }))

    const handleError = (error: AxiosError) => toast(makeToastData({
      title: `${error.response?.data?.error || 'Failed to register'}`,
      status: 'error'
    }))

    Api.createUser(values.email, values.password)
      .then(handleSuccess)
      .catch(handleError)
  }

  const onLogin = async (values: LoginFormInputs) => {
    const handleSuccess = () => toast(makeToastData({
      title: 'Logged in.',
    }))

    const handleError = (error: AxiosError) => toast(makeToastData({
      title: `${error.response?.data?.error || 'Failed to login'}`,
      status: 'error',
    }))

    Api.login(values.email, values.password)
      .then(handleSuccess)
      .catch(handleError)
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
