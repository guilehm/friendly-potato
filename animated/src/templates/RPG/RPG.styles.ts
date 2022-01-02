import styled from "styled-components"
import { Container as ChakraContainer } from "@chakra-ui/react"

export const Container = styled(ChakraContainer)``

export const GameContainer = styled.div`
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  /* z-index: 10; */
`

export const Canvas = styled.canvas`
  border: 1px solid black;
  /* box-shadow: 0px 0px 10px -3px #000000; */
  margin: auto;
  width: 100%;
  max-width: 80vh;
  max-height: 80vh;
  image-rendering: optimizeSpeed; /* STOP SMOOTHING, GIVE ME SPEED  */
  image-rendering: -moz-crisp-edges; /* Firefox                        */
  image-rendering: -o-crisp-edges; /* Opera                          */
  image-rendering: -webkit-optimize-contrast; /* Chrome (and eventually Safari) */
  image-rendering: pixelated; /* Chrome */
  image-rendering: optimize-contrast; /* CSS3 Proposed                  */
  -ms-interpolation-mode: nearest-neighbor;
`

export const KeysContainer = styled.div`
  position: absolute;
  text-align: center;
  left: 0;
  width: 100%;
  z-index: 20;
`

export const Title = styled.h2`
  font-size: 3.4rem;
  line-height: 60px;
  padding-bottom: 20px;
`

export const Paragraph = styled.p`
  font-size: 1.8rem;
`
