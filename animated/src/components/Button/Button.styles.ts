import styled from "styled-components"

type ButtonType = {
  color?: string
  bgColor?: string
}

export const Button = styled.a<ButtonType>`
  display: inline-block;
  font-size: 0.75rem;
  font-weight: 700;
  padding: 0.35em 0.65em;
  color: ${(props) => props.color || "white"};
  background-color: ${(props) => props.bgColor || "black"};
  text-align: center;
  border-radius: 0.25rem;
  white-space: nowrap;
  line-height: 1;
  margin: 2px;
  cursor: pointer;
`
