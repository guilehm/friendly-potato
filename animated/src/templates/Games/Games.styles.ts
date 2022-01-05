import styled from "styled-components"

export const GamesContainer = styled.div`
  padding-bottom: 20px;
`

export const GamesSection = styled.section`
  margin: auto;
  max-width: 980px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  grid-gap: 20px;
`

export const Title = styled.h1`
  text-align: center;
  font-size: 1.4rem;
  margin-bottom: 10px;
`
