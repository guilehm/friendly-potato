import styled from "styled-components"

export const Container = styled.div`
  display: flex;
  height: 100vh;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  overflow: hidden;
`

export const Canvas = styled.canvas`
  border: 5px solid black;
  box-shadow: 0px 0px 20px -3px #000000;
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

export const StatsList = styled.ul`
  display: flex;

  & li {
    list-style-type: none;
    padding: 5px;
  }
`

export const ProgressContainer = styled.div`
  progress::-webkit-progress-bar {
    background-color: white;
  }
`

export const Progress = styled.progress`
  background-color: white;
  border: 1px solid black;
  height: 24px;
`

export const LifeLabel = styled.label`
  position: absolute;
  left: 50%;
  transform: translate(-50%, 0);
`


export const ProgressXPContainer = styled.div`
  margin: 0;
  padding: 0;
  progress::-webkit-progress-bar {
    background-color: white;
  }
  progress::-webkit-progress-value {
    background-color: yellow;
  }
`
export const ProgressXP = styled.progress`
  height: 5px;
  background-color: yellow;
  border: 1px solid black;
`
