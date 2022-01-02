import styled from "styled-components"

export const Container = styled.div`
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  z-index: 100;
  /* background-color: yellow; */
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
