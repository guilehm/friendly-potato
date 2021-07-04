import styled from 'styled-components'


export const Section = styled.section`
  transition: ease 500ms;
  &:hover {
    transform: scale(1.05);
  }
`

export const Title = styled.h2`
  font-family: 'Press Start 2P', cursive;
  font-size: 0.8rem;
  margin: 5px 0;
`

export const Image = styled.img`
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 12px 12px 0 0;
`
