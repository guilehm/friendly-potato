import styled from 'styled-components'

export const HomeContainer = styled.div`
  padding-bottom: 20px;
`

export const HomeSection = styled.section`
  margin: auto;
  max-width: 980px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  grid-gap: 20px;

`

export const Title = styled.h1`
  text-align: center;
  font-family: 'Press Start 2P', cursive;
  font-size: 1.4rem;
  margin-bottom: 10px;
`
