import styled from "styled-components"

export const Container = styled.div`
  display: flex;
  height: 100vh;
  justify-content: center;
  align-items: center;
  flex-direction: column;
`

export const Canvas = styled.canvas`
  border: 5px solid black;
  box-shadow: 0px 0px 20px -3px #000000;
  /* margin: auto; */
  width: 100%;
  max-width: 720px;
  image-rendering: optimizeSpeed; /* STOP SMOOTHING, GIVE ME SPEED  */
  image-rendering: -moz-crisp-edges; /* Firefox                        */
  image-rendering: -o-crisp-edges; /* Opera                          */
  image-rendering: -webkit-optimize-contrast; /* Chrome (and eventually Safari) */
  image-rendering: pixelated; /* Chrome */
  image-rendering: optimize-contrast; /* CSS3 Proposed                  */
  -ms-interpolation-mode: nearest-neighbor;
`

export const ArrowContainer = styled.div`
  text-align: center;
  cursor: pointer;
`
