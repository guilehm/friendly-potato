import styled from 'styled-components'


export const Section = styled.section`
transition: ease-in-out 100ms;
  &:hover {
    transform: scale(1.05);
  }
`

export const Title = styled.h2`
  font-family: 'Press Start 2P', cursive;
  font-size: 0.8rem;
`

export const Image = styled.img`
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 12px 12px 0 0;
`
