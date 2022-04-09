import styled from "styled-components"
import Spinner from "../Spinner"

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

export const SpriteContainer = styled.section`
  text-align: center;
`

export const SpriteList = styled.ul`
  display: flex;
  align-items: center;
  justify-content: center;
`

type SpriteItemType = {
  img: string
  x: number
  y: number
  ox: number
  oy: number

  w: number
  h: number
}

export const SpriteListItem = styled.li<SpriteItemType>`
  list-style-type: none;
  cursor: pointer;
  width: ${props => props.w * 10}px;
  height: ${props => props.h * 10}px;
  margin: 2px;
  background-image: ${props => `url(${props.img})`};
  background-position: ${props => `${((props.x) * -10)}px ${((props.y) * -10)}px}`};
  background-size: 1280px 1280px;
  image-rendering: optimizeSpeed;
  image-rendering: -moz-crisp-edges;
  image-rendering: -o-crisp-edges;
  image-rendering: -webkit-optimize-contrast;
  image-rendering: pixelated;
  image-rendering: optimize-contrast;
  -ms-interpolation-mode: nearest-neighbor;
`
