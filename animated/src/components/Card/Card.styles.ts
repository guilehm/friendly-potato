import styled from "styled-components"

type SectionProps = {
  zoom: boolean
}

export const Section = styled.section<SectionProps>`
  transition: ease 500ms;
  ${(props) =>
    props.zoom
      ? `&:hover {
      transform: scale(1.05);
    }`
      : ""};
`

export const Title = styled.h2`
  font-size: 0.8rem;
  margin: 5px 0;
`

export const Image = styled.img`
  width: 100%;
  height: ${(props) => props.height ? `${props.height} px` : "200px"};
  object-fit: cover;
  border-radius: 12px 12px 0 0;
`
